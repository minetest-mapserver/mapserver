
# System overview

<img src="./Overview.png">

# Build dependencies

* go >= 1.16
* nodejs >= v17.4.0
* npm >= 8.3.0

# Create the frontend bundle

```bash
cd public
npm ci
npm run bundle
```

# Development setup

Working directory: `./server`

## Preparing the files and map

Copy your `map.sqlite` into the working directory if you want to test with
a sqlite map database

### world.mt

You need a `world.mt` too in order to make the connection to the database.
In the sqlite case:

```
gameid = minetest
backend = sqlite3
creative_mode = false
enable_damage = false
player_backend = files
```

For postgres:
```
gameid = minetest
backend = postgresql
creative_mode = true
enable_damage = true
player_backend = postgresql
pgsql_connection = host=localhost port=5432 user=postgres password=enter dbname=postgres
pgsql_player_connection = host=localhost port=5432 user=postgres password=enter dbname=postgres
```

## Running the server

* Create a `mapserver.json` with `go run . -createconfig`
* Change the value `webdev` in the `mapserver.json` to `true`
* Start the server with `go run .` or with debug output: `go run . -debug`
* The web files in `public/` can now be changed on the fly without restarting the server
