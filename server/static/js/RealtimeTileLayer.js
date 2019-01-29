var RealtimeTileLayer = (function(){
  'use strict';

  function getTileSource(layerId, x,y,zoom,cacheBust){
      return "api/tile/" + layerId + "/" + x + "/" + y + "/" + zoom + "?_=" + Date.now();
  }

  function getImageId(layerId, x, y, zoom){
      return "tile-" + layerId + "/" + x + "/" + y + "/" + zoom;
  }



  return {
    init: function(){
      WebSocketChannel.addListener("rendered-tile", function(tc){
        var id = getImageId(tc.layerid, tc.x, tc.y, tc.zoom);
        var el = document.getElementById(id);

        if (el){
            //Update src attribute if img found
            el.src = getTileSource(tc.layerid, tc.x, tc.y, tc.zoom, true);
        }
      });
    },
    create: function(layerId){
        return L.TileLayer.extend({
          createTile: function(coords){
            var tile = document.createElement('img');
            tile.src = getTileSource(layerId, coords.x, coords.y, coords.z);
            tile.id = getImageId(layerId, coords.x, coords.y, coords.z);
            return tile;
          }
        });
    }
  };


}())
