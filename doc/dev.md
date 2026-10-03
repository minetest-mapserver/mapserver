
# Build dependencies

* docker
* docker-compose

# Create the frontend bundle

```bash
docker-compose up mapserver_frontend
```

This builds the frontend into `public/dist` and rebuilds it on changes (served by the mapserver in `webdev` mode).

# Frontend development (vite)

The frontend is a Vue 3 app (`public/src`), built with [vite](https://vite.dev).
With a mapserver running on `localhost:8080` the vite dev-server (with hot-reload) can be used:

```bash
cd public
npm ci
npm run dev
# lint
npm run lint
```

The `/api` requests are proxied to `http://localhost:8080` (override with the `MAPSERVER_URL` env variable).

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