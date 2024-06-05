import AbstractIconOverlay from './AbstractIconOverlay.js';

export default AbstractIconOverlay.extend({
  initialize: function() {
    AbstractIconOverlay.prototype.initialize.call(this, "atm");
  },

  getMaxDisplayedZoom: function(){
    return 10;
  },

  getIcon: function(obj){
    var img = "pics/atm_front.png";

    if (obj.attributes.type == "wiretransfer")
      img = "pics/atm_front.png";
    else if (obj.attributes.type == "atm2")
      img = "pics/atm2_front.png";
    else if (obj.attributes.type == "atm3")
      img = "pics/atm3_front.png";

    return L.icon({
      iconUrl: img,
      iconSize:     [16, 16],
      iconAnchor:   [8, 8],
      popupAnchor:  [0, -8]
    });
  },

  createPopup: function(obj){
    var title = "ATM";

    if (obj.attributes.type == "wiretransfer")
      title = "Wiretransfer";

    return "<h4>" + title + "</h4>";
  }
});
