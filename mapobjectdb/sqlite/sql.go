package sqlite

const getMapDataPosQuery = `
select o.uid, o.type, o.mtime,
 o.x, o.y, o.z,
 o.posx, o.posy, o.posz,
 oa.key, oa.value
from objects o
left join object_attributes oa on o.uid = oa.object_uid
where o.type = ?
and o.posx >= ? and o.posy >= ? and o.posz >= ?
and o.posx <= ? and o.posy <= ? and o.posz <= ?
order by o.uid
limit ?
`

const getMapDataWithAttributeLikePosQuery = `
select o.uid, o.type, o.mtime,
 o.x, o.y, o.z,
 o.posx, o.posy, o.posz,
 oa.key, oa.value
from objects o
left join object_attributes oa on o.uid = oa.object_uid
where o.uid in (
  select object_uid from object_attributes where key = ? and value like ?
)
and o.type = ?
and o.posx >= ? and o.posy >= ? and o.posz >= ?
and o.posx <= ? and o.posy <= ? and o.posz <= ?
order by o.uid
limit ?
`

const removeMapDataAttributesQuery = `
delete from object_attributes where object_uid in (select uid from objects where posx = ? and posy = ? and posz = ?)
`

const removeMapDataQuery = `
delete from objects where posx = ? and posy = ? and posz = ?
`

const addMapDataQuery = `
insert into
objects(uid,x,y,z,posx,posy,posz,type,mtime)
values(?, ?, ?, ?, ?, ?, ?, ?, ?)
`

const addMapDataAttributeQuery = `
insert into
object_attributes(object_uid, key, value)
values(?, ?, ?)
`

const getSettingQuery = `
select value from settings where key = ?
`

const setSettingQuery = `
insert or replace into settings(key, value)
values(?, ?)
`
