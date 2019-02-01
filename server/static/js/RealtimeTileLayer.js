'use strict';

function RealtimeTileLayer(wsChannel){
  var self = this;

  wsChannel.addListener("rendered-tile", function(tc){
    var id = self.getImageId(tc.layerid, tc.x, tc.y, tc.zoom);
    var el = document.getElementById(id);

    if (el){
        //Update src attribute if img found
        el.src = self.getTileSource(tc.layerid, tc.x, tc.y, tc.zoom, true);
    }
  });
}

RealtimeTileLayer.prototype.getTileSource = function(layerId, x,y,zoom,cacheBust){
    return "api/tile/" + layerId + "/" + x + "/" + y + "/" + zoom + "?_=" + Date.now();
};

RealtimeTileLayer.prototype.getImageId = function(layerId, x, y, zoom){
    return "tile-" + layerId + "/" + x + "/" + y + "/" + zoom;
};

RealtimeTileLayer.prototype.createLayer = function(layerId){
  var self = this;

  return L.TileLayer.extend({
    createTile: function(coords){
      var tile = document.createElement('img');
      tile.src = self.getTileSource(layerId, coords.x, coords.y, coords.z);
      tile.id = self.getImageId(layerId, coords.x, coords.y, coords.z);
      return tile;
    }
  });
};
