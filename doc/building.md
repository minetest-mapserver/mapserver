
# Building the mapserver

Instructions to build the mapserver from source

## Build dependencies

* go >= 1.11 (for the binary)
* rollup >= 1.x (for the embedded js/css assets)

Ubuntu install: https://github.com/golang/go/wiki/Ubuntu

## Compile


Generate the js bundle for the frontend:
```
cd static/js
rollup -c rollup.config.js
```

Generate the `mapserver` binary:
```bash
# generate the static web files
# this step embeds the generated js/css/html files for inclusion in the resulting binary
go generate

# build the binary for the current playtform
go build

# (optionally) run the unit-tests
go test ./...

```


