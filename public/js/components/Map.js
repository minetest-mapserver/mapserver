import SimpleCRS from "../utils/SimpleCRS.js";

export default {
    props: ["lat", "lon", "zoom", "layerId"],
    mounted: function() {
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

        console.log(map);
    },
    template: /*html*/`
        <div ref="target" style="height: 100%">
            Map {{lat}} / {{lon}} / {{zoom}} / {{layerId}}
        </div>
    `
};