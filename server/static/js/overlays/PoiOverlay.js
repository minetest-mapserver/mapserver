/* exported PoiOverlay */
/* globals AbstractIconOverlay: true */

var PoiOverlay = AbstractIconOverlay.extend({
  initialize: function(wsChannel, layerMgr) {
    AbstractIconOverlay.prototype.initialize.call(this, wsChannel, layerMgr, "poi", PoiIcon);
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
