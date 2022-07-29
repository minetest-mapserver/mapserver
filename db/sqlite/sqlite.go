package sqlite

import (
	"database/sql"
	"errors"
	"mapserver/coords"
	"mapserver/db"
	"mapserver/public"
	"time"

	"github.com/sirupsen/logrus"
	_ "modernc.org/sqlite"
)

/*
sqlite extract: https://stackoverflow.com/questions/15448373/how-to-dump-a-file-stored-in-a-sqlite-database-as-a-blob
sqlite3 my.db "SELECT writefile('object0.gz', MyBlob) FROM MyTable WHERE id = 1"
*/

type Sqlite3Accessor struct {
	db       *sql.DB
	filename string
}

func (db *Sqlite3Accessor) Migrate() error {

	//RW connection
	rwdb, err := sql.Open("sqlite", "file:"+db.filename+"?mode=rw")
	if err != nil {
		return err
	}

	defer rwdb.Close()

	hasMtime := true
	_, err = rwdb.Query("select max(mtime) from blocks")
	if err != nil {
		hasMtime = false
	}

	if !hasMtime {
		log.WithFields(logrus.Fields{
			"filename": db.filename,
		}).Info("Migrating database, this might take a while depending on the mapblock-count")
		start := time.Now()

		sql, err := public.Files.ReadFile("sql/sqlite_mapdb_migrate.sql")
		if err != nil {
			return err
		}

		_, err = rwdb.Exec(string(sql))
		if err != nil {
			return err
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

	rows, err := a.db.Query(getBlocksByMtimeQuery, gtmtime, limit)
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

func (db *Sqlite3Accessor) GetBlock(pos *coords.MapBlockCoords) (*db.Block, error) {
	ppos := coords.CoordToPlain(pos)

	rows, err := db.db.Query(getBlockQuery, ppos)
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

	return nil, nil
}

func New(filename string) (*Sqlite3Accessor, error) {
	db, err := sql.Open("sqlite", "file:"+filename)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("pragma busy_timeout = 5000;")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("pragma journal_mode = wal;")
	if err != nil {
		return nil, errors.New("could not set db to wal-mode, please stop the minetest-server to allow this one-time fixup")
	}

	sq := &Sqlite3Accessor{db: db, filename: filename}
	return sq, nil
}
