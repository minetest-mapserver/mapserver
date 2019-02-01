'use strict';

function LayerManager(layers, map){

  map.on('baselayerchange', function (e) {
      console.log("baselayerchange", e.layer);
  });
  
}

LayerManager.prototype.addListener = function(listener){
};

LayerManager.prototype.getCurrentLayer = function(){
};
