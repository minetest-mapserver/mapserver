alter table blocks add column mtime bigint not null default 0;

create index BLOCKS_TIME on blocks(mtime);

create or replace function on_blocks_change() returns trigger as
$BODY$
BEGIN
    NEW.mtime = floor(EXTRACT(EPOCH from now()) * 1000);
    return NEW;
END;
$BODY$
LANGUAGE plpgsql;

create trigger blocks_update
 before insert or update
 on blocks
 for each row
 execute procedure on_blocks_change();
