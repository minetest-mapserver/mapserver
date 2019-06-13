/* exported BonesOverlay */
/* globals AbstractIconOverlay: true */

var BonesIcon = L.icon({
  iconUrl: 'pics/bones_top.png',

  iconSize:     [16, 16],
  iconAnchor:   [8, 8],
  popupAnchor:  [0, -16]
});

var BonesOverlay = AbstractIconOverlay.extend({
  initialize: function(wsChannel, layerMgr) {
    AbstractIconOverlay.prototype.initialize.call(this, wsChannel, layerMgr, "bones", BonesIcon);
  },

  createPopup: function(bones){
    return "<h4>" + bones.attributes.owner + "</h4>";
  }
});
