package postgres

const getMapDataPosQuery = `
select o.uid, o.type, o.mtime,
 o.x, o.y, o.z,
 o.posx, o.posy, o.posz,
 oa.key, oa.value
from objects o
left join object_attributes oa on o.uid = oa.object_uid
where o.type = $1
and o.posx >= $2 and o.posy >= $3 and o.posz >= $4
and o.posx <= $5 and o.posy <= $6 and o.posz <= $7
order by o.uid
limit $8
`
const getMapDataWithAttributeLikePosQuery = `
select o.uid, o.type, o.mtime,
 o.x, o.y, o.z,
 o.posx, o.posy, o.posz,
 oa.key, oa.value
from objects o
left join object_attributes oa on o.uid = oa.object_uid
where o.uid in (
  select object_uid from object_attributes where key = $8 and value ilike $9
)
and o.type = $1
and o.posx >= $2 and o.posy >= $3 and o.posz >= $4
and o.posx <= $5 and o.posy <= $6 and o.posz <= $7
order by o.uid
limit $10
`

const removeMapDataQuery = `
delete from objects where posx = $1 and posy = $2 and posz = $3
`

const addMapDataQuery = `
insert into
objects(uid, x,y,z,posx,posy,posz,type,mtime)
values($1, $2, $3, $4, $5, $6, $7, $8, $9)
`

const addMapDataAttributeQuery = `
insert into
object_attributes(object_uid, key, value)
values($1, $2, $3)
`

const getSettingQuery = `
select value from settings where key = $1
`

const setSettingQuery = `
insert into settings(key, value)
values($1, $2)
on conflict(key)
do update set value = EXCLUDED.value
`
