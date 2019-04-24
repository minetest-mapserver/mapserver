/* exported MinecartOverlay */
/* globals AbstractIconOverlay: true */
/* jshint unused: false */

var MinecartOverlay = L.LayerGroup.extend({
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
      this.minecarts = info.minecarts || [];
    }.bind(this));
  },

  createMarker: function(cart){

    var Icon = L.icon({
      iconUrl: "pics/minecart_logo.png",

      iconSize:     [32, 32],
      iconAnchor:   [16, 16],
      popupAnchor:  [0, -32]
    });

    var marker = L.marker([cart.pos.z, cart.pos.x], {icon: Icon});
    var html = "<b>Minecart</b><hr>";

    marker.bindPopup(html);

    return marker;
  },

  isCartInCurrentLayer: function(cart){
    var mapLayer = this.layerMgr.getCurrentLayer();

    return (cart.pos.y >= (mapLayer.from*16) && cart.pos.y <= (mapLayer.to*16));
  },


  onMinetestUpdate: function(info){
    var self = this;

    this.minecarts.forEach(function(cart){
      var isInLayer = self.isCartInCurrentLayer(cart);

      if (!isInLayer){
        if (self.currentObjects[train.id]){
          //train is displayed and not on the layer anymore
          //Remove the marker and reference
          self.currentObjects[train.id].remove();
          delete self.currentObjects[train.id];
        }

        return;
      }

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

    var mapLayer = this.layerMgr.getCurrentLayer();

    this.trains.forEach(function(train){
      if (!self.isTrainInCurrentLayer(train)){
        //not in current layer
        return;
      }

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
