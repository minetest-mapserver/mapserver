import debounce from '../../util/debounce.js';
import wsChannel from '../../WebSocketChannel.js';
import layerMgr from '../../LayerManager.js';

import { getMapObjects } from '../../api.js';

export default L.LayerGroup.extend({
  initialize: function(type, icon) {
    L.LayerGroup.prototype.initialize.call(this);

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

  onLayerChange: function(/*layer*/){
    this.reDraw(true);
  },

  onMapMove: function(){
    this.reDraw(false);
  },

  getIcon: function(/*ob*/){
    return this.icon;
  },

  getMaxDisplayedZoom: function(){
    return 10;
  },

  reDraw: function(full){
    var self = this;

    if (this.map.getZoom() < this.getMaxDisplayedZoom()) {
      this.clearLayers();
      this.currentObjects = {};
      return;
    }

    if (full){
      this.clearLayers();
      this.currentObjects = {};
    }

    var mapLayer = layerMgr.getCurrentLayer();
    var min = this.map.getBounds().getSouthWest();
    var max = this.map.getBounds().getNorthEast();

    var y1 = parseInt(mapLayer.from);
    var y2 = parseInt(mapLayer.to);
    var x1 = parseInt(min.lng/16);
    var x2 = parseInt(max.lng/16);
    var z1 = parseInt(min.lat/16);
    var z2 = parseInt(max.lat/16);

    getMapObjects({
      pos1: { x:x1, y:y1, z:z1 },
      pos2: { x:x2, y:y2, z:z2 },
      type: this.type
    })
    .then(function(objects){
      //TODO: remove non-existing markers

      objects.forEach(function(obj){
        var hash = self.hashPos(obj.x, obj.y, obj.z);
        var marker = self.currentObjects[hash];
        var popup, icon;

        if (marker) {
          //marker exists
          icon = self.getIcon(obj);

          if (!icon) {
            //icon does not wanna be displayed anymore
            marker.remove();
            return;
          }
          //set popup, if changed
          popup = self.createPopup(obj);
          if (popup)
            marker.setPopupContent(popup);

          //redraw icon, if changed
          marker.setIcon(icon);

        } else {
          //marker does not exist
          icon = self.getIcon(obj);

          if (!icon) {
            //icon does not want to be displayed
            return;
          }

          marker = L.marker([obj.z + 0.5, obj.x + 0.5], {icon: icon});
          popup = self.createPopup(obj);
          if (popup)
            marker.bindPopup(popup);
          marker.addTo(self);

          self.currentObjects[hash] = marker;

        }
      });
    });

  },

  onAdd: function(map) {
    this.map = map;
    map.on("zoomend", this.onMapMove);
    map.on("moveend", this.onMapMove);
    wsChannel.addListener("mapobject-created", this.onMapObjectUpdated);
    this.reDraw(true);
  },

  onRemove: function(map) {
    this.clearLayers();
    map.off("zoomend", this.onMapMove);
    map.off("moveend", this.onMapMove);
    wsChannel.removeListener("mapobject-created", this.onMapObjectUpdated);
  }

});
