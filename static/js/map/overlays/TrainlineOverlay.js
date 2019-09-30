import AbstractGeoJsonOverlay from './AbstractGeoJsonOverlay.js';

export default AbstractGeoJsonOverlay.extend({
  initialize: function() {
    AbstractGeoJsonOverlay.prototype.initialize.call(this, "train");
  },

  createGeoJson: function(objects){

    var geoJsonLayer = L.geoJSON([], {
      onEachFeature: function(feature, layer){
        if (feature.properties && feature.properties.popupContent) {
          layer.bindPopup(feature.properties.popupContent);
        }
      },
      pointToLayer: function (feature, latlng) {
        var geojsonMarkerOptions = {
          radius: 8,
          weight: 1,
          opacity: 1,
          fillOpacity: 0.8
        };

        return L.circleMarker(latlng, geojsonMarkerOptions);
      },
      style: function(feature) {
        if (feature.properties && feature.properties.color){
          return { color: feature.properties.color };
        }
      }
    });

    var lines = {}; // { "A1":[] }
    var lineColors = {}; // { "A1": "red" }

    //Sort and add lines
    objects.forEach(function(obj){
      if (!obj.attributes.line)
        return;

      var line = lines[obj.attributes.line];
      if (!line){
        line = [];
        lines[obj.attributes.line] = line;
        //default or new color
        lineColors[obj.attributes.line] = "#ff7800";
      }

      if (obj.attributes.color){
        //new color
        lineColors[obj.attributes.line] = obj.attributes.color;
      }

      line.push(obj);
    });

    //Order by index and display
    Object.keys(lines).forEach(function(linename){
      lines[linename].sort(function(a,b){
        return parseInt(a.attributes.index) - parseInt(b.attributes.index);
      });

      var coords = [];
      var stations = [];

      //Add stations
      lines[linename].forEach(function(entry){
        coords.push([entry.x, entry.z]);

        if (entry.attributes.station) {
          stations.push({
            "type": "Feature",
            "properties": {
              "name": entry.attributes.station,
              "color": lineColors[linename],
              "popupContent": "<b>Train-station (Line " + entry.attributes.line + ")</b><hr>" +
                entry.attributes.station
            },
            "geometry": {
              "type": "Point",
              "coordinates": [entry.x, entry.z]
            }
          });
        }
      });

      var feature = {
        "type":"Feature",
        "geometry": {
          "type":"LineString",
          "coordinates":coords
        },
        "properties":{
            "name": linename,
            "color": lineColors[linename],
            "popupContent": "<b>Train-line (" + linename + ")</b>"
        }
      };

      //line-points
      geoJsonLayer.addData(feature);

      //stations
      stations.forEach(function(stationfeature){
        geoJsonLayer.addData(stationfeature);
      });


    });

    return geoJsonLayer;
  }

});
