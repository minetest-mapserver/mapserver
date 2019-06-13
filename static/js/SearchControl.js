/* exported SearchControl */
/* globals SearchInput: true */
/* globals SearchMenu: true */

var SearchControl = L.Control.extend({
    initialize: function(wsChannel, opts) {
        L.Control.prototype.initialize.call(this, opts);
    },

    onAdd: function(map) {
      var div = L.DomUtil.create('div');
      m.mount(div, SearchInput);
      m.mount(document.getElementById("search-content"), {
        view: function () {
          return m(SearchMenu, {map: map});
        }
      });

      return div;
    }
});
