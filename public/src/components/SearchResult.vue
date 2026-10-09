<script setup>
import { useRouter } from 'vue-router';
import layerManager from '../LayerManager.js';
import SearchResultIcon from './SearchResultIcon.vue';
import SearchResultDescription from './SearchResultDescription.vue';

defineProps({
  result: { type: Array, required: true }
});

const router = useRouter();

function getLayerName(obj){
  const layer = layerManager.getLayerByY(obj.y);
  return layer ? layer.name : "<unknown>";
}

function goto(obj){
  const layer = layerManager.getLayerByY(obj.y) || layerManager.getCurrentLayer();
  router.push(`/map/${layer.id}/12/${obj.x}/${obj.z}`);
}
</script>

<template>
  <table class="table table-striped">
    <thead>
      <tr>
        <th>Type</th>
        <th>Owner</th>
        <th>Layer</th>
        <th>Position</th>
        <th>Description</th>
        <th>Action</th>
      </tr>
    </thead>
    <tbody>
      <tr
        v-for="(obj, i) in result"
        :key="i"
        :class="{ 'table-warning': obj.type == 'shop' && obj.attributes.stock == 0 }"
      >
        <td><SearchResultIcon :obj="obj" /></td>
        <td>{{ obj.attributes.owner }}</td>
        <td>{{ getLayerName(obj) }}</td>
        <td>
          <span class="badge text-bg-success">{{ obj.x }}/{{ obj.y }}/{{ obj.z }}</span>
        </td>
        <td><SearchResultDescription :obj="obj" /></td>
        <td>
          <button
            type="button"
            class="btn btn-secondary"
            @click="goto(obj)"
          >
            Goto <i class="fas fa-play" />
          </button>
        </td>
      </tr>
    </tbody>
  </table>
</template>
