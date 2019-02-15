'use strict';

var LuacontrollerIcon = L.icon({
  iconUrl: 'pics/jeija_luacontroller_top.png',

  iconSize:     [16, 16], //TODO: 512px :O ...
  iconAnchor:   [8, 8],
  popupAnchor:  [0, -16]
});

var LuacontrollerBurntIcon = L.icon({
  iconUrl: 'pics/jeija_luacontroller_burnt_top.png',

  iconSize:     [16, 16], //TODO: 512px :O ...
  iconAnchor:   [8, 8],
  popupAnchor:  [0, -16]
});

var LuacontrollerOverlay = AbstractIconOverlay.extend({
  initialize: function(wsChannel, layerMgr) {
    AbstractIconOverlay.prototype.initialize.call(this, wsChannel, layerMgr, "luacontroller");
  },

  getIcon: function(ctrl){
    if (ctrl.burnt)
      return LuacontrollerIcon;
    else
      return LuacontrollerBurntIcon;
  },

  createPopup: function(ctrl){
    return "<pre>" + ctrl.attributes.code + "</pre>";
  }
});
