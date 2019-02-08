
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

Default colors, see: [colors.txt](../server/static/colors.txt)

## Configuration json

All config options reside in the `mapserver.json` file with the default values
The mapserver will generate a fresh `mapserver.json` if there is none at startup.

### Example mapserver.json
```json
{
	"port": 8080,
	"enablerendering": true,
	"webdev": false,
	"webapi": {
		"enablemapblock": false,
		"secretkey": "ZJoSpysiKGlYexof"
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
	"renderingqueue": 100
}
```

### Settings

#### port
The port on which the server listens to

#### webapi.secretkey
The generated secret for the [mod bridge](./mod)

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
*from* and *to* are in blocks (not mapblocks)
Don't reuse the `id` after the tiles are generated.
If you make more substantial changes here you may have to remove all
existing tiles and start rendering from scratch.

#### renderingjobs
Number of cores to use for rendering, defaults to all available cores
