
# Incremental rendering

Incremental rendering works with the help of the *mtime* column
on the minetest database.

Every insert or update changes the *mtime* column to the current timestamp (with the help of triggers).
This way changes to the blocks can be detected by remembering the mtime of the
last query.


## Table *blocks* (minetest db)

posx	| posy	| posz	| data	| mtime
---	| ---	| ---	| ---	|
10	| 11	| 12	| ABC	| 1552977950000
20      | 21    | 22    | 123   | 1552977950010
30      | 31    | 32    | XYZ   | **1552977950020**
40      | 41    | 42    | A12   | 1552977950030
50      | 51    | 52    | B34   | 1552977950040

## Table *settings* (mapserver db)

key		| value
---		|
last\_mtime	| **1552977950020**

## Query example

The following query will return all changed blocks since the last call:

```
select posx,posy,posz,data,mtime
from blocks b
where b.mtime > 1552977950020
order by b.mtime asc
limit 1000

```

Additionally it will limit the returned rows so the mapserver can be started and stopped at any time
without processing all new data at once.

After that query the highest *mtime* is stored again in the mapserver database.

## Schedule

Incremental rendering is executed periodically:

* Without pause between calls if there is more data available (catch-up after mapserver downtime)
* With a 5 second pause between calls if there is no new data

## About realtime

Of course there are delays between placing/removing blocks and the tiles on the mapserver.
The minetest setting ** server\_map\_save\_interval ** is responsible for the delay to the mapserver (defaults to 5.3 seconds)
Don't try to decrease this value too much on your minetest instance, it has a performance impact!

