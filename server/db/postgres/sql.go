package postgres

const getBlocksByInitialTileQuery = `
select posx,posy,posz,data,mtime
from blocks b
where b.posx >= $1
and b.posy >= $2
and b.posz >= $3
and b.posx <= $4
and b.posy <= $5
and b.posz <= $6
`

const getBlockCountByInitialTileQuery = `
select count(*)
from blocks b
where b.posx >= $1
and b.posy >= $2
and b.posz >= $3
and b.posx <= $4
and b.posy <= $5
and b.posz <= $6
`

const getBlocksByMtimeQuery = `
select posx,posy,posz,data,mtime
from blocks b
where b.mtime > $1
order by b.mtime asc
limit $2
`

const countBlocksQuery = `
select count(*) from blocks where mtime >= $1 and mtime <= $2
`

const getBlockQuery = `
select posx,posy,posz,data,mtime from blocks b
where b.posx = $1
and b.posy = $2
and b.posz = $3
`
