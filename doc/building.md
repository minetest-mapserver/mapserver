
# Building the mapserver

Instructions to build the mapserver from source

## Build dependencies

* go >= 1.11 (for the binary)
* nodejs >= 22 / npm (for the embedded frontend)

Ubuntu install: https://github.com/golang/go/wiki/Ubuntu

## Compile


Build the frontend (output in `public/dist`, embedded into the binary):
```bash
cd public
npm ci
npm run build
```

Generate the `mapserver` binary:
```bash
# build the binary for the current platform
go build

# (optionally) run the unit-tests
go test ./...

```


