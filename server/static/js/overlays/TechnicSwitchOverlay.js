import AbstractIconOverlay from './AbstractIconOverlay.js';

var TechnicSwitchIcon = L.icon({
  iconUrl: 'pics/technic_water_mill_top_active.png',

  iconSize:     [16, 16],
  iconAnchor:   [8, 8],
  popupAnchor:  [0, -16]
});

export default AbstractIconOverlay.extend({
  initialize: function(wsChannel, layerMgr) {
    AbstractIconOverlay.prototype.initialize.call(this, wsChannel, layerMgr, "technicswitch", TechnicSwitchIcon);
  },

  createPopup: function(sw){
    return "<p>Active: " + sw.attributes.active + "</p>" +
      "<p>Channel: " + sw.attributes.channel + "</p>" +
      "<p>Supply: " + sw.attributes.supply + "</p>" +
      "<p>Demand: " + sw.attributes.demand + "</p>";
  }
});
