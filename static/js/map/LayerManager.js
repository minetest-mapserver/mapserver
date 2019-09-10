import RealtimeTileLayer from './RealtimeTileLayer.js';

class LayerManager {

  setup(layers){
    this.listeners = [];
    this.layers = layers;
    this.currentLayer = this.layers[0];
  }

  setupMap(wsChannel, map, currentLayerId){
    this.map = map;
    this.layerObjects = {};

    var self = this;

    //All layers
    this.layers.forEach(function(layer){
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

  switchLayer(layerId){
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
  }

  setLayerId(layerId){
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
  }

  getLayerByY(y){
    return this.layers.find(layer => (y >= (layer.from*16) && y <= (layer.to*16)));
  }

  addListener(listener){
    this.listeners.push(listener);
  }

  removeListener(listener){
    this.listeners = this.listeners.filter(function(el){
      return el != listener;
    });
  }

  getCurrentLayer(){
    return this.currentLayer;
  }
}

export default new LayerManager();
