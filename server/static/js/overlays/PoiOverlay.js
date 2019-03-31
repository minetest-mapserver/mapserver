'use strict';

var PoiIcon = L.icon({
  iconUrl:       'css/images/marker-icon.png',
	shadowUrl:     'css/images/marker-shadow.png',
	iconSize:    [25, 41],
	iconAnchor:  [12, 41],
	popupAnchor: [1, -34],
	tooltipAnchor: [16, -28],
  shadowSize: [41, 41]
});


var PoiOverlay = AbstractIconOverlay.extend({
  initialize: function(wsChannel, layerMgr) {
    AbstractIconOverlay.prototype.initialize.call(this, wsChannel, layerMgr, "poi", PoiIcon);
  },

  getMaxDisplayedZoom: function(){
    return 5;
  },

  createPopup: function(poi){
    return "<h4>" + poi.attributes.name + "</h4><hr>" +
      "<b>Owner: </b> " + poi.attributes.owner + "<br>";
  }
});
