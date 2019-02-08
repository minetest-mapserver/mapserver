'use strict';

var TechnicAnchorIcon = L.icon({
  iconUrl: 'pics/technic_admin_anchor.png',

  iconSize:     [32, 32],
  iconAnchor:   [16, 16],
  popupAnchor:  [0, -32]
});

var TechnicAnchorOverlay = AbstractIconOverlay.extend({
  initialize: function(wsChannel, layerMgr) {
    AbstractIconOverlay.prototype.initialize.call(this, wsChannel, layerMgr, "luacontroller", TechnicAnchorIcon);
  },

  createPopup: function(lcd){
    return "<p>Owner: " + lcd.attributes.owner + "</p>" +
      "<p>Radius: " + lcd.attributes.radius + "</p>" +
      "<p>Locked: " + lcd.attributes.locked + "</p>" +
      "<p>Enabled: " + lcd.attributes.enabled + "</p>";
  }
});
