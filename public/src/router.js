import { createRouter, createWebHashHistory } from 'vue-router';
import MapView from './views/MapView.vue';

// created after the hash-compat rewrite, the history reads the initial location on creation
export function createAppRouter(){
  return createRouter({
    history: createWebHashHistory(),
    routes: [
      { path: "/map/:layerId/:zoom/:lon/:lat", component: MapView, props: true },
      { path: "/search/:query", component: () => import('./views/SearchView.vue'), props: true },
      { path: "/:pathMatch(.*)*", redirect: "/map/0/12/0/0" }
    ]
  });
}
