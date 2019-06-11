import AbstractIconOverlay from './AbstractIconOverlay.js';

var LcdIcon = L.icon({
  iconUrl: 'pics/lcd_lcd.png',

  iconSize:     [16, 16],
  iconAnchor:   [8, 8],
  popupAnchor:  [0, -16]
});

export default AbstractIconOverlay.extend({
  initialize: function(wsChannel, layerMgr) {
    AbstractIconOverlay.prototype.initialize.call(this, wsChannel, layerMgr, "digilinelcd", LcdIcon);
  },

  createPopup: function(lcd){
    return "<pre>" + lcd.attributes.text + "</pre>";
  }
});
