package sqlite


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

const removeMapDataQuery = `
delete from objects where posx = ? and posy = ? and posz = ?
`

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

const getSettingQuery = `
select value from settings where key = ?
`

const setSettingQuery = `
insert or replace into settings(key, value)
values(?, ?)
`
