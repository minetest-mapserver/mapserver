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
    this.onMinetestUpdate = this.onMinetestUpdate.bind(this);

    //update players all the time
    this.wsChannel.addListener("minetest-info", function(info){
      this.players = info.players || [];
    }.bind(this));
  },

  createMarker: function(player){
    var marker = L.marker([player.pos.z, player.pos.x], {icon: PlayerIcon});
    marker.bindPopup(player.name);

    return marker;
  },

  onMinetestUpdate: function(info){
    //TODO incremental update
    var self = this;

    this.players.forEach(function(player){
      if (self.currentObjects[player.name]){
        //marker exists
        self.currentObjects[player.name].setLatLng([player.pos.z, player.pos.x]);
        //setPopupContent

      } else {
        //marker does not exist
        var marker = self.createMarker(player);
        marker.addTo(self);

        self.currentObjects[player.name] = marker;
      }
    });

    Object.keys(self.currentObjects).forEach(function(existingName){
      var playerIsActive = self.players.find(function(p){
        return p.name == existingName;
      });

      if (!playerIsActive){
        //player
        self.currentObjects[existingName].remove();
        delete self.currentObjects[existingName];
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

      var marker = self.createMarker(player);
      marker.addTo(self);
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
