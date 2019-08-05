
# Changelog

## Next

* Train overlay decluttering
* Updated builtin colors
* Hide travelnets and missions by pattern "(P)"
* Added benchmark
* Minor speed improvements

## 3.0.0

* Javascript ES6 frontend
* Param2 coloring
* More player infos (RTT)
* Search ability for LCD displays
* City border support
* Label support (street, city names, regions)
* Various bugfixes

## 2.2.0

* Added ATM overlay
* Added locator overlay
* Add ability for bone-owner search
* Separate mapserver_mod into own repository

## 2.1.1

* force tcp v4 for listener

## 2.1.0

* Added search bar
* Shop info with stock
* Colored POI's
* Icons for POI's
* Exported mapserver mod to own repository: https://github.com/thomasrudin/mapserver_mod

## 2.0.1
* Added *configversion* to `mapserver.json`
* Added *defaultoverlays* to `mapserver.json`
* Added trainline overlay
* Added border overlay
* Added simple js/css bundler (concat)
* Fix mismatching timestamp issues

## 2.0.0
* Fixed websocket connect issue (origin check)
* Fixed multiple layer issue
* Fixed mapblock version support issue
* *Breaking change:* Reworked `mapserver.json` layer config (now mapblock range, not blocks)
* Reworked build pipeline
* Optional crafting recipes for mapobjects

## 1.0.0
* Extended [configuration](config.md)
* [Monitoring](prometheus.md)
* [Mod](mod.md)-bridge
* Player and game-stats display (player-count, lag)
* More [mapobjects](mapobjects.md)
* Fixed tile generation glitch
* Tile-backend is now file-based (needs a complete re-rendering)
* Postgres map backend

## 0.1.0 (2019-02-08)
* Realtime tiles
* Mapobjects
* Sqlite backend

## 0.0.1 - 0.0.3
* Alpha releases
