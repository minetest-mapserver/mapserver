<script setup>
import { ref, watch } from 'vue';
import { useRouter } from 'vue-router';
import { getMapObjects } from '../api.js';
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
  um_area_forsale: "description"
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

async function search(query){
  busy.value = true;
  error.value = null;
  result.value = [];

  try {
    const results = await Promise.all(
      Object.entries(searchFields).map(([type, key]) => searchFor(type, key, query))
    );
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
