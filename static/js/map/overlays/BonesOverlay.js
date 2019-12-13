import AbstractIconOverlay from './AbstractIconOverlay.js';

var BonesIcon = L.icon({
  iconUrl: 'pics/bones_top.png',

  iconSize:     [16, 16],
  iconAnchor:   [8, 8],
  popupAnchor:  [0, -16]
});

export default AbstractIconOverlay.extend({
  initialize: function() {
    AbstractIconOverlay.prototype.initialize.call(this, "bones", BonesIcon);
  },

  createPopup: function(bones){
    return "<h4>" + bones.attributes.owner + "</h4>";
  }
});
