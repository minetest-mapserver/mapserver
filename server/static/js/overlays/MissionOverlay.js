'use strict';

var MissionIcon = L.icon({
  iconUrl: 'pics/mission_32px.png',

  iconSize:     [32, 32],
  iconAnchor:   [16, 16],
  popupAnchor:  [0, -32]
});

var MissionOverlay = AbstractIconOverlay.extend({
  initialize: function(wsChannel, layerMgr) {
    AbstractIconOverlay.prototype.initialize.call(this, wsChannel, layerMgr, "mission", MissionIcon);
  },

  createPopup: function(mission){
    return "<p>Owner: " + mission.attributes.owner + "</p>" +
      "<p>Name: " + mission.attributes.name + "</p>" +
      "<p>Successcount: " + mission.attributes.successcount + "</p>" +
      "<p>Failcount: " + mission.attributes.failcount + "</p>";
  }
});
