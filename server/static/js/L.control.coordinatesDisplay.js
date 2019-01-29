(function(){
  // coord display

  L.Control.CoordinatesDisplay = L.Control.extend({
      onAdd: function(map) {
        var div = L.DomUtil.create('div', 'leaflet-bar leaflet-custom-display');
        function update(ev){
          var latlng = ev.latlng;
          div.innerHTML = "X:" + parseInt(latlng.lng) + " Z:" + parseInt(latlng.lat);
        }

        map.on('mousemove', update);
        map.on('click', update);
        map.on('touch', update);

        return div;
      },

      onRemove: function(map) {
      }
  });

  L.control.coordinatesDisplay = function(opts) {
      return new L.Control.CoordinatesDisplay(opts);
  }

})()
