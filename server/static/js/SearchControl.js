/* exported SearchControl */

var SearchControl = L.Control.extend({
    initialize: function(wsChannel, opts) {
        L.Control.prototype.initialize.call(this, opts);
    },

    onAdd: function() {
      var div = L.DomUtil.create('div', 'leaflet-bar leaflet-custom-display');

      var View = {
        view: function(){
          return m("input[type=text]", { placeholder: "Search", class: "form-control" });
        }
      };

      m.mount(div, SearchInput);
      m.mount(document.getElementById("search-content"), SearchMenu);

      return div;
    }
});
