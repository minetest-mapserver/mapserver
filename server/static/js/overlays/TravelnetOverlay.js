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

    this.currentObjects = [];

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

    if (this._map.getZoom() < 10) {
      this.clearLayers();
      this.currentObjects = [];
      return;
    }

    if (full){
      this.clearLayers();
      this.currentObjects = [];
    }

    var mapLayer = this.layerMgr.getCurrentLayer()
    var min = this._map.getBounds().getSouthWest();
    var max = this._map.getBounds().getNorthEast();

    var y1 = parseInt(mapLayer.from/16);
    var y2 = parseInt(mapLayer.to/16);
    var x1 = parseInt(min.lng);
    var x2 = parseInt(max.lng);
    var z1 = parseInt(min.lat);
    var z2 = parseInt(max.lat);

    //TODO: get coords
    api.getMapObjects(
      x1, y1, y1,
      x2, y2, z2,
      "travelnet")
    .then(function(travelnets){
      //TODO: remove non-existing markers, add new ones
      if (!full){
        self.clearLayers();
      }

      travelnets.forEach(function(travelnet){
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
    this.clearLayers();
    map.off("zoomend", this.onMapMove);
    map.off("moveend", this.onMapMove);
    this.layerMgr.removeListener(this.onLayerChange);
  }
});
