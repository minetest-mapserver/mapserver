-- recreate tables (can't add a referential constraint on an existing table)
create table objects_ng(
	uid text primary key not null,
	x int,
	y int,
	z int,
	posx int,
	posy int,
	posz int,
	type varchar,
	mtime bigint
);

create table object_attributes_ng(
	object_uid text not null,
	key varchar not null,
	value varchar not null,
	FOREIGN KEY (object_uid) references objects_ng(uid) ON DELETE CASCADE
	primary key(object_uid, key)
);

-- migrate data
insert into objects_ng(uid,x,y,z,posx,posy,posz,type,mtime)
	select id,x,y,z,posx,posy,posz,type,mtime from objects;
insert into object_attributes_ng(object_uid, key, value)
	select objectid, key, value from object_attributes;

-- remove old tables
drop table object_attributes;
drop table objects;

-- rename tables
alter table objects_ng rename to objects;
alter table object_attributes_ng rename to object_attributes;

-- journal mode, just in case
pragma journal_mode=wal;