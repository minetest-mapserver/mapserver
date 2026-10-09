import L from 'leaflet';

export default L.TileLayer.extend({

  initialize: function(wsChannel, layerId, map) {
    L.TileLayer.prototype.initialize.call(this);

    this.wsChannel = wsChannel;
    this.layerId = layerId;
    this.map = map;
    this.onRenderedTile = this.onRenderedTile.bind(this);
  },

  onRenderedTile: function(tc){
    if (tc.layerid != this.layerId){
      //ignore other layers
      return;
    }

    if (tc.zoom != this.map.getZoom()){
      //ignore other zoom levels
      return;
    }

    var id = this.getImageId(tc.x, tc.y, tc.zoom);
    var el = document.getElementById(id);

    if (el){
        //Update src attribute if img found
        el.src = this.getTileSource(tc.x, tc.y, tc.zoom, true);
    }
  },

  onAdd: function(map){
    this.wsChannel.addListener("rendered-tile", this.onRenderedTile);
    return L.TileLayer.prototype.onAdd.call(this, map);
  },

  onRemove: function(map){
    this.wsChannel.removeListener("rendered-tile", this.onRenderedTile);
    return L.TileLayer.prototype.onRemove.call(this, map);
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
    tile.onload = () => {
      tile.onload = undefined;
      done(null, tile);
    };
    tile.onerror = e => done(e, tile);

    return tile;
  }
});
