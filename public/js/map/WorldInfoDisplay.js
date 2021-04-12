/* exported WorldInfoDisplay */
import WorldStats from '../components/WorldStats.js';

// coord display
export default L.Control.extend({
    initialize: function(wsChannel, opts) {
        L.Control.prototype.initialize.call(this, opts);
        this.wsChannel = wsChannel;
    },

    onAdd: function() {
      var div = L.DomUtil.create('div', 'leaflet-bar leaflet-custom-display');

      this.wsChannel.addListener("minetest-info", function(info){
        m.render(div, WorldStats(info));
      });

      return div;
    }
});
