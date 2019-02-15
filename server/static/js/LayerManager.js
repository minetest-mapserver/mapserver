'use strict';

function LayerManager(layers, map){
  this.listeners = [];
  this.currentLayer = layers[0];
  this.layers = layer;

  map.on('baselayerchange', function (e) {
      console.log("baselayerchange", e.layer);
      //TODO
  });

}

LayerManager.prototype.setLayerId = function(layerId){
  var self = this;
  this.layers.forEach(function(layer){
    if (layer.id == layerId){
      self.currentLayer = layer;
    }
  });
},


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
