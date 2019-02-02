'use strict';

var TravelnetIcon = L.icon({
  iconUrl: 'pics/travelnet_inv.png',

  iconSize:     [64, 64],
  iconAnchor:   [32, 32],
  popupAnchor:  [0, -32]
});

var TravelnetOverlay = L.LayerGroup.extend({
  initialize: function(wsChannel, layerMgr) {
    L.LayerGroup.prototype.initialize.call(this);

    this.layerMgr = layerMgr;
    this.wsChannel = wsChannel;

    this.currentLayers = [];

    this.onLayerChange = this.onLayerChange.bind(this);
    this.onMapMove = debounce(this.onMapMove.bind(this), 50);
  },

  onLayerChange: function(layer){
    this.reDraw(true);
  },

  onMapMove: function(){
    this.reDraw(false);
  },

  reDraw: function(full){
    var self = this;

    if (full){
      this.clearLayers();
      this.currentLayers = [];
    }

    //TODO: get coords
    api.getMapObjects(-10,-10,-10,10,10,10,"travelnet")
    .then(function(travelnets){
      //TODO: remove non-existing markers, add new ones
      if (!full){
        self.clearLayers();
      }

      travelnets.forEach(function(travelnet){

        console.log(travelnet);

        var popup = "<h4>" + travelnet.attributes.station_name + "</h4><hr>" +
          "<b>X: </b> " + travelnet.x + "<br>" +
          "<b>Y: </b> " + travelnet.y + "<br>" +
          "<b>Z: </b> " + travelnet.z + "<br>" +
          "<b>Network: </b> " + travelnet.attributes.station_network + "<br>" +
          "<b>Owner: </b> " + travelnet.attributes.owner + "<br>";

        var marker = L.marker([travelnet.z, travelnet.x], {icon: TravelnetIcon});
        marker.bindPopup(popup).addTo(self);
      });
    })

  },

  onAdd: function(map) {
    map.on("zoomend", this.onMapMove);
    map.on("moveend", this.onMapMove);
    this.layerMgr.addListener(this.onLayerChange);
    this.reDraw(true)
  },

  onRemove: function(map) {
    map.off("zoomend", this.onMapMove);
    map.off("moveend", this.onMapMove);
    this.layerMgr.removeListener(this.onLayerChange);
  }
});
