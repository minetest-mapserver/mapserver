import AbstractIconOverlay from './AbstractIconOverlay.js';

var LuacontrollerIcon = L.icon({
  iconUrl: 'pics/jeija_luacontroller_top.png',

  iconSize:     [16, 16],
  iconAnchor:   [8, 8],
  popupAnchor:  [0, -16]
});

var LuacontrollerBurntIcon = L.icon({
  iconUrl: 'pics/jeija_luacontroller_burnt_top.png',

  iconSize:     [16, 16],
  iconAnchor:   [8, 8],
  popupAnchor:  [0, -16]
});

export default AbstractIconOverlay.extend({
  initialize: function(wsChannel, layerMgr) {
    AbstractIconOverlay.prototype.initialize.call(this, wsChannel, layerMgr, "luacontroller");
  },

  getIcon: function(ctrl){
    if (ctrl.attributes.burnt)
      return LuacontrollerBurntIcon;
    else
      return LuacontrollerIcon;
  },

  createPopup: function(ctrl){
    return "LuaController" + (ctrl.attributes.burnt ? "(Burnt)" : "");
  }
});
