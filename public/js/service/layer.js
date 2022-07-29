
import store from '../store/config.js';

export const getLayerById = id => store.layers.find(l => l.id == id);