'use strict';

var LabelOverlay = AbstractIconOverlay.extend({
  initialize: function(wsChannel, layerMgr) {
    AbstractIconOverlay.prototype.initialize.call(this, wsChannel, layerMgr, "digiterm");
  },

  getIcon: function(lbl){
    return L.divIcon({html: lbl.attributes.text});
  },

  createPopup: function(lbl){
    return "<pre>" + lbl.attributes.text + "</pre>";
  }
});
