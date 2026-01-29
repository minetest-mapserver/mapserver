package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"mapserver/coords"
	"mapserver/db"
	"mapserver/types"
	"time"

	"github.com/sirupsen/logrus"
	_ "modernc.org/sqlite"
)

/*
sqlite extract: https://stackoverflow.com/questions/15448373/how-to-dump-a-file-stored-in-a-sqlite-database-as-a-blob
sqlite3 my.db "SELECT writefile('object0.gz', MyBlob) FROM MyTable WHERE id = 1"
*/

type Sqlite3Accessor struct {
	db         *sql.DB
	filename   string
	legacy_pos bool // legacy pos-column (instead of x/y/z columns)
}

func (db *Sqlite3Accessor) Migrate() error {

	//RW connection
	rwdb, err := sql.Open("sqlite", db.filename+"?mode=rw")
	if err != nil {
		return err
	}

	defer rwdb.Close()

	hasMtime := true
	_, err = rwdb.Query("select max(mtime) from blocks")
	if err != nil {
		hasMtime = false
	}

	_, err = rwdb.Query("select x, y, z from blocks limit 1")
	if err != nil {
		// x/y/z fields not found, set legacy flag
		db.legacy_pos = true
	}

	if !hasMtime {
		log.WithFields(logrus.Fields{
			"filename": db.filename,
		}).Info("Migrating database, this might take a while depending on the mapblock-count")
		start := time.Now()

		// create pos(mtime) column
		_, err = rwdb.Exec(createMtimeColumnQuery)
		if err != nil {
			return fmt.Errorf("create mtime column failed: %v", err)
		}

		// create trigger to update mtime column
		if db.legacy_pos {
			_, err = rwdb.Exec(createMtimeUpdateTriggerPosLegacy)
			if err != nil {
				return fmt.Errorf("create legacy trigger failed: %v", err)
			}
		} else {
			_, err = rwdb.Exec(createMtimeUpdateTriggerPosXYZ)
			if err != nil {
				return fmt.Errorf("create trigger failed: %v", err)
			}
		}

		t := time.Now()
		elapsed := t.Sub(start)
		log.WithFields(logrus.Fields{"elapsed": elapsed}).Info("Migration completed")
	}

	return nil
}

func convertRows(pos int64, data []byte, mtime int64) *db.Block {
	c := coords.PlainToCoord(pos)
	return &db.Block{Pos: c, Data: data, Mtime: mtime}
}

func (a *Sqlite3Accessor) FindBlocksByMtime(gtmtime int64, limit int) ([]*db.Block, error) {
	blocks := make([]*db.Block, 0)

	if a.legacy_pos {
		rows, err := a.db.Query(getBlocksByMtimeQueryLegacy, gtmtime, limit)
		if err != nil {
			return nil, err
		}

		defer rows.Close()

		for rows.Next() {
			var pos int64
			var data []byte
			var mtime int64

			err = rows.Scan(&pos, &data, &mtime)
			if err != nil {
				return nil, err
			}

			mb := convertRows(pos, data, mtime)
			blocks = append(blocks, mb)
		}
	} else {
		rows, err := a.db.Query(getBlocksByMtimeQuery, gtmtime, limit)
		if err != nil {
			return nil, err
		}

		defer rows.Close()

		for rows.Next() {
			mb := &db.Block{Pos: &types.MapBlockCoords{}}

			err = rows.Scan(&mb.Pos.X, &mb.Pos.Y, &mb.Pos.Z, &mb.Data, &mb.Mtime)
			if err != nil {
				return nil, err
			}

			blocks = append(blocks, mb)
		}
	}

	return blocks, nil
}

func (db *Sqlite3Accessor) CountBlocks() (int, error) {
	rows, err := db.db.Query(countBlocksQuery)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	if rows.Next() {
		var count int64

		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}

		return int(count), nil
	}

	return 0, nil
}

func (db *Sqlite3Accessor) GetTimestamp() (int64, error) {
	rows, err := db.db.Query(getTimestampQuery)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	if rows.Next() {
		var ts int64

		err = rows.Scan(&ts)
		if err != nil {
			return 0, err
		}

		return ts, nil
	}

	return 0, nil
}

func (a *Sqlite3Accessor) GetBlock(pos *types.MapBlockCoords) (*db.Block, error) {

	if a.legacy_pos {
		ppos := coords.CoordToPlain(pos)
		rows, err := a.db.Query(getBlockQueryLegacy, ppos)
		if err != nil {
			return nil, err
		}

		defer rows.Close()

		if rows.Next() {
			var pos int64
			var data []byte
			var mtime int64

			err = rows.Scan(&pos, &data, &mtime)
			if err != nil {
				return nil, err
			}

			mb := convertRows(pos, data, mtime)
			return mb, nil
		}
	} else {
		rows, err := a.db.Query(getBlockQuery, pos.X, pos.Y, pos.Z)
		if err != nil {
			return nil, err
		}

		defer rows.Close()

		if rows.Next() {
			mb := &db.Block{Pos: &types.MapBlockCoords{}}
			err = rows.Scan(&mb.Pos.X, &mb.Pos.Y, &mb.Pos.Z, &mb.Data, &mb.Mtime)
			if err != nil {
				return nil, err
			}

			return mb, nil
		}
	}

	return nil, nil
}

func (a *Sqlite3Accessor) intQuery(q string, params ...interface{}) int {
	rows, err := a.db.Query(q, params...)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var result int
		err = rows.Scan(&result)
		if err != nil {
			panic(err)
		}

		return result
	}

	panic("no result!")
}

func New(filename string) (*Sqlite3Accessor, error) {
	db, err := sql.Open("sqlite", filename+"?mode=ro")
	if err != nil {
		return nil, err
	}

	// sqlite connection limit
	db.SetMaxOpenConns(1)

	err = EnableWAL(db)
	if err != nil {
		return nil, err
	}

	sq := &Sqlite3Accessor{db: db, filename: filename}
	return sq, nil
}

func EnableWAL(db *sql.DB) error {
	result := db.QueryRow("pragma journal_mode;")
	var mode string
	err := result.Scan(&mode)
	if err != nil {
		return err
	}

	if mode != "wal" {
		_, err = db.Exec("pragma journal_mode = wal;")
		if err != nil {
			return errors.New("couldn't switch the db-journal to wal-mode, please stop the minetest-engine to allow doing this or do it manually: " + err.Error())
		}
	}

	return nil
}
