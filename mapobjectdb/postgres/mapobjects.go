package postgres

import (
	"database/sql"
	"mapserver/coords"
	"mapserver/mapobjectdb"
	"unicode/utf8"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (a *PostgresAccessor) GetMapData(q *mapobjectdb.SearchQuery) ([]*mapobjectdb.MapObject, error) {

	var rows *sql.Rows
	var err error

	if q.AttributeLike == nil {
		//plain pos search
		rows, err = a.db.Query(getMapDataPosQuery,
			q.Type,
			q.Pos1.X, q.Pos1.Y, q.Pos1.Z,
			q.Pos2.X, q.Pos2.Y, q.Pos2.Z,
			*q.Limit,
		)

	} else {
		//attribute like search
		rows, err = a.db.Query(getMapDataWithAttributeLikePosQuery,
			q.Type,
			q.Pos1.X, q.Pos1.Y, q.Pos1.Z,
			q.Pos2.X, q.Pos2.Y, q.Pos2.Z,
			q.AttributeLike.Key, q.AttributeLike.Value,
			*q.Limit,
		)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*mapobjectdb.MapObject, 0)
	var currentObj *mapobjectdb.MapObject
	var currentUID string

	for rows.Next() {
		var id string
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

		if currentUID == "" || currentUID != id {
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
			currentUID = id

			result = append(result, currentObj)

		}

		currentObj.Attributes[key] = value
	}

	return result, nil
}

func (a *PostgresAccessor) RemoveMapData(pos *coords.MapBlockCoords) error {
	_, err := a.db.Exec(removeMapDataQuery, pos.X, pos.Y, pos.Z)
	return err
}

func (a *PostgresAccessor) AddMapData(data *mapobjectdb.MapObject) error {

	for k, v := range data.Attributes {
		if !utf8.Valid([]byte(v)) {
			// invalid utf8, skip insert into db
			fields := logrus.Fields{
				"type":  data.Type,
				"value": v,
				"key":   k,
			}
			log.WithFields(fields).Info("Migration completed")
			return nil
		}
	}

	uid := uuid.NewString()
	_, err := a.db.Exec(addMapDataQuery,
		uid,
		data.X, data.Y, data.Z,
		data.MBPos.X, data.MBPos.Y, data.MBPos.Z,
		data.Type, data.Mtime)

	if err != nil {
		return err
	}

	for k, v := range data.Attributes {
		_, err := a.db.Exec(addMapDataAttributeQuery, uid, k, v)
		if err != nil {
			return err
		}
	}

	return nil
}
