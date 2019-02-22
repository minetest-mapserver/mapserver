'use strict';

//TODO
var TrainIcon = L.icon({
  iconUrl: 'pics/sam.png',

  iconSize:     [16, 32],
  iconAnchor:   [8, 16],
  popupAnchor:  [0, -16]
});

var TrainOverlay = L.LayerGroup.extend({
  initialize: function(wsChannel, layerMgr) {
    L.LayerGroup.prototype.initialize.call(this);

    this.layerMgr = layerMgr;
    this.wsChannel = wsChannel;

    this.currentObjects = {}; // name => marker
    this.trains = [];

    this.reDraw = this.reDraw.bind(this);
    this.onMinetestUpdate = this.onMinetestUpdate.bind(this);

    //update players all the time
    this.wsChannel.addListener("minetest-info", function(info){
      this.trains = info.trains || [];
    }.bind(this));
  },

  createMarker: function(train){
    var marker = L.marker([train.pos.z, train.pos.x], {icon: TrainIcon});
    marker.bindPopup("Train");

    return marker;
  },

  onMinetestUpdate: function(info){
    var self = this;

    this.trains.forEach(function(train){
      if (self.currentObjects[train.id]){
        //marker exists
        self.currentObjects[train.id].setLatLng([train.pos.z, train.pos.x]);
        //setPopupContent

      } else {
        //marker does not exist
        var marker = self.createMarker(train);
        marker.addTo(self);

        self.currentObjects[train.id] = marker;
      }
    });

    Object.keys(self.currentObjects).forEach(function(existingId){
      var trainIsActive = self.trains.find(function(t){
        return t.id == existingId;
      });

      if (!trainIsActive){
        //train
        self.currentObjects[existingId].remove();
        delete self.currentObjects[existingId];
      }
    });
  },

  reDraw: function(){
    var self = this;
    this.currentObjects = {};
    this.clearLayers();

    var mapLayer = this.layerMgr.getCurrentLayer()

    this.trains.forEach(function(train){
      //TODO: filter layer

      var marker = self.createMarker(train);
      marker.addTo(self);
      self.currentObjects[train.id] = marker;
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
