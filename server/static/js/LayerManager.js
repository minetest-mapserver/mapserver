/* exported LayerManager */
/* globals RealtimeTileLayer: true */

function LayerManager(wsChannel, layers, map, currentLayerId){
  this.listeners = [];
  this.currentLayer = layers[0];
  this.layers = layers;
  this.map = map;
  this.layerObjects = {};

  var self = this;

  //All layers
  layers.forEach(function(layer){
    var tileLayer = new RealtimeTileLayer(wsChannel, layer.id, map);
    self.layerObjects[layer.name] = tileLayer;
    if (layer.id == currentLayerId){
      tileLayer.addTo(map);
      self.currentLayer = layer;
    }
  });

  map.on('baselayerchange', function (e) {
      self.setLayerId(e.layer.layerId);
  });

}

LayerManager.prototype.switchLayer = function(layerId){
  var self = this;
  Object.keys(this.layerObjects).forEach(function(key){
    var layerObj = self.layerObjects[key];
    if (self.map.hasLayer(layerObj)){
      self.map.removeLayer(layerObj);
    }
  });

  Object.keys(this.layerObjects).forEach(function(key){
    var layerObj = self.layerObjects[key];
    if (layerObj.layerId == layerId){
      self.map.addLayer(layerObj);
    }
  });
};

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

LayerManager.prototype.getLayerByY = function(y){
  return this.layers.find(function(layer){
    return (y >= (layer.from*16) && y <= (layer.to*16));
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
