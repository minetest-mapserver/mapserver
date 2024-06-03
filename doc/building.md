
# Building the mapserver

Instructions to build the mapserver from source.

## Docker instructions

Build the image using the provided `Dockerfile`. That's all you have to do. The compiling process will be done inside containers.

## Build dependencies

* go >= 1.11 (for the binary)
* rollup >= 1.x (for the embedded js/css assets)

Ubuntu install: https://github.com/golang/go/wiki/Ubuntu

## Compile


Generate the js bundle for the frontend:
```
cd public
rollup -c rollup.config.js
```

Generate the `mapserver` binary:
```bash
# build the binary for the current platform
go build

# (optionally) run the unit-tests
go test ./...

```


