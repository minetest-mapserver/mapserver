<script setup>
import { shallowRef, onMounted, onBeforeUnmount, watch } from 'vue';
import { useRouter } from 'vue-router';
import layerManager from '../LayerManager.js';
import config from '../config.js';
import { createMap } from '../map/MapFactory.js';
import SearchInput from '../components/SearchInput.vue';
import LayerSelector from '../components/LayerSelector.vue';
import View3DButton from '../components/View3DButton.vue';
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
// last position written to the route by the map itself
let lastRoute = null;

function updateRoute(){
  const center = map.getCenter();
  const layerId = layerManager.getCurrentLayer().id;

  lastRoute = {
    zoom: map.getZoom(),
    lon: Math.floor(center.lng),
    lat: Math.floor(center.lat)
  };

  router.replace(`/map/${layerId}/${lastRoute.zoom}/${lastRoute.lon}/${lastRoute.lat}`);
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
  // the route update is async: compare against what the map wrote, not the live map state,
  // otherwise a late echo of our own update snaps the map back mid-gesture (pinch/drag)
  if (lastRoute &&
      +props.zoom == lastRoute.zoom &&
      +props.lon == lastRoute.lon &&
      +props.lat == lastRoute.lat){
    return;
  }
  map.setView([+props.lat, +props.lon], +props.zoom);
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
    <div class="d-flex mb-3 gap-2">
      <View3DButton :get-map="() => map" />
      <LayerSelector />
    </div>
  </Teleport>
  <Teleport
    v-if="bottomRight"
    :to="bottomRight"
  >
    <WorldStats />
  </Teleport>
</template>
