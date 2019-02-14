package postgres

const getBlocksByMtimeQuery = `
select posx,posy,posz,data,mtime
from blocks b
where b.mtime > ?
order by b.mtime asc
limit ?
`

const countBlocksQuery = `
select count(*) from blocks b where b.mtime >= ? and b.mtime <= ?
`

const getBlockQuery = `
select posx,posy,posz,data,mtime from blocks b
where b.posx = ?
and b.posy = ?
and b.posz = ?
`
