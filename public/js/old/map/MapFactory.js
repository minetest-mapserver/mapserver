import wsChannel from '../WebSocketChannel.js';
import SimpleCRS from './SimpleCRS.js';
import CoordinatesDisplay from './CoordinatesDisplay.js';
import WorldInfoDisplay from './WorldInfoDisplay.js';
import TopRightControl from './TopRightControl.js';
import Overlaysetup from './Overlaysetup.js';
import CustomOverlay from './CustomOverlay.js';
import RealtimeTileLayer from './RealtimeTileLayer.js';

import config from '../config.js';


export function createMap(node, layerId, zoom, lat, lon){

  const cfg = config.get();

  const map = L.map(node, {
    minZoom: 2,
    maxZoom: 12,
    center: [lat, lon],
    zoom: zoom,
    crs: SimpleCRS,
    maxBounds: L.latLngBounds(
      L.latLng(-31000, -31000),
      L.latLng(31000, 31000)
    )
  });

  map.attributionControl.addAttribution('<a href="https://github.com/minetest-mapserver/mapserver">Minetest Mapserver</a>');

  var tileLayer = new RealtimeTileLayer(wsChannel, layerId, map);
  tileLayer.addTo(map);

  //All overlays
  var overlays = {};
  Overlaysetup(cfg, map, overlays);
  CustomOverlay(map, overlays);

  new CoordinatesDisplay({ position: 'bottomleft' }).addTo(map);
  new WorldInfoDisplay(wsChannel, { position: 'bottomright' }).addTo(map);
  new TopRightControl({ position: 'topright' }).addTo(map);

  //layer control
  L.control.layers({}, overlays, { position: "topright" }).addTo(map);

  return map;
}
