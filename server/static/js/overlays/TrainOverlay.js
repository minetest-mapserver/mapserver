'use strict';

function getTrainImageUrlForType(type){
  switch(type){
    case "advtrains:subway_wagon":
      return "pics/advtrains/advtrains_subway_wagon_inv.png";
    case "advtrains:engine_japan":
      return  "pics/advtrains/advtrains_engine_japan_inv.png";
    case "advtrains:engine_steam":
      return  "pics/advtrains/advtrains_engine_steam_inv.png";
    case "advtrains:engine_industrial":
      return  "pics/advtrains/advtrains_engine_industrial_inv.png";
    case "advtrains:wagon_wood":
      return  "pics/advtrains/advtrains_wagon_wood_inv.png";
    case "advtrains:wagon_box":
      return  "pics/advtrains/advtrains_wagon_box_inv.png";
    default:
      //TODO: fallback image
      return "pics/advtrains/advtrains_subway_wagon_inv.png";
  }
}

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

    //search for wagin in front (whatever "front" is...)
    var type;
    var lowest_pos = 100;
    train.wagons.forEach(function(w){
      if (w.pos_in_train < lowest_pos){
        lowest_pos = w.pos_in_train;
        type = w.type;
      }
    });

    var Icon = L.icon({
      iconUrl: getTrainImageUrlForType(type),

      iconSize:     [16, 16],
      iconAnchor:   [8, 8],
      popupAnchor:  [0, -16]
    });

    var marker = L.marker([train.pos.z, train.pos.x], {icon: Icon});

    var html = "<b>Train</b><hr>";

    html += "<b>Name:</b> " + train.text_outside + "<br>";
    html += "<b>Line:</b> " + train.line + "<br>";
    html += "<b>Velocity:</b> "+ Math.floor(train.velocity*10)/10 + "<br>";

    html += "<b>Composition: </b>";
    train.wagons.forEach(function(w){
      var iconUrl =  getTrainImageUrlForType(w.type);
      html += "<img src='"+iconUrl+"'>";
    });

    marker.bindPopup(html);

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
