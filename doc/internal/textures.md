
# convert

```sh
sudo apt install imagemagick
convert markers-soft.png -crop 36x46+0+0 -resize 28% mapserver_poi_red.png
convert markers-soft.png -crop 36x46+36+0 -resize 28% mapserver_poi_orange.png
convert markers-soft.png -crop 36x46+72+0 -resize 28% mapserver_poi_green.png
convert markers-soft.png -crop 36x46+108+0 -resize 28% mapserver_poi_blue.png
convert markers-soft.png -crop 36x46+144+0 -resize 28% mapserver_poi_violet.png
convert markers-soft.png -crop 36x46+180+0 -resize 28% mapserver_poi_brown.png
```
