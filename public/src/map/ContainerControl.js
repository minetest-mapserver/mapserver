
// empty leaflet control, used as a mount-point (teleport target) for vue components
export default L.Control.extend({
  options: {
    className: ""
  },

  onAdd: function() {
    var div = L.DomUtil.create('div', this.options.className);

    // don't drag/zoom the map while interacting with the control
    L.DomEvent.disableClickPropagation(div);
    L.DomEvent.disableScrollPropagation(div);

    return div;
  }
});
