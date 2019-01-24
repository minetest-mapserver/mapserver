
```sql
update blocks set mtime = 0 where pos = (select max(pos) from blocks);
select max(mtime) from blocks;
```
