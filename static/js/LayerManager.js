import RealtimeTileLayer from './map/RealtimeTileLayer.js';

class LayerManager {

  setup(layers){
    this.layers = layers;
    this.currentLayer = this.layers[0];
  }

  setLayerId(layerId){
    var self = this;
    this.layers.forEach(function(layer){
      if (layer.id == layerId){
        self.currentLayer = layer;
        return;
      }
    });

    if (layerId != this.currentLayer.id){
      // layer not found
      this.currentLayer = this.layers[0];
    }
  }

  getLayerByY(y){
    return this.layers.find(layer => (y >= (layer.from*16) && y <= (layer.to*16)));
  }

  getCurrentLayer(){
    return this.currentLayer;
  }
}

export default new LayerManager();
