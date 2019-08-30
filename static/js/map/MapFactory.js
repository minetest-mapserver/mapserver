import wsChannel from '../WebSocketChannel.js';
import SimpleCRS from './SimpleCRS.js';
import CoordinatesDisplay from './CoordinatesDisplay.js';
import WorldInfoDisplay from './WorldInfoDisplay.js';
import SearchControl from './SearchControl.js';
import Overlaysetup from './Overlaysetup.js';
import CustomOverlay from './CustomOverlay.js';
import layerManager from './LayerManager.js';
import config from '../config.js';


export function createMap(node, layerId, zoom, lat, lon){

  const cfg = config.get();

  const map = L.map(node, {
    minZoom: 2,
    maxZoom: 12,
    center: [lat, lon],
    zoom: zoom,
    crs: SimpleCRS
  });

  map.attributionControl.addAttribution('<a href="https://github.com/minetest-tools/mapserver">Minetest Mapserver</a>');

  var overlays = {};

  layerManager.setup(wsChannel, cfg.layers, map, layerId);

  //All overlays
  Overlaysetup(cfg, map, overlays, wsChannel, layerManager);
  CustomOverlay(map, overlays);

  new CoordinatesDisplay({ position: 'bottomleft' }).addTo(map);
  new WorldInfoDisplay(wsChannel, { position: 'bottomright' }).addTo(map);

  if (cfg.enablesearch){
    new SearchControl(wsChannel, { position: 'topright' }).addTo(map);
  }

  //layer control
  L.control.layers(layerManager.layerObjects, overlays, { position: "topright" }).addTo(map);

  return map;
}
