import AbstractIconOverlay from './AbstractIconOverlay.js';

var TravelnetIcon = L.icon({
  iconUrl: 'pics/travelnet_inv.png',

  iconSize:     [64, 64],
  iconAnchor:   [32, 32],
  popupAnchor:  [0, -32]
});

export default AbstractIconOverlay.extend({
  initialize: function(wsChannel, layerMgr) {
    AbstractIconOverlay.prototype.initialize.call(this, wsChannel, layerMgr, "travelnet", TravelnetIcon);
  },

  createPopup: function(travelnet){
    return "<h4>" + travelnet.attributes.station_name + "</h4><hr>" +
      "<b>X: </b> " + travelnet.x + "<br>" +
      "<b>Y: </b> " + travelnet.y + "<br>" +
      "<b>Z: </b> " + travelnet.z + "<br>" +
      "<b>Network: </b> " + travelnet.attributes.station_network + "<br>" +
      "<b>Owner: </b> " + travelnet.attributes.owner + "<br>";
  }
});
