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
