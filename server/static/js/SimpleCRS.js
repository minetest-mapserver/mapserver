/* exported SimpleCRS */

var SimpleCRS = L.Util.extend({}, L.CRS.Simple, {
    scale: function (zoom) {
        return Math.pow(2, zoom-9);
    }
});
