'use strict';

// coord display
var CoordinatesDisplay = L.Control.extend({
    onAdd: function(map) {
      var div = L.DomUtil.create('div', 'leaflet-bar leaflet-custom-display');

      var hoverCoord, clickCoord

      function updateHover(ev){
        hoverCoord = ev.latlng;
        update();
      }

      function updateClick(ev){
        clickCoord = ev.latlng;
        update();
      }

      function update(){
        var html = "";
        if (hoverCoord)
          html += = "X=" + parseInt(hoverCoord.lng) + " Z=" + parseInt(hoverCoord.lat);

        if (clickCoord)
          html += = " (marked: X=" + parseInt(clickCoord.lng) + " Z=" + parseInt(clickCoord.lat) + ")";

        div.innerHTML = html;
      }

      map.on('mousemove', updateHover);
      map.on('click', updateClick);
      map.on('touch', updateClick);

      return div;
    },

    onRemove: function(map) {
    }
});
