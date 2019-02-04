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

    this.currentObjects = {}; // name => marker
    this.players = [];

    this.reDraw = this.reDraw.bind(this);
    this.onMapMove = debounce(this.onMapMove.bind(this), 50);
    this.onMinetestUpdate = this.onMinetestUpdate.bind(this);

    //update players all the time
    this.wsChannel.addListener("minetest-info", function(info){
      this.players = info.players || [];
    }.bind(this));
  },

  onMinetestUpdate: function(info){
    //TODO incremental update
    var self = this;

    this.players.forEach(function(player){
      if (self.currentObjects[player.name]){
        //marker exists
      } else {
        //marker does not exist
      }
    });

    Object.keys(self.currentObjects).forEach(function(existingName){
      var playerIsActive = this.players.find(function(p){
        return p.name == existingName;
      });

      if (!playerIsActive){
        //player
      }
    });
  },

  reDraw: function(){
    var self = this;
    this.currentObjects = {};
    this.clearLayers();

    var mapLayer = this.layerMgr.getCurrentLayer()

    this.players.forEach(function(player){
      //TODO: filter layer

      var marker = L.marker([travelnet.z, travelnet.x], {icon: TravelnetIcon});
      marker.bindPopup(player.name).addTo(self);

      self.currentObjects[player.name] = marker;
    });

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
