'use strict';

var PlayerIcon = L.icon({
  iconUrl: 'pics/sam.png',

  iconSize:     [16, 32],
  iconAnchor:   [8, 16],
  popupAnchor:  [0, -16]
});

var PlayerOverlay = L.LayerGroup.extend({
  initialize: function(wsChannel, layerMgr) {
    L.LayerGroup.prototype.initialize.call(this);

    this.layerMgr = layerMgr;
    this.wsChannel = wsChannel;

    this.currentObjects = [];

    this.onLayerChange = this.onLayerChange.bind(this);
    this.reDraw = this.reDraw.bind(this);
    this.onMapMove = debounce(this.onMapMove.bind(this), 50);

    this.wsChannel.addListener("minetest-info", this.onMinetestUpdate.bind(this));
  },

  onLayerChange: function(layer){
    this.reDraw();
  },

  onMinetestUpdate: function(info){
    //TODO
  },

  reDraw: function(){
    var self = this;

    this.clearLayers();

    var mapLayer = this.layerMgr.getCurrentLayer()
    //TODO
  },

  onAdd: function(map) {
    this.layerMgr.addListener(this.reDraw);
    this.reDraw(true)
  },

  onRemove: function(map) {
    this.clearLayers();
    this.layerMgr.removeListener(this.reDraw);
  }
});
