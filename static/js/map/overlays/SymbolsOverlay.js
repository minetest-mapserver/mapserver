import AbstractIconOverlay from './AbstractIconOverlay.js';

var SymbolIcon = L.icon({
  iconUrl: 'pics/symbols/ign_symbol_star_red.png',

  iconSize:     [16, 16],
  iconAnchor:   [8, 8],
//	tooltipAnchor:[0, -16]
});

export default AbstractIconOverlay.extend({
  initialize: function(wsChannel, layerMgr) {
    AbstractIconOverlay.prototype.initialize.call(this, wsChannel, layerMgr, "symbol", SymbolIcon);

    this.tooltipOptions = {permanent: true, direction: 'top', className: 'tooltip-textonly'};
  },

	getIcon: function(/*ob*/){
    return this.icon;
  },
/*
	var marker = new L.marker([39.5, -77.3], { opacity: 0.01 }); //opacity may be set to zero
	marker.bindTooltip("My Label", {permanent: true, className: "my-label", offset: [0, 0] });
	marker.addTo(map);
*/
  createTooltip: function(symbol){
    return symbol.attributes.minimap_text;
  }
});
