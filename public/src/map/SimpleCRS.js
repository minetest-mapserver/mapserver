import L from 'leaflet';

// zoom 9 is 1:1 (one pixel per node)
const ZOOM_OFFSET = 9;

export default L.Util.extend({}, L.CRS.Simple, {
    scale: function (zoom) {
        return Math.pow(2, zoom - ZOOM_OFFSET);
    },

    // inverse of scale(), used by pinch-zoom (map.getScaleZoom) and fitBounds/flyTo
    zoom: function (scale) {
        return Math.log(scale) / Math.LN2 + ZOOM_OFFSET;
    }
});
