import AbstractIconOverlay from './AbstractIconOverlay.js';

const IconWood = L.icon({
  iconUrl: 'pics/default_sign_wood.png',
  iconSize:     [16, 16],
  iconAnchor:   [8, 8],
  popupAnchor:  [0, -8]
});

const IconSteel = L.icon({
  iconUrl: 'pics/default_sign_steel.png',
  iconSize:     [16, 16],
  iconAnchor:   [8, 8],
  popupAnchor:  [0, -8]
});

export default AbstractIconOverlay.extend({
  initialize: function() {
    AbstractIconOverlay.prototype.initialize.call(this, "sign");
  },

  getIcon: function(obj) {
    if (obj.attributes.material === "steel") {
      return IconSteel;
    } else {
      return IconWood;
    }
  },

  createPopup: function(sign) {
    return sign.attributes.display_text;
  }
});
