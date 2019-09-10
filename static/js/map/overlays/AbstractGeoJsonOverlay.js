import debounce from '../../util/debounce.js';
import { getMapObjects } from '../../api.js';

export default L.LayerGroup.extend({
  initialize: function(wsChannel, layerMgr, type) {
    L.LayerGroup.prototype.initialize.call(this);

    this.layerMgr = layerMgr;
    this.wsChannel = wsChannel;
    this.type = type;

    this.onLayerChange = this.onLayerChange.bind(this);
    this.onMapMove = debounce(this.onMapMove.bind(this), 50);
  },

  onLayerChange: function(){
    this.reDraw();
  },

  onMapMove: function(){
    this.reDraw();
  },

  createStyle: function(){
	   //TODO: default style
  },

  createGeoJson: function(objects){
    var self = this;

    var geoJsonLayer = L.geoJSON([], {
        onEachFeature: function(feature, layer){
            if (feature.properties && feature.properties.popupContent) {
                layer.bindPopup(feature.properties.popupContent);
            }
        },
        style: self.createStyle.bind(self)
    });

    objects.forEach(function(obj){
      geoJsonLayer.addData(self.createFeature(obj));
    });

    return geoJsonLayer;
  },

  getMaxDisplayedZoom: function(){
    return 7;
  },

  getMinDisplayedZoom: function(){
    return 13;
  },

  reDraw: function(){
    var self = this;
    var zoom = this.map.getZoom();

    if (zoom < this.getMaxDisplayedZoom() || zoom > this.getMinDisplayedZoom()) {
      this.clearLayers();
      return;
    }

    var mapLayer = this.layerMgr.getCurrentLayer();
    var min = this._map.getBounds().getSouthWest();
    var max = this._map.getBounds().getNorthEast();

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
      self.clearLayers();

      var geoJsonLayer = self.createGeoJson(objects);
      geoJsonLayer.addTo(self);
    });

  },

  onAdd: function(map) {
    this.map = map;
    map.on("zoomend", this.onMapMove);
    map.on("moveend", this.onMapMove);
    this.layerMgr.addListener(this.onLayerChange);
    this.reDraw(true);
  },

  onRemove: function(map) {
    this.clearLayers();
    map.off("zoomend", this.onMapMove);
    map.off("moveend", this.onMapMove);
    this.layerMgr.removeListener(this.onLayerChange);
  }

});
