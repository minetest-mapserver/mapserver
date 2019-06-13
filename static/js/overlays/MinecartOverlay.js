/* exported MinecartOverlay */
/* globals AbstractIconOverlay: true */
/* jshint unused: false */

var MinecartOverlay = L.LayerGroup.extend({
  initialize: function(wsChannel, layerMgr) {
    L.LayerGroup.prototype.initialize.call(this);

    this.layerMgr = layerMgr;
    this.wsChannel = wsChannel;

    this.currentObjects = {}; // name => marker
    this.minecarts = [];

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
    html += "<b>Id: </b> " + cart.id + "<br>";

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
        if (self.currentObjects[cart.id]){
          //cart is displayed and not on the layer anymore
          //Remove the marker and reference
          self.currentObjects[cart.id].remove();
          delete self.currentObjects[cart.id];
        }

        return;
      }

      if (self.currentObjects[cart.id]){
        //marker exists
        self.currentObjects[cart.id].setLatLng([cart.pos.z, cart.pos.x]);
        //setPopupContent

      } else {
        //marker does not exist
        var marker = self.createMarker(cart);
        marker.addTo(self);

        self.currentObjects[cart.id] = marker;
      }
    });

    Object.keys(self.currentObjects).forEach(function(existingId){
      var cartIsActive = self.minecarts.find(function(t){
        return t.id == existingId;
      });

      if (!cartIsActive){
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

    this.minecarts.forEach(function(cart){
      if (!self.isCartInCurrentLayer(cart)){
        //not in current layer
        return;
      }

      var marker = self.createMarker(cart);
      marker.addTo(self);
      self.currentObjects[cart.id] = marker;
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
