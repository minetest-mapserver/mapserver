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

    this.currentObjects = []; //{obj:{}, marker: {}}

    this.reDraw = this.reDraw.bind(this);
    this.onMapMove = debounce(this.onMapMove.bind(this), 50);
    this.onMinetestUpdate = this.onMinetestUpdate.bind(this);
  },

  onMinetestUpdate: function(info){
    //TODO incremental update
  },

  reDraw: function(){
    //TODO full update
    var self = this;
    this.clearLayers();

    var mapLayer = this.layerMgr.getCurrentLayer()
    //TODO
  },

  onAdd: function(map) {
    this.layerMgr.addListener(this.reDraw);
    this.wsChannel.addListener("minetest-info", this.onMinetestUpdate);
    this.reDraw();
  },

  onRemove: function(map) {
    this.clearLayers();
    this.layerMgr.removeListener(this.reDraw);
    this.wsChannel.removeListener("minetest-info", this.onMinetestUpdate);
  }
});
