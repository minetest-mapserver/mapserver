package postgres

const migrateScript = `
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
`

const getBlocksByMtimeQuery = `
select posx,posy,posz,data,mtime
from blocks b
where b.mtime > ?
order by b.mtime asc
limit ?
`

const getLastBlockQuery = `
select posx,posy,posz,data,mtime
from blocks b
where b.mtime = 0
and b.posx >= ?
and b.posy >= ?
and b.posz >= ?
order by b.posx asc, b.posy asc, b.posz asc, b.mtime asc
limit ?
`

const countBlocksQuery = `
select count(*) from blocks b where b.mtime >= ? and b.mtime <= ?
`

const getBlockQuery = `
select pos,data,mtime from blocks b
where b.posx = ?
and b.posy = ?
and b.posz = ?
`
