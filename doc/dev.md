
# Build dependencies

* docker
* docker-compose

# Create the frontend bundle

```bash
docker-compose up mapserver_frontend
```

# Development setup (sqlite)

```bash
# start the engine in the first window/shell
docker-compose -f docker-compose.yml -f docker-compose.sqlite.yml up minetest
# and the mapserver in another
docker-compose -f docker-compose.yml -f docker-compose.sqlite.yml up mapserver
```
