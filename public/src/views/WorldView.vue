<script setup>
import { onBeforeUnmount, onMounted, ref, watch } from 'vue';
import { useRouter } from 'vue-router';
import layerManager from '../LayerManager.js';
import { createView } from '@buckaroobanzay/luanti-viewer';
import { getViewNodedefs } from '../api.js';
import { getViewMedia } from '../util/ViewMediaLoader.js';
import { createViewChunkLoader } from '../util/ViewChunkLoader.js';
import throttle from '../util/throttle.js';

// camera position and the point it looks at in node coordinates
const props = defineProps({
  posX: { type: String, required: true },
  posY: { type: String, required: true },
  posZ: { type: String, required: true },
  targetX: { type: String, required: true },
  targetY: { type: String, required: true },
  targetZ: { type: String, required: true }
});

function currentRoutePath(){
  return [props.posX, props.posY, props.posZ, props.targetX, props.targetY, props.targetZ].join("/");
}

const router = useRouter();
const viewElement = ref(null);
const error = ref(null);

// node name -> definition, fetched once
let nodedefs;
// element of the current view, the viewer releases its resources once it is removed from the document
let container;

// latest camera position reported by the viewer
let cameraPos = null;

function view2d(){
  const x = Math.floor(cameraPos ? cameraPos.x : +props.posX);
  const y = Math.floor(cameraPos ? cameraPos.y : +props.posY);
  const z = Math.floor(cameraPos ? cameraPos.z : +props.posZ);
  const layer = layerManager.getLayerByY(y) || layerManager.getCurrentLayer();
  // map coordinates: lat = z, lon = x
  router.push(`/map/${layer.id}/12/${x}/${z}`);
}

// last position written to the url, route changes caused by it must not recreate the view
let appliedPos = null;

const updateRoute = throttle(camera => {
  const { position: p, target: t } = camera;
  // nodes for the position, one decimal for the target to keep the viewing direction
  const round1 = v => Math.round(v * 10) / 10;
  appliedPos = [
    Math.floor(p.x), Math.floor(p.y), Math.floor(p.z),
    round1(t.x), round1(t.y), round1(t.z)
  ].join("/");
  router.replace(`/view/${appliedPos}`);
}, 500);

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
  const tx = +props.targetX, ty = +props.targetY, tz = +props.targetZ;
  if (![x, y, z, tx, ty, tz].every(Number.isFinite)){
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

  const el = container = document.createElement("div");
  el.className = "full-screen";
  viewElement.value.appendChild(el);

  cameraPos = null;
  createView(el, {
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
      target: { x: tx, y: ty, z: tz }
    },
    events: {
      camera: camera => {
        cameraPos = camera.position;
        if (container === el){
          updateRoute(camera);
        }
      }
    }
  });
}

onMounted(showView);
onBeforeUnmount(removeView);
watch(props, () => {
  // ignore the changes written by the view itself
  if (currentRoutePath() != appliedPos){
    showView();
  }
});
</script>

<template>
  <div
    ref="viewElement"
    class="full-screen"
  >
    <button
      type="button"
      class="btn btn-secondary view-2d"
      @click="view2d"
    >
      <i class="fa fa-map" />
      View 2D
    </button>
    <div
      v-if="error"
      class="alert alert-danger m-2"
    >
      {{ error }}
    </div>
  </div>
</template>

<style scoped>
.view-2d {
  position: absolute;
  top: 10px;
  right: 10px;
  z-index: 1000;
}
</style>
