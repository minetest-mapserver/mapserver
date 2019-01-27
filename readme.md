Minetest mapserver
=======

Realtime mapserver for Minetest

# Features

## Current features

* Initial and incremental map rendering
* Configurable layers (default: "Base" from y -16 to 160)

## Planned Features

* POI markers / mod integration
* Player markers and infos (HP, breath, etc)
* Street names / Train stuff

# Installation / Getting started

**Please make a backup of your world in case something goes wrong**

* Download the binary from the [releases](https://github.com/thomasrudin-mt/mapserver/releases) for your architecture and platform
* Drop the binary into your world folder (the one with the `world.mt` and `map.sqlite` files)
* Start the mapserver via command-line: `./mapserver` or `./mapserver.exe`
* Point your browser to `http://127.0.0.1:8080`

# Configuration

All config options reside in the `mapserver.json` file with the default values
Please stop the server if you make any changes there.

# Development state

* Early beta
* Working basic features (map rendering)
* Successor of http://github.com/thomasrudin-mt/minetest-tile-server

# Bugs

There will be bugs, please file them in the *issues* page.

# Contributions

Contributions are always welcome via pull/merge requests
