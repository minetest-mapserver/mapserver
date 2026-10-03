<script setup>
import { shallowRef, onMounted, onBeforeUnmount, watch } from 'vue';
import { useRouter } from 'vue-router';
import layerManager from '../LayerManager.js';
import config from '../config.js';
import { createMap } from '../map/MapFactory.js';
import SearchInput from '../components/SearchInput.vue';
import LayerSelector from '../components/LayerSelector.vue';
import WorldStats from '../components/WorldStats.vue';

const props = defineProps({
  layerId: { type: String, required: true },
  zoom: { type: String, required: true },
  lon: { type: String, required: true },
  lat: { type: String, required: true }
});

const router = useRouter();
const cfg = config.get();

const mapElement = shallowRef(null);
// leaflet control containers, used as teleport targets
const topRight = shallowRef(null);
const bottomRight = shallowRef(null);

let map;

function updateRoute(){
  const center = map.getCenter();
  const layerId = layerManager.getCurrentLayer().id;

  router.replace(`/map/${layerId}/${map.getZoom()}/` +
    `${Math.floor(center.lng)}/${Math.floor(center.lat)}`);
}

function setupMap(){
  layerManager.setLayerId(props.layerId);

  const result = createMap(
    mapElement.value,
    layerManager.getCurrentLayer().id,
    +props.zoom,
    +props.lat,
    +props.lon
  );

  map = result.map;
  topRight.value = result.topRight;
  bottomRight.value = result.bottomRight;

  map.on('zoomend', updateRoute);
  map.on('moveend', updateRoute);
}

onMounted(setupMap);

onBeforeUnmount(() => {
  map.remove();
});

watch(() => props.layerId, () => {
  //layer changed, recreate map
  map.remove();
  setupMap();
});

watch(() => [props.zoom, props.lat, props.lon], () => {
  // position changed from outside (history navigation, edited url)
  const center = map.getCenter();
  if (+props.zoom != map.getZoom() ||
      +props.lon != Math.floor(center.lng) ||
      +props.lat != Math.floor(center.lat)){
    map.setView([+props.lat, +props.lon], +props.zoom);
  }
});
</script>

<template>
  <div
    ref="mapElement"
    class="full-screen"
  />
  <Teleport
    v-if="topRight"
    :to="topRight"
  >
    <SearchInput v-if="cfg.enablesearch" />
    <LayerSelector />
  </Teleport>
  <Teleport
    v-if="bottomRight"
    :to="bottomRight"
  >
    <WorldStats />
  </Teleport>
</template>
