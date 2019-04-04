
# Mapobject-Database

The mapobject database is for all the mapobjects (POI, Shops, etc)
Per default it is initialized as a sqlite3 database in `mapserver.sqlite`

If you want to store your mapobjects in a proper database or need performance at
scale you can configure a postgresql server in your `world.mt`:

```
pgsql_mapserver_connection = host=127.0.0.1 port=5432 user=postgres password=enter dbname=postgres
```

The syntax is the same as for `pgsql_connection`
