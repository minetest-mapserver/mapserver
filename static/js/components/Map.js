import wsChannel from '../WebSocketChannel.js';
import SimpleCRS from '../SimpleCRS.js';
import CoordinatesDisplay from '../CoordinatesDisplay.js';
import WorldInfoDisplay from '../WorldInfoDisplay.js';
import SearchControl from '../SearchControl.js';
import Overlaysetup from '../Overlaysetup.js';
import layerManager from '../LayerManager.js';
import config from '../config.js';

export default {
  view(vnode){
    return m("div", { class: "full-screen" });
  },

  oncreate(vnode){
    console.log("oncreate", vnode);
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

    new CoordinatesDisplay({ position: 'bottomleft' }).addTo(map);
    new WorldInfoDisplay(wsChannel, { position: 'bottomright' }).addTo(map);

    if (cfg.enablesearch){
      new SearchControl(wsChannel, { position: 'topright' }).addTo(map);
    }

    //layer control
    L.control.layers(layerManager.layerObjects, overlays, { position: "topright" }).addTo(map);

    //TODO: overlay persistence (state, localstorage)
    //TODO: update hash
  },

  onbeforeupdate(newVnode, oldVnode) {
      return false;
  },

  onupdate(vnode){
    console.log("onupdate", vnode);

    //TODO: compare and update center,zoom,layer
  },

  onremove(vnode){
    console.log("onremove", vnode);
    vnode.state.map.remove();
  }
}
