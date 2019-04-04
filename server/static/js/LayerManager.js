/* exported LayerManager */

function LayerManager(layers, map){
  this.listeners = [];
  this.currentLayer = layers[0];
  this.layers = layers;

  var self = this;

  map.on('baselayerchange', function (e) {
      self.setLayerId(e.layer.layerId);
  });

}

LayerManager.prototype.setLayerId = function(layerId){
  var self = this;
  this.layers.forEach(function(layer){
    if (layer.id == layerId){
      self.currentLayer = layer;
      self.listeners.forEach(function(listener){
        listener(layer);
      });
      return;
    }
  });
};


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
