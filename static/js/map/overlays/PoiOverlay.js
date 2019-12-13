import AbstractIconOverlay from './AbstractIconOverlay.js';

export default AbstractIconOverlay.extend({
  initialize: function() {
    AbstractIconOverlay.prototype.initialize.call(this, "poi");
  },

  getIcon: function(obj){
    return L.AwesomeMarkers.icon({
      icon: obj.attributes.icon || "home",
      prefix: "fa",
      markerColor: obj.attributes.color || "blue"
    });
  },

  getMaxDisplayedZoom: function(){
    return 5;
  },

  createPopup: function(poi){
    return "<h4>" + poi.attributes.name + "</h4><hr>" +
      "<b>Owner: </b> " + poi.attributes.owner + "<br>";
  }
});
