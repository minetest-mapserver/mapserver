/* exported LabelOverlay */
/* globals AbstractIconOverlay: true */

var LabelIcon = L.Icon.extend({
  initialize: function(options) {
    L.Icon.prototype.initialize.call(this, options);
  },

  createIcon: function() {
		var div = document.createElement('div'),
		    options = this.options;

		div.innerHTML = options.html || "";

		return div;
  }
});

var LabelOverlay = AbstractIconOverlay.extend({
  initialize: function(wsChannel, layerMgr) {
    AbstractIconOverlay.prototype.initialize.call(this, wsChannel, layerMgr, "label");
  },

  getIcon: function(lbl){
    return new LabelIcon({
      iconAnchor:   [15, 50],
      iconSize:     [30, 100],
      html: "<svg height='30' width='100'><text x='0' y='15'>" + lbl.attributes.text + "</text></svg>"
    });
  },

  createPopup: function(lbl){
    return "<p>" + lbl.attributes.text + "</p>";
  }
});
