
export default L.Util.extend({}, L.CRS.Simple, {
    scale: function (zoom) {
        return Math.pow(2, zoom-9);
    }
});
