
(function(){

    var crs = L.Util.extend({}, L.CRS.Simple, {
        //transformation: L.transformation(0.001, 0, -0.001, 0),
        scale: function (zoom) {
            return Math.pow(2, zoom-9);
        }
    });

    var initialZoom = 11;
    var initialCenter = [0, 0];

    var map = L.map('image-map', {
      minZoom: 2,
      maxZoom: 12,
      center: initialCenter,
      zoom: initialZoom,
      crs: crs
    });

    function getTileSource(layerId, x,y,zoom,cacheBust){
        return "api/tile/" + layerId + "/" + x + "/" + y + "/" + zoom + (cacheBust ? "?_=" + Date.now() : "");
    }

    function getImageId(layerId, x, y, zoom){
        return "tile-" + layerId + "/" + x + "/" + y + "/" + zoom;
    }

    function createTileLayer(layerId){
        return L.TileLayer.extend({
          createTile: function(coords){
            var tile = document.createElement('img');
            tile.src = getTileSource(layerId, coords.x, coords.y, coords.z);
            tile.id = getImageId(layerId, coords.x, coords.y, coords.z);
            return tile;
          }
        });
    }

    function updateTile(data){
        var id = getImageId(data.layerId, data.x, data.y, data.zoom);
        var el = document.getElementById(id);

        if (el){
            //Update src attribute if img found
            el.src = getTileSource(data.layerId, data.x, data.y, data.zoom, true);
        }
    }

    var layers = {};

    var Layer = createTileLayer(0);
    var tileLayer = new Layer();
    tileLayer.addTo(map);

    L.control.layers(layers, {}).addTo(map);

    // coord display

    L.Control.CoordinatesDisplay = L.Control.extend({
        onAdd: function(map) {
      var div = L.DomUtil.create('div', 'leaflet-bar leaflet-custom-display');
      function update(ev){
        var latlng = ev.latlng;
        div.innerHTML = "X:" + parseInt(latlng.lng) + " Z:" + parseInt(latlng.lat);
      }

      map.on('mousemove', update);
      map.on('click', update);
      map.on('touch', update);

      return div;
        },

        onRemove: function(map) {
        }
    });

    L.control.coordinatesDisplay = function(opts) {
        return new L.Control.CoordinatesDisplay(opts);
    }

    var el = L.control.coordinatesDisplay({ position: 'bottomleft' });
    el.addTo(map);



})()
