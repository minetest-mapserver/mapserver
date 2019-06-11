import AbstractIconOverlay from './AbstractIconOverlay.js';

var DigitermIcon = L.icon({
  iconUrl: 'pics/digiterms_beige_front.png',

  iconSize:     [16, 16],
  iconAnchor:   [8, 8],
  popupAnchor:  [0, -16]
});

export default AbstractIconOverlay.extend({
  initialize: function(wsChannel, layerMgr) {
    AbstractIconOverlay.prototype.initialize.call(this, wsChannel, layerMgr, "digiterm", DigitermIcon);
  },

  createPopup: function(lcd){
    return "<pre>" + lcd.attributes.display_text + "</pre>";
  }
});
