
# Parameters

Mapserver command line parameters:

## Help
`./mapserver -help`
Shows all available commands


## Version
`./mapserver -version`
Shows the version, architecture and operating system:

```
Mapserver version: 0.0.2
OS: linux
Architecture: amd64
```


## Debug mode
`./mapserver -debug`
Enables the debug mode
It is advisable to pipe the debug output to a file for later inspection:

```
./mapserver -debug > debug.txt
```

## Config
`./mapserver -createconfig`
Creates a config and exits
