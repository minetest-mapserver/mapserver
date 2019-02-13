PRAGMA foreign_keys = ON;
PRAGMA journal_mode = MEMORY;

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

create table if not exists settings(
	key varchar primary key not null,
	value varchar not null
);
