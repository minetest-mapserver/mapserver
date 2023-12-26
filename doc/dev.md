
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
docker-compose up minetest
# and the mapserver in another
docker-compose up mapserver
```

# Development setup (postgres)

```bash
# start postgres in the background
docker-compose -f docker-compose.yml -f docker-compose.postgres.yml up -d postgres
# start the engine in the first window/shell
docker-compose -f docker-compose.yml -f docker-compose.postgres.yml up minetest
# and the mapserver in another
docker-compose -f docker-compose.yml -f docker-compose.postgres.yml up mapserver
```

Utilities:
```sh
# psql
docker-compose -f docker-compose.yml -f docker-compose.postgres.yml exec postgres psql -U postgres
```