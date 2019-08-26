
import { getConfig } from './api.js';
//import { setup } from './map.js';
import routes from './routes.js';
import wsChannel from './WebSocketChannel.js';
import config from './config.js';

//TODO: migrate #/layer/zoom/lat/lon to #!/map/...
//TODO: migrate #/zoom/lat/lon to #!/map/...

getConfig().then(cfg => {
  config.set(cfg);
  wsChannel.connect();
  m.route(document.getElementById("app"), "/map/0/13/0/0", routes);
  //setup(cfg);
});
