import SearchInput from '../components/SearchInput.js';

export default L.Control.extend({
    initialize: function(wsChannel, opts) {
        L.Control.prototype.initialize.call(this, opts);
    },

    onAdd: function(map) {
      var div = L.DomUtil.create('div');
      m.mount(div, SearchInput);
      return div;
    }
});
