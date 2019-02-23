'use strict';

var TechnicQuarryIcon = L.icon({
  iconUrl: 'pics/default_tool_mesepick.png',

  iconSize:     [16, 16],
  iconAnchor:   [8, 8],
  popupAnchor:  [0, -16]
});

var TechnicQuarryOverlay = AbstractIconOverlay.extend({
  initialize: function(wsChannel, layerMgr) {
    AbstractIconOverlay.prototype.initialize.call(this, wsChannel, layerMgr, "technicquarry", TechnicQuarryIcon);
  },

  createPopup: function(quarry){
    return "<p>Owner: " + quarry.attributes.owner + "</p>" +
      "<p>Dug: " + quarry.attributes.dug + "</p>" +
      "<p>Enabled: " + quarry.attributes.enabled + "</p>";
  }
});
