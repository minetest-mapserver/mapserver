Minetest mapserver
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-6-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->
=======

![](https://github.com/minetest-mapserver/mapserver/workflows/jshint/badge.svg)
![](https://github.com/minetest-mapserver/mapserver/workflows/go-test/badge.svg)
![](https://github.com/minetest-mapserver/mapserver/workflows/build/badge.svg)

![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/minetest-mapserver/mapserver)
![GitHub repo size](https://img.shields.io/github/repo-size/minetest-mapserver/mapserver.svg)
![GitHub closed issues](https://img.shields.io/github/issues-closed/minetest-mapserver/mapserver.svg)
![GitHub issues](https://img.shields.io/github/issues/minetest-mapserver/mapserver)

![GitHub All Releases](https://img.shields.io/github/downloads/minetest-mapserver/mapserver/total)
![Docker Pulls](https://img.shields.io/docker/pulls/minetestmapserver/mapserver)

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/minetest-mapserver/mapserver)
[![Go Report Card](https://goreportcard.com/badge/github.com/minetest-mapserver/mapserver)](https://goreportcard.com/report/github.com/minetest-mapserver/mapserver)
[![Coverage Status](https://coveralls.io/repos/github/minetest-mapserver/mapserver/badge.svg)](https://coveralls.io/github/minetest-mapserver/mapserver)

<img src="./doc/pics/General_map_preview.png">

Realtime mapserver for [Minetest](https://minetest.net)

Demo: [Pandorabox Server map](https://pandorabox.io/map/#-1782.25/493.5/10)

# Documentation

* [Installation](doc/install.md)
* [Building](doc/building.md)
* [Mapserver mod](doc/mod.md)
* [Parameters](doc/params.md)
* [Search](doc/search.md)
* [Configuration](doc/config.md)
* [Recommended specs](doc/recommended_specs.md)
* [Stats webfragment](doc/stats_webfragment.md)
* [Web API](doc/api.md)
* [Contribution](doc/contrib.md)
* [Development](doc/dev.md)
* [License](doc/license.md)
* [Changelog](doc/changelog.md)

# How it works

See: [Incremental rendering](doc/incrementalrendering.md)

# Compatibility

* Minetest 0.4.15 - 0.4.17.1
* Minetest 5.0

# Features

## Current features

* Click-and-run installation
* Initial and incremental map rendering
* Param2 coloring
* Realtime rendering and map-updating
* Realtime player and world stats
* [Search](doc/search.md) bar
* Configurable layers (default: "Base" from y -16 to 160)
* POI [markers](doc/mapobjects.md) / [mod](doc/mod.md) integration
* Protector display
* LCD Displays as markers
* Monitoring with [Prometheus](doc/prometheus.md)

## Planned Features

* Isometric view
* Skin support
* Route planning (via travelnets / trains)

# Supported map-databases
The connection is auto-detected from your `world.mt`:

* Sqlite3
* PostgreSql

# Screenshots

## Web interface
<img src="./pics/web.png">

## Terminal
<img src="./pics/terminal.png">

## Map objects (as markers)
Enable/Disable those in the [Configuration](doc/config.md)
See:  [mapobjects](doc/mapobjects.md)


# Bugs

There will be bugs, please file them in the [issues](https://github.com/minetest-mapserver/mapserver/issues) page.

# Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://github.com/BuckarooBanzay"><img src="https://avatars.githubusercontent.com/u/39065740?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Buckaroo Banzai</b></sub></a><br /><a href="https://github.com/BuckarooBanzay/mapserver/commits?author=BuckarooBanzay" title="Code">💻</a></td>
    <td align="center"><a href="http://photo.pyrollo.com/"><img src="https://avatars.githubusercontent.com/u/13189280?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Pierre-Yves Rollo</b></sub></a><br /><a href="https://github.com/BuckarooBanzay/mapserver/commits?author=pyrollo" title="Code">💻</a></td>
    <td align="center"><a href="http://peter.nerlich4u.de/"><img src="https://avatars.githubusercontent.com/u/10530729?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Peter Nerlich</b></sub></a><br /><a href="https://github.com/BuckarooBanzay/mapserver/commits?author=PeterNerlich" title="Code">💻</a></td>
    <td align="center"><a href="https://github.com/crocsg"><img src="https://avatars.githubusercontent.com/u/34553036?v=4?s=100" width="100px;" alt=""/><br /><sub><b>StephaneG</b></sub></a><br /><a href="https://github.com/BuckarooBanzay/mapserver/commits?author=crocsg" title="Code">💻</a></td>
    <td align="center"><a href="https://arvados.org/"><img src="https://avatars.githubusercontent.com/u/149135?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Ward Vandewege</b></sub></a><br /><a href="https://github.com/BuckarooBanzay/mapserver/commits?author=cure" title="Code">💻</a></td>
    <td align="center"><a href="https://blog.bmarwell.de/"><img src="https://avatars.githubusercontent.com/u/1413391?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Benjamin Marwell</b></sub></a><br /><a href="https://github.com/BuckarooBanzay/mapserver/commits?author=bmarwell" title="Documentation">📖</a></td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!