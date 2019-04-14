/* exported ShopOverlay */
/* globals AbstractIconOverlay: true */

var ShopIcon = L.icon({
  iconUrl: 'pics/shop.png',
  iconSize:     [32, 32],
  iconAnchor:   [16, 16],
  popupAnchor:  [0, -16]
});

var ShopEmptyIcon = L.icon({
  iconUrl: 'pics/shop_empty.png',
  iconSize:     [32, 32],
  iconAnchor:   [16, 16],
  popupAnchor:  [0, -16]
});


var ShopOverlay = AbstractIconOverlay.extend({
  initialize: function(wsChannel, layerMgr) {
    AbstractIconOverlay.prototype.initialize.call(this, wsChannel, layerMgr, "shop", ShopIcon);
  },

  getMaxDisplayedZoom: function(){
    return 5;
  },

  createPopup: function(poi){
    return "<h4>" + poi.attributes.type + "</h4><hr>" +
      "<b>Owner: </b> " + poi.attributes.owner + "<br>";
  }
});
