
import { getConfig } from './api.js';
import routes from './routes.js';
import wsChannel from './WebSocketChannel.js';
import config from './config.js';
import { hashCompat } from './compat.js';
import layerManager from './LayerManager.js';

// hash route compat
hashCompat();

getConfig().then(cfg => {

  layerManager.setup(cfg.layers);
  config.set(cfg);
  wsChannel.connect();
  m.route(document.getElementById("app"), "/map/0/12/0/0", routes);
});
