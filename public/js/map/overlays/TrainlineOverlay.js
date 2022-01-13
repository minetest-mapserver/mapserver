import AbstractGeoJsonOverlay from './AbstractGeoJsonOverlay.js';
import { getMapObjects } from '../../api.js';

export default AbstractGeoJsonOverlay.extend({
  initialize: function() {
    AbstractGeoJsonOverlay.prototype.initialize.call(this, "train");
    this.cache = {}; // { "A1":[] }
    this.pendingQueries = [];
    this.lastLayer = L.geoJSON([], {}); // empty dummy, will be replaced later
  },
  
  createGeoJson: function(objects){
    var self = this;
    // which unique lines do objects belong to?
    var lines = [];
    objects.forEach(function(obj){
      if (obj.attributes.line && lines.indexOf(obj.attributes.line) == -1) {
        lines.push(obj.attributes.line);
      }
    });
    // query for each line, add to cache
    lines.forEach(function(l){
      if (!self.cache[l]){
        // only request if not in cache.
        // if someone changed the train lines, the user has to reload. sorry.
        self.pendingQueries.push(l);
        getMapObjects({
          type: self.type,
          attributelike: {
            key: "line",
            value: l
          }
        })
        .then(function(objects){
          self.cache[l] = objects;
          self.pendingQueries = self.pendingQueries.filter(function(e){
            return e != l;
          });
          if (self.pendingQueries.length == 0) {
            // trigger a redraw
            // this will only happen if anything changed since it runs only after we completed a query
            self.clearLayers();

            var geoJsonLayer = self.createGeoJsonInternal(self.cache);
            geoJsonLayer.addTo(self);
          }
        });
      }
    });
    
    return self.lastLayer;
  },

  createGeoJsonInternal: function(lines){

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

    var lineColors = {}; // { "A1": "red" }

    // already sorted, determine color
    Object.keys(lines).forEach(function(linename){
      if (!lineColors[linename]){
        lineColors[linename] = "#ff7800";
        for (var i = lines[linename].length-1; i >= 0; i--) {
          // find the last element specifying a color
          // as was previous behaviour, but be more efficient
          if (lines[linename][i].attributes.color){
            lineColors[linename] = lines[linename][i].attributes.color;
            break;
          }
        }
      }
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

    this.lastLayer = geoJsonLayer;
    return geoJsonLayer;
  }

});
