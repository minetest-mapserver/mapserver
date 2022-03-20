import Map from "../components/Map.js";

export default {
    components: {
        "map-component": Map
    },
    template: /*html*/`
        <map-component
            :lat="$route.params.lat"
            :lon="$route.params.lon"
            :zoom="$route.params.zoom"
            :layerId="$route.params.layerId"
        />
    `
};