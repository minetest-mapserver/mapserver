package postgres

import (
	"database/sql"
	"mapserver/coords"
	"mapserver/db"
	"mapserver/public"
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
		log.Info("Migrating database, this might take a while depending on the mapblock-count")
		start := time.Now()

		sql, err := public.Files.ReadFile("sql/postgres_mapdb_migrate.sql")
		if err != nil {
			return err
		}

		_, err = db.db.Exec(string(sql))
		if err != nil {
			return err
		}
		t := time.Now()
		elapsed := t.Sub(start)
		log.WithFields(logrus.Fields{"elapsed": elapsed}).Info("Migration completed")
	}

	return nil
}

func convertRows(posx, posy, posz int, data []byte, mtime int64) *db.Block {
	c := coords.NewMapBlockCoords(posx, posy, posz)
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
		var posx, posy, posz int
		var data []byte
		var mtime int64

		err = rows.Scan(&posx, &posy, &posz, &data, &mtime)
		if err != nil {
			return nil, err
		}

		mb := convertRows(posx, posy, posz, data, mtime)
		blocks = append(blocks, mb)
	}

	return blocks, nil
}

func (this *PostgresAccessor) CountBlocks(frommtime, tomtime int64) (int, error) {
	rows, err := this.db.Query(countBlocksQuery, frommtime, tomtime)
	if err != nil {
		panic(err)
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

func (db *PostgresAccessor) GetTimestamp() (int64, error) {
	rows, err := db.db.Query(getTimestampQuery)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	if rows.Next() {
		var ts float64

		err = rows.Scan(&ts)
		if err != nil {
			return 0, err
		}

		return int64(ts), nil
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
		var posx, posy, posz int
		var data []byte
		var mtime int64

		err = rows.Scan(&posx, &posy, &posz, &data, &mtime)
		if err != nil {
			return nil, err
		}

		mb := convertRows(posx, posy, posz, data, mtime)
		return mb, nil
	}

	return nil, nil
}

func New(connStr string) (*PostgresAccessor, error) {
	db, err := sql.Open("postgres", connStr+" sslmode=disable")
	if err != nil {
		return nil, err
	}

	sq := &PostgresAccessor{db: db}
	return sq, nil
}
