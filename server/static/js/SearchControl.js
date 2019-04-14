/* exported SearchControl */

var SearchControl = L.Control.extend({
    initialize: function(wsChannel, opts) {
        L.Control.prototype.initialize.call(this, opts);
    },

    onAdd: function() {
      var div = L.DomUtil.create('div');

      m.mount(div, SearchInput);
      m.mount(document.getElementById("search-content"), SearchMenu);

      return div;
    }
});
