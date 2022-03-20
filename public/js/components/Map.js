import SimpleCRS from "../utils/SimpleCRS.js";
import RealtimeTileLayer from '../utils/RealtimeTileLayer.js';
import ws from '../service/ws.js';

export default {
    props: ["lat", "lon", "zoom", "layerId"],
    mounted: function() {
        console.log("Map::mounted", this.lat, this.lon, this.zoom, this.layerId);
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

        map.attributionControl.addAttribution('<a href="https://github.com/minetest-mapserver/mapserver">Minetest Mapserver</a>');

        var tileLayer = new RealtimeTileLayer(ws, this.layerId, map);
        tileLayer.addTo(map);
      
        console.log(map);
    },
    methods: {
        updateMap: function() {
            console.log("Map::updateMap", this.lat, this.lon, this.zoom, this.layerId);
        }
    },
    watch: {
        "$route": "updateMap"
    },
    template: /*html*/`
        <div ref="target" style="height: 100%"></div>
    `
};