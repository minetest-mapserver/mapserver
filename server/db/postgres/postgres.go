package postgres

import (
	"database/sql"
	"mapserver/coords"
	"mapserver/db"
	"mapserver/vfs"
	"time"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type PostgresAccessor struct {
	db *sql.DB
}

func (db *PostgresAccessor) Migrate() error {
	hasMtime := true
	_, err := db.db.Query("select max(mtime) from blocks")
	if err != nil {
		hasMtime = false
	}

	if !hasMtime {
		log.Info("Migrating database")
		start := time.Now()
		_, err = db.db.Exec(vfs.FSMustString(false, "/sql/postgres_mapdb_migrate.sql"))
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

func (this *PostgresAccessor) FindBlocksByMtime(gtmtime int64, limit int) ([]*db.Block, error) {
	blocks := make([]*db.Block, 0)

	rows, err := this.db.Query(getBlocksByMtimeQuery, gtmtime, limit)
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

func (this *PostgresAccessor) CountBlocks(frommtime, tomtime int64) (int, error) {
	rows, err := this.db.Query(countBlocksQuery, frommtime, tomtime)
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

func (this *PostgresAccessor) GetBlock(pos *coords.MapBlockCoords) (*db.Block, error) {
	rows, err := this.db.Query(getBlockQuery, pos.X, pos.Y, pos.Z)
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

func New(connStr string) (*PostgresAccessor, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	sq := &PostgresAccessor{db: db}
	return sq, nil
}
