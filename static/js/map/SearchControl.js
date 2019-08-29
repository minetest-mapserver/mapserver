import SearchMenu from '../search/SearchMenu.js';
import SearchInput from '../search/SearchInput.js';

export default L.Control.extend({
    initialize: function(wsChannel, opts) {
        L.Control.prototype.initialize.call(this, opts);
    },

    onAdd: function(map) {
      var div = L.DomUtil.create('div');
      m.mount(div, SearchInput);
      m.mount(document.getElementById("search-content"), {
        view: () => m(SearchMenu, {map: map})
      });

      return div;
    }
});
