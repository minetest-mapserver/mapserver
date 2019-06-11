/* exported Hashroute */

export default {

  setup: function(map, layerMgr){
    function updateHash(){
      var center = map.getCenter();
      window.location.hash =
        layerMgr.getCurrentLayer().id + "/" +
        center.lng + "/" + center.lat + "/" + map.getZoom();
    }

    map.on('zoomend', updateHash);
    map.on('moveend', updateHash);
    map.on('baselayerchange', updateHash);
    updateHash();
  },

  getLayerId: function(){
    var hashParts = window.location.hash.substring(1).split("/");
    if (hashParts.length == 4){
      //new format
      return +hashParts[0];

    }

    return 0;
  },

  getZoom: function(){
    var hashParts = window.location.hash.substring(1).split("/");
    if (hashParts.length == 3){
      //old format
      return +hashParts[2];

    } else if (hashParts.length == 4){
      //new format
      return +hashParts[3];

    }

    return 11;
  },

  getCenter: function(){
    var hashParts = window.location.hash.substring(1).split("/");
    if (hashParts.length == 3){
      //old format
      return [+hashParts[1], +hashParts[0]];

    } else if (hashParts.length == 4){
      //new format
      return [+hashParts[2], +hashParts[1]];

    }

    return [0, 0];
  }

};
