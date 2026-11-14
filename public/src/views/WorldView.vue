<script setup>
import { onBeforeUnmount, onMounted, ref, watch } from 'vue';
import { createView } from '@buckaroobanzay/luanti-viewer';
import { getViewNodedefs } from '../api.js';
import { getViewMedia } from '../util/ViewMediaLoader.js';
import { createViewChunkLoader } from '../util/ViewChunkLoader.js';

// camera position in node coordinates
const props = defineProps({
  posX: { type: String, required: true },
  posY: { type: String, required: true },
  posZ: { type: String, required: true }
});

const viewElement = ref(null);
const error = ref(null);

// node name -> definition, fetched once
let nodedefs;
// element of the current view, the viewer releases its resources once it is removed from the document
let container;

function removeView(){
  if (container){
    container.remove();
    container = null;
  }
}

function showView(){
  removeView();
  error.value = null;

  const x = +props.posX, y = +props.posY, z = +props.posZ;
  if (![x, y, z].every(Number.isFinite)){
    error.value = "Invalid position";
    return;
  }

  if (!nodedefs){
    nodedefs = getViewNodedefs();
    nodedefs.catch(e => {
      nodedefs = null;
      error.value = `Could not load the node definitions: ${e.message}`;
    });
  }

  container = document.createElement("div");
  container.className = "full-screen";
  viewElement.value.appendChild(container);

  createView(container, {
    getNodeDefinitions: async name => {
      const def = (await nodedefs)[name];
      return def ? { name, ...def } : null;
    },
    getMapblock: createViewChunkLoader(),
    getMedia: getViewMedia,
    // mapblocks around the camera
    viewRange: 6,
    camera: {
      position: { x, y, z },
      target: { x, y, z: z + 10 }
    }
  });
}

onMounted(showView);
onBeforeUnmount(removeView);
watch(props, showView);
</script>

<template>
  <div
    ref="viewElement"
    class="full-screen"
  >
    <div
      v-if="error"
      class="alert alert-danger m-2"
    >
      {{ error }}
    </div>
  </div>
</template>
