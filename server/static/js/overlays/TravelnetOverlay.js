'use strict';

var TravelnetIcon = L.icon({
  iconUrl: 'pics/travelnet_inv.png',

  iconSize:     [64, 64],
  iconAnchor:   [32, 32],
  popupAnchor:  [0, -32]
});

var TravelnetOverlay = L.LayerGroup.extend({
  initialize: function(wsChannel, layerMgr) {
    this.layerMgr = layerMgr;
    this.wsChannel = wsChannel;

    this.onLayerChange = this.onLayerChange.bind(this);
    this.onMapMove = this.onMapMove.bind(this);
  },

  onLayerChange: function(layer){
    this.reDraw(true);
  },

  onMapMove: function(){
    this.reDraw(false);
  },

  reDraw: function(full){
    var self = this;

    if (full)
      this.clearLayers();

    //TODO: get coords
    api.getMapObjects(-10,-10,-10,10,10,10,"travelnet")
    .then(function(travelnets){
      //TODO: remove non-existing markers, add new ones
      if (!full)
        this.clearLayers();

      console.log(travelnets);

      //TODO: attributes, coords
      var marker = L.marker([travelnet.z, travelnet.x], {icon: TravelnetIcon});

      travelnets.forEach(function(travelnet){
        var popup = "<h4>" + travelnet.name + "</h4><hr>" +
          "<b>X: </b> " + travelnet.x + "<br>" +
          "<b>Y: </b> " + travelnet.y + "<br>" +
          "<b>Z: </b> " + travelnet.z + "<br>" +
          "<b>Network: </b> " + travelnet.network + "<br>" +
          "<b>Owner: </b> " + travelnet.owner + "<br>";

        marker.bindPopup(popup).addTo(self);
      });
    })

  },

  onAdd: function(map) {
    this.layerMgr.addListener(this.onLayerChange);
    console.log("TravelnetOverlay.onAdd", map);
  },

  onRemove: function(map) {
    this.layerMgr.removeListener(this.onLayerChange);
    console.log("TravelnetOverlay.onRemove");
  }
});
