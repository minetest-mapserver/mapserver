
# Stuck incremental rendering

In case the incremental rendering is stuck (a bug between 1.0.1 and 2.0.0)
NOTE: this is for the mapserver database, NOT the map database!

## Sqlite

Execute this command in the sqlite database (`sqlite3 ./mapserver.sqlite`)

```
update settings
set value = strftime('%s', 'now')
where key = 'last_mtime';
```

This will reset the current rendering time to now

## Postgres

For postgresql:

Enter the shell (for example):
```
psql -U postgres
```

And reset the current mtime
```
update settings
set value = floor(EXTRACT(EPOCH from now()) * 1000)
where key = 'last_mtime';
```

