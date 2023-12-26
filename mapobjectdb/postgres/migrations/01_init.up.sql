
create table if not exists objects(
	id serial primary key,
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
	objectid serial,
	key varchar not null,
	value varchar not null,
	FOREIGN KEY (objectid) references objects(id) ON DELETE CASCADE,
	primary key(objectid, key)
);

create index if not exists object_attributes_key_value on object_attributes(key, value);

create table if not exists settings(
	key varchar primary key not null,
	value varchar not null
);
