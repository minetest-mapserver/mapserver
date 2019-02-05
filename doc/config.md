
# Configuration

All config options reside in the `mapserver.json` file with the default values
Please stop the server if you make any changes there.

The mapserver will generate a fresh `mapserver.json` if there is none at startup.

## Example json
```json
{
	"port": 8080,
	"enablerendering": true,
	"webdev": false,
	"webapi": {
		"enablemapblock": false,
		"secretkey": "ZJoSpysiKGlYexof"
	},
	"renderstate": {
		"initialrun": false,
		"legacyprocessed": 16111,
		"lastx": 3,
		"lasty": 3,
		"lastz": 8,
		"lastmtime": 0
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

## Settings

### port
The port on which the server listens to

### webapi.secretkey
The generated secret for the [mod bridge](./install)

### layers
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

### renderingjobs
Number of cores to use for rendering, defaults to all available cores
