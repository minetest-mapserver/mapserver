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
    AbstractIconOverlay.prototype.initialize.call(this, wsChannel, layerMgr, "shop");
  },

  getMaxDisplayedZoom: function(){
    return 5;
  },

  getIcon: function(obj){
    if (obj.attributes.stock > 0)
      return ShopIcon;
    else
      return ShopEmptyIcon;
  },

  createPopup: function(obj){
    console.log(obj)
    return "<h4>" + obj.attributes.type + "</h4><hr>" +
      "<b>Owner: </b> " + obj.attributes.owner + "<br>" +
      "<b>Input: </b> " + obj.attributes.in_count + " x " + obj.attributes.in_item + "<br>" +
      "<b>Output: </b> " + obj.attributes.out_count + " x " + obj.attributes.out_item + "<br>" +
      "<b>Stock: </b> " + obj.attributes.stock + "<br>";
  }
});
