import AbstractIconOverlay from './AbstractIconOverlay.js';

var IconStore = {
  'ign_symbol_star_red.png': L.icon({
    iconUrl: 'pics/symbols/ign_symbol_star_red.png',
    iconSize:     [16, 16],
    iconAnchor:   [8, 8]})
};

export default AbstractIconOverlay.extend({
  initialize: function() {
    AbstractIconOverlay.prototype.initialize.call(this, "symbol", IconStore['ign_symbol_star_red.png']);

    this.tooltipOptions = {permanent: true, direction: 'top', className: 'tooltip-textonly'};
  },

  getIcon: function(symbol){
    if (! IconStore[symbol.attributes.minimap_symbol])
      IconStore[symbol.attributes.minimap_symbol] = L.icon({
        iconUrl: 'pics/symbols/'+symbol.attributes.minimap_symbol,
        iconSize:     [16, 16],
        iconAnchor:   [8, 8],
      });

    return IconStore[symbol.attributes.minimap_symbol];
  },

  createTooltip: function(symbol){
    return symbol.attributes.minimap_text;
  }
});
