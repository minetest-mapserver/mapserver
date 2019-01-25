package mapobjectdb

import (
	"mapserver/coords"
)

const getMapDataPosQuery = `
select o.id, o.type, o.mtime,
 o.x, o.y, o.z,
 o.posx, o.posy, o.posz,
 oa.key, oa.value
from objects o
left join object_attributes oa on o.id = oa.objectid
where o.type = ?
and o.posx >= ? and o.posy >= ? and o.posz >= ?
and o.posx <= ? and o.posy <= ? and o.posz <= ?
order by o.id
`


func (db *Sqlite3Accessor) GetMapData(q SearchQuery) ([]*MapObject, error) {
	rows, err := db.db.Query(getMapDataPosQuery,
		q.Type,
		q.Pos1.X, q.Pos1.Y, q.Pos1.Z,
		q.Pos2.X, q.Pos2.Y, q.Pos2.Z,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*MapObject, 0)
	var currentObj *MapObject

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

		if currentObj == nil {
			//TODO
		} else {
			//TODO
		}

	}

	return result, nil
}

const removeMapDataQuery = `
delete from objects where posx = ? and posy = ? and posz = ?
`

func (db *Sqlite3Accessor) RemoveMapData(pos *coords.MapBlockCoords) error {
	_, err := db.db.Exec(removeMapDataQuery, pos.X, pos.Y, pos.Z)
	return err
}

const addMapDataQuery = `
insert into
objects(x,y,z,posx,posy,posz,type,mtime)
values(?, ?, ?, ?, ?, ?, ?, ?)
`

const addMapDataAttributeQuery = `
insert into
object_attributes(objectid, key, value)
values(?, ?, ?)
`

func (db *Sqlite3Accessor) AddMapData(data *MapObject) error {
	res, err := db.db.Exec(addMapDataQuery,
		data.X, data.Y, data.Z,
		data.MBPos.X, data.MBPos.Y, data.MBPos.Z,
		data.Type, data.Mtime)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return err
	}

	for k, v := range data.Attributes {
		//TODO: batch insert
		_, err := db.db.Exec(addMapDataAttributeQuery, id, k, v)

		if err != nil {
			return err
		}
	}

	return nil
}
