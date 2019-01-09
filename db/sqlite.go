package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"mapserver/coords"
	"time"
	"errors"
)

var log *logrus.Entry
func init(){
	log = logrus.WithFields(logrus.Fields{"prefix": "db/sqlite.Migrate"})
}

const migrateScript = `
alter table blocks add mtime timestamp default NULL;
update blocks set mtime = current_timestamp;
create index blocks_mtime on blocks(mtime);

CREATE TRIGGER update_blocks_mtime_insert after insert on blocks for each row
begin
update blocks set mtime = current_timestamp where pos = new.pos;
end;

CREATE TRIGGER update_blocks_mtime_update after update on blocks for each row
begin
update blocks set mtime = current_timestamp where pos = old.pos;
end;
`

type Sqlite3Accessor struct {
	db *sql.DB
	filename string
}

func (db *Sqlite3Accessor) Migrate() error {

	//RW connection
	rwdb, err := sql.Open("sqlite3", db.filename + "?mode=rw")
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
		log.WithFields(logrus.Fields{"filename":db.filename}).Info("Migrating database")
		start := time.Now()
		_, err = rwdb.Exec(migrateScript)
		if err != nil {
			return err
		}
		t := time.Now()
		elapsed := t.Sub(start)
		log.WithFields(logrus.Fields{"elapsed":elapsed}).Info("Migration completed")
	}

	return nil
}

func (db *Sqlite3Accessor) FindLatestBlocks(mintime int64, limit int) ([]Block, error) {
	return make([]Block, 0), nil
}

func (db *Sqlite3Accessor) FindBlocks(pos1, pos2 coords.MapBlockCoords) ([]Block, error) {
	return make([]Block, 0), nil
}

const countBlockQuery = `
select count(*) from blocks b
where b.pos >= ? and b.pos <= ?
`

func (db *Sqlite3Accessor) CountBlocks(pos1, pos2 coords.MapBlockCoords) (int, error) {
	ppos1 := coords.CoordToPlain(pos1)
	ppos2 := coords.CoordToPlain(pos2)

	if ppos1 > ppos2 {
		ppos1, ppos2 = ppos2, ppos1
	}

	rows, err := db.db.Query(countBlockQuery, ppos1, ppos2)
	if err != nil {
		return 0, err
	}

	if !rows.Next() {
		return 0, errors.New("No results")
	}

	var count int
	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func NewSqliteAccessor(filename string) (*Sqlite3Accessor, error) {
	db, err := sql.Open("sqlite3", filename + "?mode=ro")
	if err != nil {
		return nil, err
	}

	sq := &Sqlite3Accessor{db: db, filename: filename}
	return sq, nil
}
