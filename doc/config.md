
# Configuration

## Colors

There are builtin default colors, if you want new ones or override them
just put your `colors.txt` in the same directory the mapserver runs.

Example `colors.txt`:
```
# scifi nodes
scifi_nodes:slope_white 240 240 240
scifi_nodes:slope_vent 120 120 120
scifi_nodes:white2 240 240 240
```

Default colors, see: [colors directory](../public/colors)

## Configuration json

All config options reside in the `mapserver.json` file with the default values
The mapserver will generate a fresh `mapserver.json` if there is none at startup.

### Example mapserver.json
```json
{
	"port": 8080,
	"enableprometheus": true,
	"enablerendering": true,
	"webdev": false,
	"webapi": {
		"enablemapblock": false,
		"secretkey": "OYHuTRbhSQXkHcwu"
	},
	"layers": [
		{
			"id": 0,
			"name": "Base",
			"to": 160,
			"from": -16
		}
	],
	"renderingfetchlimit": 1000,
	"renderingjobs": 2,
	"renderingqueue": 100,
	"incrementalrenderingtimer": "5s",
	"mapobjects": {
		"bones": true,
		"protector": true,
		"technic": true,
		"luacontroller": true,
		"digiterms": true,
		"digilines": true,
		"travelnet": true,
		"mapserver": true,
		"mission": true,
		"jumpdrive": true,
		"smartshop": true,
		"fancyvend": true,
		"atm": true
	},
	"mapblockaccessor": {
		"expiretime": "10s",
		"purgetime": "15s",
		"maxitems": 5000
	},
    "skins": {
      "enableskinsdb": true,
      "skinspath": "/path/to/minetest/mods/skinsdb/textures"
    }
}
```

### Settings

#### port
The port on which the server listens to

#### webapi.secretkey
The generated secret for the [mod bridge](./mod.md)

#### layers
The layers as a list
More layers can be added here:
```json
"layers": [
  {
    "id": 0,
    "name": "Base",
    "to": 160,
    "from": -16
  },
  {
    "id": 1,
    "name": "Space",
    "to": 1600,
    "from": 1000
  }
],
```
*from* and *to* are in mapblocks (16x16x16 blocks)
Don't reuse the `id` after the tiles are generated.
If you make more substantial changes here you may have to remove all
existing tiles and start rendering from scratch.

#### renderingjobs
Number of cores to use for rendering, defaults to all available cores.
If CPU-limiting is desired, this is a good spot to begin with (Set to 1 for single-thread)

#### renderingfetchlimit
Number of mapblocks to collect at once while rendering:
* More means faster but also more RAM usage
* Less means slower but less RAM usage

For a small system (Raspberry PI) a setting of 1000 is ok.
Faster system can use the default (10'000)

#### incrementalrenderingtimer
Duration to let pass between incremental rendering executions.

#### enableprometheus
Enables the [Prometheus](./prometheus.md) metrics endpoint

#### mapblockaccessor.maxitems
Number of mapblocks to keep in memory, dial this down if you have memory issues

#### skins.enableskinsdb
Enables support for serving/displaying custom player skins provided by the SkinsDB mod.

#### skins.skinspath
The path to where SkinsDB textures are stored. This should be the SkinsDB textures directory.

Example: `/path/to/minetest/mods/skinsdb/textures`
