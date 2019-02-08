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
    this.onMapObjectUpdated = this.onMapObjectUpdated.bind(this);
    this.onMapMove = debounce(this.onMapMove.bind(this), 50);
  },

  //websocket update
  onMapObjectUpdated: function(obj){
    var hash = self.hashPos(obj.x, obj.y, obj.z);
    var marker = self.currentObjects[hash];

    if (marker) {
      //marker exists
      marker.setPopupContent(self.createPopup(obj));
    }
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

  getIcon: function(obj){
    return this.icon;
  },

  reDraw: function(full){
    var self = this;

    if (this.map.getZoom() < 10) {
      this.clearLayers();
      this.currentObjects = {};
      return;
    }

    if (full){
      this.clearLayers();
      this.currentObjects = {};
    }

    var mapLayer = this.layerMgr.getCurrentLayer()
    var min = this.map.getBounds().getSouthWest();
    var max = this.map.getBounds().getNorthEast();

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
        var marker = self.currentObjects[hash];

        if (marker) {
          //marker exists
          marker.setPopupContent(self.createPopup(obj));

        } else {
          //marker does not exist
          marker = L.marker([obj.z, obj.x], {icon: self.getIcon(obj)});
          marker.bindPopup(self.createPopup(obj));
          marker.addTo(self);

          self.currentObjects[hash] = marker;

        }
      });
    })

  },

  onAdd: function(map) {
    this.map = map;
    map.on("zoomend", this.onMapMove);
    map.on("moveend", this.onMapMove);
    this.layerMgr.addListener(this.onLayerChange);
    this.wsChannel.addListener("mapobject-created", this.onMapObjectUpdated);
    this.reDraw(true)
  },

  onRemove: function(map) {
    this.clearLayers();
    map.off("zoomend", this.onMapMove);
    map.off("moveend", this.onMapMove);
    this.layerMgr.removeListener(this.onLayerChange);
    this.wsChannel.removeListener("mapobject-created", this.onMapObjectUpdated);
  }

});
