'use strict';

var AbstractIconOverlay = L.LayerGroup.extend({
  initialize: function(wsChannel, layerMgr, type, icon) {
    L.LayerGroup.prototype.initialize.call(this);

    this.layerMgr = layerMgr;
    this.wsChannel = wsChannel;
    this.type = type;
    this.icon = icon;

    this.currentObjects = {};

    this.onLayerChange = this.onLayerChange.bind(this);
    this.onMapMove = debounce(this.onMapMove.bind(this), 50);
  },

  hashPos: function(x,y,z){
    return x + "/" + y + "/" + z;
  },

  onLayerChange: function(layer){
    this.reDraw(true);
  },

  onMapMove: function(){
    this.reDraw(false);
  },

  reDraw: function(full){
    var self = this;

    if (this._map.getZoom() < 10) {
      this.clearLayers();
      this.currentObjects = {};
      return;
    }

    if (full){
      this.clearLayers();
      this.currentObjects = {};
    }

    var mapLayer = this.layerMgr.getCurrentLayer()
    var min = this._map.getBounds().getSouthWest();
    var max = this._map.getBounds().getNorthEast();

    var y1 = parseInt(mapLayer.from/16);
    var y2 = parseInt(mapLayer.to/16);
    var x1 = parseInt(min.lng);
    var x2 = parseInt(max.lng);
    var z1 = parseInt(min.lat);
    var z2 = parseInt(max.lat);

    api.getMapObjects(
      x1, y1, y1,
      x2, y2, z2,
      this.type)
    .then(function(objects){
      //TODO: remove non-existing markers

      objects.forEach(function(obj){
        var hash = self.hashPos(obj.x, obj.y, obj.z);

        if (self.currentObjects[hash]) {
          //marker exists
          //TODO: update popup

        } else {
          //marker does not exist
          var marker = L.marker([obj.z, obj.x], {icon: self.icon});
          marker.bindPopup(self.createPopup(obj));
          marker.addTo(self);

          self.currentObjects[hash] = marker;

        }
      });
    })

  },

  onAdd: function(map) {
    map.on("zoomend", this.onMapMove);
    map.on("moveend", this.onMapMove);
    this.layerMgr.addListener(this.onLayerChange);
    this.reDraw(true)
  },

  onRemove: function(map) {
    this.clearLayers();
    map.off("zoomend", this.onMapMove);
    map.off("moveend", this.onMapMove);
    this.layerMgr.removeListener(this.onLayerChange);
  }

});
