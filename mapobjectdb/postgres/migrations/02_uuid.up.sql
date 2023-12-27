
-- add objects.uid
alter table objects add uid uuid not null default gen_random_uuid();
create index objects_uid on objects(uid);
alter table objects add constraint objects_uid_unique UNIQUE (uid);

-- add object_attributes.object_uid and update references
alter table object_attributes add object_uid uuid;
update object_attributes set object_uid = (select uid from objects o where o.id = objectid);
alter table object_attributes alter column object_uid set not null;
alter table object_attributes
    add constraint object_attributes_object_uid_fk
    FOREIGN KEY (object_uid)
    REFERENCES objects(uid)
    on delete cascade;
create index object_attributes_object_uid on object_attributes(object_uid);

-- drop old id's
alter table object_attributes drop objectid;
alter table objects drop id;