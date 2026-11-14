<script setup>
import { ref, watch } from 'vue';
import { useRouter } from 'vue-router';
import { getMapObjects, getRespawnPlaces } from '../api.js';
import SearchResult from '../components/SearchResult.vue';

const props = defineProps({
  query: { type: String, required: true }
});

// mapobject type -> searched attribute
const searchFields = {
  shop: "out_item",
  poi: "name",
  train: "station",
  travelnet: "station_name",
  bones: "owner",
  locator: "name",
  label: "text",
  digiterm: "display_text",
  digilinelcd: "text",
  um_area_forsale: "description",
  sign: "display_text"
};

const router = useRouter();

function backToMap(){
  if (window.history.state && window.history.state.back){
    router.back();
  } else {
    router.push("/");
  }
}

const busy = ref(false);
const error = ref(null);
const result = ref([]);

function searchFor(type, key, valuelike){
  return getMapObjects({
    pos1: { x:-2048, y:-2048, z:-2048 },
    pos2: { x:2048, y:2048, z:2048 },
    type: type,
    attributelike: {
      key: key,
      value: "%" + valuelike +"%"
    }
  });
}

// respawn places are not mapobjects, filter them here and return them in the mapobject format
async function searchRespawnPlaces(valuelike){
  const needle = valuelike.toLowerCase();
  const places = await getRespawnPlaces();

  return Object.entries(places || {})
    .filter(([key, place]) => {
      return key.toLowerCase().includes(needle) ||
        (place.full_name || "").toLowerCase().includes(needle);
    })
    .map(([key, place]) => {
      const pos = place.pos;
      return {
        mapblock: {
          x: Math.floor(pos.x / 16),
          y: Math.floor(pos.y / 16),
          z: Math.floor(pos.z / 16)
        },
        x: pos.x,
        y: pos.y,
        z: pos.z,
        type: "respawn_place",
        attributes: {
          name: place.full_name || key,
          full_name: place.full_name,
          key: key,
          color: place.color,
          icon: place.icon
        }
      };
    });
}

async function search(query){
  busy.value = true;
  error.value = null;
  result.value = [];

  try {
    const results = await Promise.all([
      ...Object.entries(searchFields).map(([type, key]) => searchFor(type, key, query)),
      searchRespawnPlaces(query)
    ]);
    result.value = results.flatMap(r => r || []);
  } catch (e) {
    error.value = e;
  } finally {
    busy.value = false;
  }
}

watch(() => props.query, search, { immediate: true });
</script>

<template>
  <div class="container-fluid">
    <h5>
      <button
        type="button"
        class="btn btn-secondary"
        @click="backToMap"
      >
        <i class="fa fa-map" /> Map
      </button>
      Search results for "{{ query }}"
    </h5>

    <div v-if="busy">
      <i class="fa fa-spinner fa-spin" /> Searching...
    </div>
    <div
      v-else-if="error"
      class="alert alert-danger"
    >
      {{ error }}
    </div>
    <div v-else-if="result.length == 0">
      Nothing found
    </div>
    <SearchResult
      v-else
      :result="result"
    />
  </div>
</template>
