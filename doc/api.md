
# Api documentation

REST Api documentation

## Mapobjects

Query for `bones` / `poi`/ `shop` / etc

* **Path:** `api/mapobjects`
* **Method:** `POST`
* **Consumes:** `application/json`
* **Produces:** `application/json`

POST-Payload:
```json
{
  "pos1": {
    "x":-2048,
    "y":-2048,
    "z":-2048
  },
  "pos2": {
    "x":2048,
    "y":2048,
    "z":2048
  },
  "type":"bones"
}
```

**Note**: `pos1` and `pos2` are in mapblocks


Example query:
```bash
curl 'http://127.0.0.1:8080/api/mapobjects/' \
 -H 'Content-Type: application/json; charset=utf-8' \
 --data '{"pos1":{"x":-2048,"y":-2048,"z":-2048},"pos2":{"x":2048,"y":2048,"z":2048},"type":"bones"}' \
 | jq
```

Result:
```json
[{
    "mapblock": {
      "x": -1671,
      "y": 0,
      "z": -82
    },
    "x": -26729,
    "y": 1,
    "z": -1306,
    "type": "bones",
    "mtime": 1554099532,
    "attributes": {
      "owner": "Brzezowski58",
      "time": "0"
    }
}]
```
