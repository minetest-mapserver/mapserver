#!/bin/sh
curl -X POST -k -H 'Content-Type: application/json' 'https://pandorabox.io/map/api/mapobjects/' --data '{
 "type":"train",
 "pos1":{"x":-2048,"y":-2048,"z":-2048},
 "pos2":{"x":2048,"y":2048,"z":2048}
}' | jq

