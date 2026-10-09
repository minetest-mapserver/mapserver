<script setup>
defineProps({
  obj: { type: Object, required: true }
});

// mapobject type -> attribute used as description
const textAttributes = {
  travelnet: "station_name",
  bones: "owner",
  label: "text",
  digiterm: "display_text",
  digilinelcd: "text",
  locator: "name",
  poi: "name"
};
</script>

<template>
  <span v-if="obj.type == 'train'">
    <span>{{ obj.attributes.station }}</span>
    {{ " " }}
    <span class="badge text-bg-info">{{ obj.attributes.line }}</span>
  </span>

  <span v-else-if="obj.type == 'shop'">
    Shop, trading
    <span class="badge text-bg-primary">
      {{ obj.attributes.out_count }}x<i class="fa fa-cart-arrow-down" />
    </span>
    <span class="badge text-bg-info">{{ obj.attributes.out_item }}</span>
    for
    <span class="badge text-bg-primary">
      {{ obj.attributes.in_count }}x<i class="fa fa-money-bill" />
    </span>
    <span class="badge text-bg-info">{{ obj.attributes.in_item }}</span>
    Stock:
    <span class="badge text-bg-info">{{ obj.attributes.stock }}</span>
  </span>

  <span v-else-if="obj.type == 'um_area_forsale'">
    {{ obj.attributes.description || "No Description" }}
  </span>

  <span v-else-if="textAttributes[obj.type]">
    {{ obj.attributes[textAttributes[obj.type]] }}
  </span>

  <template v-else>
    {{ obj.type }}
  </template>
</template>
