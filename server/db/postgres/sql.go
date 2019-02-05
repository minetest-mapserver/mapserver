package postgres

const migrateScript = `
alter table blocks add mtime integer default 0;
create index blocks_mtime on blocks(mtime);

CREATE TRIGGER update_blocks_mtime_insert after insert on blocks for each row
begin
update blocks set mtime = strftime('%s', 'now') where pos = new.pos;
end;

CREATE TRIGGER update_blocks_mtime_update after update on blocks for each row
begin
update blocks set mtime = strftime('%s', 'now') where pos = old.pos;
end;
`

const getBlocksByMtimeQuery = `
select pos,data,mtime
from blocks b
where b.mtime > ?
order by b.mtime asc
limit ?
`

const getLastBlockQuery = `
select pos,data,mtime
from blocks b
where b.mtime = 0
and b.pos > ?
order by b.pos asc, b.mtime asc
limit ?
`

const countBlocksQuery = `
select count(*) from blocks b where b.mtime >= ? and b.mtime <= ?
`


const getBlockQuery = `
select pos,data,mtime from blocks b where b.pos = ?
`
