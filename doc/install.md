
# Mapserver installation

**Please make a backup of your world in case something goes wrong**

* Download the binary from the [releases](https://github.com/thomasrudin-mt/mapserver/releases) for your architecture and platform
* Drop the binary into your world folder (the one with the `world.mt` and `map.sqlite` files)
* Start the mapserver via command-line: `./mapserver` or `./mapserver.exe`
* Point your browser to `http://127.0.0.1:8080`

For additional infos (lag,time,players => active mode) on the mapserver interface you should install the [mapserver-mod](mod.md)

## Performance / Scalability

For small to medium setups the default values should suffice.
If you have a bigger map (say: above 10 GB) you should configure the mapserver accordingly:

* See [Mapobject-Database](./mapobjectdb.md) for a scalable mapserver database
