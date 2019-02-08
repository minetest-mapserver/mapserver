'use strict';

var TechnicSwitchIcon = L.icon({
  iconUrl: 'pics/technic_water_mill_top_active.png.png',

  iconSize:     [16, 16],
  iconAnchor:   [8, 8],
  popupAnchor:  [0, -16]
});

var TechnicSwitchOverlay = AbstractIconOverlay.extend({
  initialize: function(wsChannel, layerMgr) {
    AbstractIconOverlay.prototype.initialize.call(this, wsChannel, layerMgr, "technicswitch", TechnicQuarryIcon);
  },

  createPopup: function(sw){
    return "<p>Active: " + sw.attributes.active + "</p>" +
      "<p>Channel: " + sw.attributes.channel + "</p>" +
      "<p>Supply: " + sw.attributes.supply + "</p>" +
      "<p>Demand: " + sw.attributes.demand + "</p>";
  }
});
