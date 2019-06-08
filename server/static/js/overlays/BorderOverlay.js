/* exported BorderOverlay */
/* globals AbstractGeoJsonOverlay: true */

var BorderOverlay = AbstractGeoJsonOverlay.extend({
  initialize: function(wsChannel, layerMgr) {
    AbstractGeoJsonOverlay.prototype.initialize.call(this, wsChannel, layerMgr, "border");
  },

  getMaxDisplayedZoom: function(){
    return 1;
  },

  getMinDisplayedZoom: function(){
    return 9;
  },

  createGeoJson: function(objects){
    var geoJsonLayer = L.geoJSON([], {
      onEachFeature: function(feature, layer){
        if (feature.properties && feature.properties.popupContent) {
          layer.bindPopup(feature.properties.popupContent);
        }
      },
      style: function(feature) {
        if (feature.properties && feature.properties.color){
          return {
            color: feature.properties.color,
            fill: feature.properties.color,
            opacity: 0.3
          };
        }
      }
    });

    var borders = [];
    var borderColors = {}; // { name: color }

    objects.forEach(function(obj){
      if (!obj.attributes.name)
        return;

      var border = borders[obj.attributes.name];
      if (!border){
        border = [];
        borders[obj.attributes.name] = border;
        borderColors[obj.attributes.name] = "#ff7800";
      }

      if (obj.attributes.color){
        borderColors[obj.attributes.name] = obj.attributes.color;
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

      // closing border
      coords.push([
        borders[bordername][0].x,
        borders[bordername][0].z
      ])

      var feature = {
        "type":"Feature",
        "geometry": {
          "type":"LineString",
          "coordinates":coords
        },
        "properties":{
            "name": bordername,
            "color": borderColors[bordername],
            "popupContent": "<b>Border (" + bordername + ")</b>"
        }
      };

      //line-points
      geoJsonLayer.addData(feature);
    });

    return geoJsonLayer;
  }

});
