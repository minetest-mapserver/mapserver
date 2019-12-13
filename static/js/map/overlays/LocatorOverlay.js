import AbstractIconOverlay from './AbstractIconOverlay.js';

export default AbstractIconOverlay.extend({
  initialize: function() {
    AbstractIconOverlay.prototype.initialize.call(this, "locator");
  },

  getMaxDisplayedZoom: function(){
    return 6;
  },

  getIcon: function(obj){
    var img = "pics/locator_beacon_level1.png";

    if (obj.attributes.level == "2")
      img = "pics/locator_beacon_level2.png";
    else if (obj.attributes.level == "3")
      img = "pics/locator_beacon_level3.png";

    L.icon({
      iconUrl: img,
      iconSize:     [32, 32],
      iconAnchor:   [16, 16],
      popupAnchor:  [0, -16]
    });
  },

  createPopup: function(obj){
    return "<h4>Locator</h4><hr>" +
      "<b>Owner: " + obj.attributes.owner + "</b><br>" +
      "<b>Name:</b> " + obj.attributes.name + "<br>" +
      "<b>Level:</b> " + obj.attributes.level + "<br>";
  }
});
