import { shallowRef } from 'vue';

const config = shallowRef(null);

export default {
  get(){
    return config.value;
  },

  set(cfg){
    config.value = cfg;
  }
};
