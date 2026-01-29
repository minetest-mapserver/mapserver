package sqlite

const getBlocksByMtimeQuery = `
	select x,y,z,data,mtime
	from blocks b
	where b.mtime > ?
	order by b.mtime asc
	limit ?
`

const getBlocksByMtimeQueryLegacy = `
	select pos,data,mtime
	from blocks b
	where b.mtime > ?
	order by b.mtime asc
	limit ?
`

const countBlocksQuery = `
	select count(*) from blocks b
`

const getBlockQuery = `
	select x,y,z,data,mtime from blocks b
	where b.x = ?
	and b.y = ?
	and b.z = ?
`

const getBlockQueryLegacy = `
	select pos,data,mtime from blocks b where b.pos = ?
`

const getTimestampQuery = `
	select strftime('%s', 'now')
`

const createMtimeColumnQuery = `
	alter table blocks add mtime integer default 0;
	create index blocks_mtime on blocks(mtime)
`

const createMtimeUpdateTriggerPosXYZ = `
	CREATE TRIGGER IF NOT EXISTS update_blocks_mtime_insert after insert on blocks for each row
	begin
	update blocks set mtime = strftime('%s', 'now')
		where x = new.x and y = new.y and z = new.z;
	end;

	CREATE TRIGGER IF NOT EXISTS update_blocks_mtime_update after update on blocks for each row
	begin
	update blocks set mtime = strftime('%s', 'now')
		where x = new.x and y = new.y and z = new.z;
	end;
`

const createMtimeUpdateTriggerPosLegacy = `
	CREATE TRIGGER IF NOT EXISTS update_blocks_mtime_insert after insert on blocks for each row
	begin
	update blocks set mtime = strftime('%s', 'now')
		where pos = new.pos;
	end;

	CREATE TRIGGER IF NOT EXISTS update_blocks_mtime_update after update on blocks for each row
	begin
	update blocks set mtime = strftime('%s', 'now')
		where pos = old.pos;
	end;
`

const getLastBlockQueryLegacy = `
	select pos,data,mtime
	from blocks b
	where b.pos > ?
	order by b.pos asc, b.mtime asc
	limit ?
`

const getBlocksByInitialTileQuery = `
	select x,y,z,data,mtime
	from blocks b
	where b.x >= ?
	and b.y >= ?
	and b.z >= ?
	and b.x <= ?
	and b.y <= ?
	and b.z <= ?
`
