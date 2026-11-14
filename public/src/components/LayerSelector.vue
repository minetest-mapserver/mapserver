<script setup>
import { useRoute, useRouter } from 'vue-router';
import layerManager from '../LayerManager.js';

const route = useRoute();
const router = useRouter();
const currentLayer = layerManager.currentLayer;

function onChange(e){
  router.push({
    params: { ...route.params, layerId: e.target.value }
  });
}
</script>

<template>
  <!-- display layer selector only if there is a choice -->
  <select
    v-if="layerManager.layers.length > 1"
    class="form-control"
    :value="currentLayer.id"
    @change="onChange"
  >
    <option
      v-for="layer in layerManager.layers"
      :key="layer.id"
      :value="layer.id"
    >
      {{ layer.name }}
    </option>
  </select>
</template>
