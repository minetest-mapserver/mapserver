package sqlite

const migrateScript = `
PRAGMA foreign_keys = ON;
PRAGMA journal_mode = MEMORY;
PRAGMA synchronous = OFF; --TODO: this is just ridiculously slow otherwise...

create table if not exists objects(
	id integer primary key autoincrement,
  x int,
  y int,
	z int,
	posx int,
	posy int,
	posz int,
	type varchar,
  mtime bigint
);

create index if not exists objects_pos on objects(posx,posy,posz);
create index if not exists objects_pos_type on objects(posx,posy,posz,type);

create table if not exists object_attributes(
	objectid integer not null,
	key varchar not null,
	value varchar not null,
	FOREIGN KEY (objectid) references objects(id) ON DELETE CASCADE
	primary key(objectid, key)
);

create index if not exists object_attributes_key_value on object_attributes(key, value);

create table if not exists tiles(
  data blob,
  mtime bigint,
  layerid int,
  x int,
  y int,
  zoom int,
  primary key(x,y,zoom,layerid)
);

create table if not exists settings(
	key varchar primary key not null,
	value varchar not null
);
`

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

const getTileQuery = `
select data,mtime from tiles t
where t.layerid = ?
and t.x = ?
and t.y = ?
and t.zoom = ?
`

const setTileQuery = `
insert or replace into tiles(x,y,zoom,layerid,data,mtime)
values(?, ?, ?, ?, ?, ?)
`

const removeTileQuery = `
delete from tiles
where x = ? and y = ? and zoom = ? and layerid = ?
`

const getSettingQuery = `
select value from settings where key = ?
`

const setSettingQuery = `
insert or replace into settings(key, value)
values(?, ?)
`
