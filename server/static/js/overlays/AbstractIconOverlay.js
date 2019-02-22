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
    var hash = this.hashPos(obj.x, obj.y, obj.z);
    var marker = this.currentObjects[hash];

    if (marker) {
      //marker exists
      var popup = this.createPopup(obj);
      if (popup)
        marker.setPopupContent(popup);

      marker.setIcon(this.getIcon(obj));
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
    var x1 = parseInt(min.lng/16);
    var x2 = parseInt(max.lng/16);
    var z1 = parseInt(min.lat/16);
    var z2 = parseInt(max.lat/16);

    api.getMapObjects({
      pos1: { x:x1, y:y1, z:z1 },
      pos2: { x:x2, y:y2, z:z2 },
      type: this.type
    })
    .then(function(objects){
      //TODO: remove non-existing markers

      objects.forEach(function(obj){
        var hash = self.hashPos(obj.x, obj.y, obj.z);
        var marker = self.currentObjects[hash];

        if (marker) {
          //marker exists

          //set popup, if changed
          var popup = self.createPopup(obj);
          if (popup)
            marker.setPopupContent(popup);

          //redraw icon, if changed
          marker.setIcon(self.getIcon(obj));

        } else {
          //marker does not exist
          marker = L.marker([obj.z, obj.x], {icon: self.getIcon(obj)});
          var popup = self.createPopup(obj);
          if (popup)
            marker.bindPopup(popup);
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
