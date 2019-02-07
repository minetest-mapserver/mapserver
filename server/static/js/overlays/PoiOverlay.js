'use strict';

var PoiOverlay = AbstractIconOverlay.extend({
  initialize: function(wsChannel, layerMgr) {
    AbstractIconOverlay.prototype.initialize.call(this, wsChannel, layerMgr, "poi", L.Icon.Default);
  },

  createPopup: function(poi){
    return "<h4>" + poi.attributes.name + "</h4><hr>" +
      "<b>Owner: </b> " + poi.attributes.owner + "<br>";
  }
});
