import 'leaflet/dist/leaflet.css';

import './assets/css/bootstrap.min.css';
import './assets/css/fontawesome.min.css';
import 'leaflet.awesome-markers/dist/leaflet.awesome-markers.css';
import './assets/css/custom.css';

import { createApp } from 'vue';
import App from './App.vue';
import { createAppRouter } from './router.js';
import { getConfig } from './api.js';
import wsChannel from './WebSocketChannel.js';
import config from './config.js';
import layerManager from './LayerManager.js';

getConfig()
.then(cfg => {
  layerManager.setup(cfg.layers);
  config.set(cfg);

  if (cfg.pagename) {
    document.title = cfg.pagename;
  }

  wsChannel.connect();

  createApp(App)
    .use(createAppRouter())
    .mount("#app");
})
.catch(e => {
  document.getElementById("app").textContent = e;
});
