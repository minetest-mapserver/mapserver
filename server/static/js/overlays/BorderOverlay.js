/* exported BorderOverlay */
/* globals AbstractGeoJsonOverlay: true */

var BorderOverlay = AbstractGeoJsonOverlay.extend({
  initialize: function(wsChannel, layerMgr) {
    AbstractGeoJsonOverlay.prototype.initialize.call(this, wsChannel, layerMgr, "border");
  },

  createGeoJson: function(objects){
    var geoJsonLayer = L.geoJSON([], {
      onEachFeature: function(feature, layer){
        if (feature.properties && feature.properties.popupContent) {
          layer.bindPopup(feature.properties.popupContent);
        }
      }
    });

    var borders = [];

    objects.forEach(function(obj){
      if (!obj.attributes.name)
        return;

      var border = borders[obj.attributes.name];
      if (!border){
        border = [];
        borders[obj.attributes.name] = border;
      }

      border.push(obj);
    });

    //Order by index and display
    Object.keys(borders).forEach(function(bordername){
      borders[bordername].sort(function(a,b){
        return parseInt(a.attributes.index) - parseInt(b.attributes.index);
      });

      var coords = [];

      //Add stations
      borders[bordername].forEach(function(entry){
        coords.push([entry.x, entry.z]);
      });

      var feature = {
        "type":"Feature",
        "geometry": {
          "type":"LineString",
          "coordinates":coords
        },
        "properties":{
            "name": bordername,
            "popupContent": "<b>Border (" + bordername + ")</b>"
        }
      };

      //line-points
      geoJsonLayer.addData(feature);
    });

    return geoJsonLayer;
  }

});
