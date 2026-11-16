import { createRouter, createWebHashHistory } from 'vue-router';
import MapView from './views/MapView.vue';

// created after the hash-compat rewrite, the history reads the initial location on creation
export function createAppRouter(){
  return createRouter({
    history: createWebHashHistory(),
    routes: [
      { path: "/map/:layerId/:zoom/:lon/:lat", component: MapView, props: true },
      { path: "/search/:query", component: () => import('./views/SearchView.vue'), props: true },
      // camera position and the point it looks at, in node coordinates
      { path: "/view/:pos_x/:pos_y/:pos_z/:target_x/:target_y/:target_z", component: () => import('./views/WorldView.vue'), props: route => ({
        posX: route.params.pos_x, posY: route.params.pos_y, posZ: route.params.pos_z,
        targetX: route.params.target_x, targetY: route.params.target_y, targetZ: route.params.target_z
      }) },
      // position only: look along +z
      { path: "/view/:pos_x/:pos_y/:pos_z", redirect: to => {
        const { pos_x, pos_y, pos_z } = to.params;
        return `/view/${pos_x}/${pos_y}/${pos_z}/${pos_x}/${pos_y}/${+pos_z + 10}`;
      } },
      { path: "/:pathMatch(.*)*", redirect: "/map/0/12/0/0" }
    ]
  });
}
