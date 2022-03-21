import SimpleCRS from "../utils/SimpleCRS.js";
import RealtimeTileLayer from '../utils/RealtimeTileLayer.js';
import ws from '../service/ws.js';
import { getLayerById } from "../service/layer.js";

export default {
    props: ["lat", "lon", "zoom", "layerId"],
    mounted: function() {
        const layer = getLayerById(this.layerId);
        console.log("Map::mounted", this.lat, this.lon, this.zoom, this.layerId, layer);

        const map = L.map(this.$refs.target, {
            minZoom: 2,
            maxZoom: 12,
            center: [this.lat, this.lon],
            zoom: this.zoom,
            crs: SimpleCRS,
            maxBounds: L.latLngBounds(
              L.latLng(-31000, -31000),
              L.latLng(31000, 31000)
            )
        });


        const updateLink = () => {
            const center = map.getCenter();
            const lon = Math.floor(center.lng);
            const lat = Math.floor(center.lat);
            console.log("Map::updateLink", map.getZoom(), lon, lat);
            // change hash route
            this.$router.push({
                name: "map",
                params: {
                    lat: lat,
                    lon: lon,
                    zoom: map.getZoom(),
                    layerId: this.layerId
                }
            });
        };

        // listen for route change
        map.on('zoomend', updateLink);
        map.on('moveend', updateLink);

        // add attribution
        map.attributionControl.addAttribution('<a href="https://github.com/minetest-mapserver/mapserver">Minetest Mapserver</a>');

        // TODO: all layers
        var tileLayer = new RealtimeTileLayer(ws, this.layerId, map);
        tileLayer.addTo(map);
      
        console.log(map);
    },
    methods: {
        updateMap: function() {
            const layer = getLayerById(this.layerId);
            console.log("Map::updateMap", this.lat, this.lon, this.zoom, this.layerId, layer);
        }
    },
    watch: {
        "$route": "updateMap"
    },
    template: /*html*/`
        <div ref="target" style="height: 100%"></div>
    `
};