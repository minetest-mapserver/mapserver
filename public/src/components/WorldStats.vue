<script setup>
import { computed } from 'vue';
import { useWsEvent } from '../composables/useWsEvent.js';

const info = useWsEvent("minetest-info");

//0 - 24'000
const isNight = computed(() => info.value.time < 5500 || info.value.time > 19000);

const time = computed(() => {
  const hour = Math.floor(info.value.time/1000);
  const minute = Math.floor((info.value.time % 1000) / 1000 * 60);
  return `${hour}:${String(minute).padStart(2, "0")}`;
});

const lagColor = computed(() => {
  if (info.value.max_lag > 1.2)
    return "red";
  if (info.value.max_lag > 0.8)
    return "orange";
  return "green";
});
</script>

<template>
  <div v-if="info">
    <span class="fa fa-users" />{{ info.players ? info.players.length : 0 }}
    <span
      class="fa fa-wifi"
      :style="{ color: lagColor }"
    />{{ Math.floor(info.max_lag*1000) }} ms
    <span class="fa fa-clock" />
    <span
      v-if="isNight"
      class="fa fa-moon"
      style="color: blue;"
    />
    <span
      v-else
      class="fa fa-sun"
      style="color: orange;"
    />{{ time }}
  </div>
</template>
