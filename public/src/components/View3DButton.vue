<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import layerManager from '../LayerManager.js';
import { getStartHeight } from '../api.js';

const props = defineProps({
  // returns the current leaflet map
  getMap: { type: Function, required: true }
});

const router = useRouter();
const busy = ref(false);

async function view3d(){
  const center = props.getMap().getCenter();
  const x = Math.floor(center.lng);
  const z = Math.floor(center.lat);

  busy.value = true;
  try {
    const result = await getStartHeight(layerManager.getCurrentLayer().id, x, z);
    if (result){
      router.push(`/view/${x}/${result.y + 1}/${z}`);
    }
  } catch (e){
    // fail silently
    console.error(e);
  } finally {
    busy.value = false;
  }
}
</script>

<template>
  <button
    type="button"
    class="btn btn-secondary text-nowrap"
    :disabled="busy"
    @click="view3d"
  >
    <i class="fa fa-cube" />
    View 3D
  </button>
</template>
