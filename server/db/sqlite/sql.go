package sqlite

const getBlocksByMtimeQuery = `
select pos,data,mtime
from blocks b
where b.mtime > ?
order by b.mtime asc
limit ?
`

const countBlocksQuery = `
select count(*) from blocks b where b.mtime >= ? and b.mtime <= ?
`

const getBlockQuery = `
select pos,data,mtime from blocks b where b.pos = ?
`
