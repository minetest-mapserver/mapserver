import wsChannel from '../WebSocketChannel.js';
import SimpleCRS from './SimpleCRS.js';
import CoordinatesDisplay from './CoordinatesDisplay.js';
import WorldInfoDisplay from './WorldInfoDisplay.js';
import SearchControl from './SearchControl.js';
import Overlaysetup from './Overlaysetup.js';
import CustomOverlay from './CustomOverlay.js';
import layerManager from './LayerManager.js';
import config from '../config.js';

export default {
  view(vnode){
    return m("div", { class: "full-screen" });
  },

  oncreate(vnode){
    const cfg = config.get();

    var map = L.map(vnode.dom, {
      minZoom: 2,
      maxZoom: 12,
      center: [+vnode.attrs.lat, +vnode.attrs.lon],
      zoom: +vnode.attrs.zoom,
      crs: SimpleCRS
    });

    vnode.state.map = map;

    map.attributionControl.addAttribution('<a href="https://github.com/minetest-tools/mapserver">Minetest Mapserver</a>');

    var overlays = {};

    layerManager.setup(wsChannel, cfg.layers, map, +vnode.attrs.layerId);

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

    function updateHash(){
      const center = map.getCenter();
      const layerId = layerManager.getCurrentLayer().id;

      m.route.set(`/map/${layerId}/${map.getZoom()}/${center.lng}/${center.lat}`);
    }

    map.on('zoomend', updateHash);
    map.on('moveend', updateHash);
    map.on('baselayerchange', updateHash);
  },

  onbeforeupdate(newVnode, oldVnode) {
    const center = newVnode.state.map.getCenter();
    const newAattrs = newVnode.attrs;

    return newAattrs.layerId != layerManager.getCurrentLayer().id ||
      newAattrs.zoom != newVnode.state.map.getZoom() ||
      Math.abs(newAattrs.lat - center.lat) > 0.1 ||
      Math.abs(newAattrs.lat - center.lat) > 0.1;
  },

  onupdate(vnode){
    layerManager.switchLayer(+vnode.attrs.layerId);
    vnode.state.map.setView([+vnode.attrs.lat, +vnode.attrs.lon], +vnode.attrs.zoom);
  },

  onremove(vnode){
    vnode.state.map.remove();
  }
}
