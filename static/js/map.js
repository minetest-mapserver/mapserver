import wsChannel from './WebSocketChannel.js';
import Hashroute from './Hashroute.js';
import SimpleCRS from './SimpleCRS.js';
import CoordinatesDisplay from './CoordinatesDisplay.js';
import WorldInfoDisplay from './WorldInfoDisplay.js';
import SearchControl from './SearchControl.js';
import Overlaysetup from './Overlaysetup.js';
import layerManager from './LayerManager.js';

export function setup(cfg){

  wsChannel.connect();

  var map = L.map('image-map', {
    minZoom: 2,
    maxZoom: 12,
    center: Hashroute.getCenter(),
    zoom: Hashroute.getZoom(),
    crs: SimpleCRS
  });

  map.attributionControl.addAttribution('<a href="https://github.com/thomasrudin-mt/mapserver">Minetest Mapserver</a>');

  var overlays = {};

  layerManager.setup(wsChannel, cfg.layers, map, Hashroute.getLayerId());

  //All overlays
  Overlaysetup(cfg, map, overlays, wsChannel, layerManager);

  new CoordinatesDisplay({ position: 'bottomleft' }).addTo(map);
  new WorldInfoDisplay(wsChannel, { position: 'bottomright' }).addTo(map);

  if (cfg.enablesearch){
    new SearchControl(wsChannel, { position: 'topright' }).addTo(map);
  }

  //layer control
  L.control.layers(layerManager.layerObjects, overlays, { position: "topright" }).addTo(map);

  Hashroute.setup(map, layerManager);
}
