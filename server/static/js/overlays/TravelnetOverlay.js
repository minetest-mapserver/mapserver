var TravelnetOverlay = (function(){
  'use strict';

  var TravelnetIcon = L.icon({
		iconUrl: 'pics/travelnet_inv.png',

		iconSize:     [64, 64],
		iconAnchor:   [32, 32],
		popupAnchor:  [0, -32]
	});

  return L.LayerGroup.extend({
    onAdd: function(map) {
      console.log("TravelnetOverlay.onAdd");
    },

    onRemove: function(map) {
      console.log("TravelnetOverlay.onRemove");
    }
  });

}());
