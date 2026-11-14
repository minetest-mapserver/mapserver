import L from 'leaflet';

export default L.Control.extend({
    onAdd: function(map) {
      var div = L.DomUtil.create('div', 'leaflet-bar leaflet-custom-display');

      var hoverCoord, clickCoord;

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
          html = html + "X=" + Math.floor(hoverCoord.lng) + " Z=" + Math.floor(hoverCoord.lat);

        if (clickCoord)
          html = html + " (marked: X=" + Math.floor(clickCoord.lng) + " Z=" + Math.floor(clickCoord.lat) + ")";

        div.innerHTML = html;
      }

      map.on('mousemove', updateHover);
      map.on('click', updateClick);
      map.on('touch', updateClick);

      return div;
    }
});
