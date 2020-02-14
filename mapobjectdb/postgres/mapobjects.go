package postgres

import (
	"database/sql"
	"mapserver/coords"
	"mapserver/mapobjectdb"
)

func (db *PostgresAccessor) GetMapData(q *mapobjectdb.SearchQuery) ([]*mapobjectdb.MapObject, error) {

	var rows *sql.Rows
	var err error
	var limit = 1000

	if q.Limit != nil {
		limit = *q.Limit
	}

	if q.AttributeLike == nil {
		//plain pos search
		rows, err = db.db.Query(getMapDataPosQuery,
			q.Type,
			q.Pos1.X, q.Pos1.Y, q.Pos1.Z,
			q.Pos2.X, q.Pos2.Y, q.Pos2.Z,
			limit,
		)

	} else {
		//attribute like search
		rows, err = db.db.Query(getMapDataWithAttributeLikePosQuery,
			q.Type,
			q.Pos1.X, q.Pos1.Y, q.Pos1.Z,
			q.Pos2.X, q.Pos2.Y, q.Pos2.Z,
			q.AttributeLike.Key, q.AttributeLike.Value,
			limit,
		)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*mapobjectdb.MapObject, 0)
	var currentObj *mapobjectdb.MapObject
	var currentId *int64

	for rows.Next() {
		var id int64
		var Type string
		var mtime int64
		var x, y, z int
		var posx, posy, posz int
		var key, value string

		err = rows.Scan(&id, &Type, &mtime,
			&x, &y, &z, &posx, &posy, &posz,
			&key, &value,
		)

		if err != nil {
			return nil, err
		}

		if currentId == nil || *currentId != id {
			pos := coords.NewMapBlockCoords(posx, posy, posz)
			mo := &mapobjectdb.MapObject{
				MBPos:      pos,
				Type:       Type,
				X:          x,
				Y:          y,
				Z:          z,
				Mtime:      mtime,
				Attributes: make(map[string]string),
			}

			currentObj = mo
			currentId = &id

			result = append(result, currentObj)

		}

		currentObj.Attributes[key] = value
	}

	return result, nil
}

func (db *PostgresAccessor) RemoveMapData(pos *coords.MapBlockCoords) error {
	_, err := db.db.Exec(removeMapDataQuery, pos.X, pos.Y, pos.Z)
	return err
}

func (db *PostgresAccessor) AddMapData(data *mapobjectdb.MapObject) error {
	res := db.db.QueryRow(addMapDataQuery,
		data.X, data.Y, data.Z,
		data.MBPos.X, data.MBPos.Y, data.MBPos.Z,
		data.Type, data.Mtime)

	lastInsertId := 0
	err := res.Scan(&lastInsertId)

	if err != nil {
		return err
	}

	for k, v := range data.Attributes {
		//TODO: batch insert
		_, err := db.db.Exec(addMapDataAttributeQuery, lastInsertId, k, v)

		if err != nil {
			return err
		}
	}

	return nil
}
