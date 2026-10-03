<script>
import { ref } from 'vue';

// keep the query between map re-creations
const query = ref("");
</script>

<script setup>
import { useRouter } from 'vue-router';

const router = useRouter();

function doSearch(){
  if (query.value.trim() == ""){
    return;
  }
  router.push(`/search/${encodeURIComponent(query.value)}`);
}
</script>

<template>
  <div class="input-group mb-3">
    <input
      v-model="query"
      type="text"
      placeholder="Search"
      class="form-control"
      @keydown.enter="doSearch"
    >
    <div
      class="input-group-append"
      @click="doSearch"
    >
      <span class="input-group-text">
        <i class="fa fa-search" />
      </span>
    </div>
  </div>
</template>
