/* exported WorldInfoDisplay */

var worldInfoRender = function(info){
  return "Lag: " + parseInt(info.max_lag*10)/10 + " Time: " + parseInt(info.time)/1000;
}

// coord display
var WorldInfoDisplay = L.Control.extend({
    initialize: function(wsChannel, opts) {
        L.Control.prototype.initialize.call(this, opts);
        this.wsChannel = wsChannel;
    },

    onAdd: function() {
      var div = L.DomUtil.create('div', 'leaflet-bar leaflet-custom-display');

      this.wsChannel.addListener("minetest-info", function(info){
        m.render(div, worldInfoRender(info))
      });

      return div;
    }
});
