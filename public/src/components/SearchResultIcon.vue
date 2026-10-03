<script setup>
import { computed } from 'vue';

const props = defineProps({
  obj: { type: Object, required: true }
});

const images = {
  travelnet: "pics/travelnet_inv.png",
  bones: "pics/bones_top.png",
  label: "pics/mapserver_label.png",
  digiterm: "pics/digiterms_beige_front.png",
  digilinelcd: "pics/lcd_lcd.png",
  um_area_forsale: "pics/um_area_forsale_sign_alpha.png"
};

const image = computed(() => {
  const attributes = props.obj.attributes;

  switch (props.obj.type){
    case "locator":
      if (attributes.level == "2")
        return "pics/locator_beacon_level2.png";
      if (attributes.level == "3")
        return "pics/locator_beacon_level3.png";
      return "pics/locator_beacon_level1.png";

    case "shop":
      return attributes.stock == 0 ? "pics/shop_empty.png" : "pics/shop.png";

    default:
      return images[props.obj.type];
  }
});
</script>

<template>
  <i
    v-if="obj.type == 'train'"
    class="fa fa-subway"
  />
  <div
    v-else-if="obj.type == 'poi'"
    style="position: relative"
    :class="'awesome-marker awesome-marker-icon-' + (obj.attributes.color || 'blue')"
  >
    <i :class="'fa fa-' + (obj.attributes.icon || 'home')" />
  </div>
  <img
    v-else-if="image"
    :src="image"
  >
  <template v-else>
    {{ obj.type }}
  </template>
</template>
