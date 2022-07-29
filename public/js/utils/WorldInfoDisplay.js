import WorldStats from '../components/WorldStats.js';

export default L.Control.extend({
    initialize: function(wsChannel, opts) {
        L.Control.prototype.initialize.call(this, opts);
        this.wsChannel = wsChannel;
    },

    onAdd: function() {
      var div = L.DomUtil.create('div', 'leaflet-bar leaflet-custom-display');
      const app = Vue.createApp(WorldStats);
      app.mount(div);
      return div;
    }
});
