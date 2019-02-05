package postgres

import (
	"database/sql"
	_ "https://github.com/lib/pq"
	"mapserver/coords"
	"mapserver/db"
)


type PostgresAccessor struct {
	db       *sql.DB
}

func (db *PostgresAccessor) Migrate() error {
	return nil
}

func convertRows(pos int64, data []byte, mtime int64) db.Block {
	c := coords.PlainToCoord(pos)
	return db.Block{Pos: c, Data: data, Mtime: mtime}
}


func (this *PostgresAccessor) FindBlocksByMtime(gtmtime int64, limit int) ([]db.Block, error) {
	blocks := make([]db.Block, 0)

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

func (this *PostgresAccessor) FindLegacyBlocksByPos(lastpos coords.MapBlockCoords, limit int) ([]db.Block, error) {
	blocks := make([]db.Block, 0)
	pc := coords.CoordToPlain(lastpos)

	rows, err := this.db.Query(getLastBlockQuery, pc, limit)
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


func (this *PostgresAccessor) GetBlock(pos coords.MapBlockCoords) (*db.Block, error) {
	ppos := coords.CoordToPlain(pos)

	rows, err := this.db.Query(getBlockQuery, ppos)
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
		return &mb, nil
	}

	return nil, nil
}

func NewPostgresAccessor(filename string) (*PostgresAccessor, error) {
	return nil, nil
}
