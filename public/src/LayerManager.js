import { ref, computed } from 'vue';

const layers = ref([]);
const currentLayerId = ref(null);

const currentLayer = computed(() =>
  layers.value.find(layer => layer.id == currentLayerId.value) || layers.value[0]
);

// reactive layer state, shared between the vue components and the leaflet overlays
class LayerManager {

  get layers(){
    return layers.value;
  }

  get currentLayer(){
    return currentLayer;
  }

  setup(list){
    layers.value = list;
    currentLayerId.value = list.length > 0 ? list[0].id : null;
  }

  setLayerId(layerId){
    const layer = layers.value.find(l => l.id == layerId);
    currentLayerId.value = layer ? layer.id : layers.value[0].id;
  }

  getLayerByY(y){
    return layers.value.find(layer => (y >= (layer.from*16) && y <= (layer.to*16)));
  }

  getCurrentLayer(){
    return currentLayer.value;
  }
}

export default new LayerManager();
