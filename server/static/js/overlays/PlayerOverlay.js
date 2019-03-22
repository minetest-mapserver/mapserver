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

    var html = "<b>" + player.name + "</b>";
    html += "<hr>";

    for (var i=0; i<Math.floor(player.hp / 2); i++)
      html += "<img src='pics/heart.png'>";

    if (player.hp % 2 == 1)
      html += "<img src='pics/heart_half.png'>";

    html += "<br>";

    for (var i=0; i<Math.floor(player.breath / 2); i++)
      html += "<img src='pics/bubble.png'>";

    if (player.breath % 2 == 1)
      html += "<img src='pics/bubble_half.png'>";


    marker.bindPopup(html);
    return marker;
  },

  isPlayerInCurrentLayer: function(player){
    var mapLayer = this.layerMgr.getCurrentLayer()

    return (player.pos.y >= mapLayer.from && player.pos.y <= mapLayer.to)
  },

  onMinetestUpdate: function(info){
    var self = this;

    this.players.forEach(function(player){
      var isInLayer = self.isPlayerInCurrentLayer(player);

      if (!isInLayer){
        if (self.currentObjects[player.name]){
          //player is displayed and not on the layer anymore
          //Remove the marker and reference
          self.currentObjects[player.name].remove();
          delete self.currentObjects[player.name];
        }

        return;
      }

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
      if (!self.isPlayerInCurrentLayer(player)){
        //not in current layer
        return;
      }

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
