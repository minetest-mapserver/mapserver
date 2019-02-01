'use strict';

function LayerManager(layers, map){
  this.listeners = [];
  this.currentLayer = layers[0];

  map.on('baselayerchange', function (e) {
      console.log("baselayerchange", e.layer);
      //TODO
  });

}

LayerManager.prototype.addListener = function(listener){
  this.listeners.push(listener);
};

LayerManager.prototype.removeListener = function(listener){
  this.listeners = this.listeners.filter(function(el){
    return el != listener;
  });
};

LayerManager.prototype.getCurrentLayer = function(){
  return this.currentLayer;
};
