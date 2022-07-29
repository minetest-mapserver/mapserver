
export default L.TileLayer.extend({

  initialize: function(wsChannel, layerId, map) {
    L.TileLayer.prototype.initialize.call(this);

    var self = this;
    this.layerId = layerId;

    wsChannel.addListener("rendered-tile", function(tc){
      if (tc.layerid != self.layerId){
        //ignore other layers
        return;
      }

      if (tc.zoom != map.getZoom()){
        //ignore other zoom levels
        return;
      }

      var id = self.getImageId(tc.x, tc.y, tc.zoom);
      var el = document.getElementById(id);

      if (el){
          //Update src attribute if img found
          el.src = self.getTileSource(tc.x, tc.y, tc.zoom, true);
      }
    });
  },

  getTileSource: function(x,y,zoom,cacheBust){
      return "api/tile/" + this.layerId + "/" + x + "/" + y + "/" + zoom + (cacheBust ? "?_=" + Date.now() : "");
  },

  getImageId: function(x, y, zoom){
      return "tile-" + this.layerId + "/" + x + "/" + y + "/" + zoom;
  },

  createTile: function(coords, done){
    var tile = document.createElement('img');
    tile.src = this.getTileSource(coords.x, coords.y, coords.z, true);
    tile.id = this.getImageId(coords.x, coords.y, coords.z);

    // trigger callbacks
    tile.onload = () => done(null, tile);
    tile.onerror = e => done(e, tile);

    return tile;
  }
});
