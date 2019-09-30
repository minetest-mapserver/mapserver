import layerManager from '../LayerManager.js';
import { createMap } from '../map/MapFactory.js';

function setupMap(vnode){
  const map = createMap(
    vnode.dom,
    layerManager.getCurrentLayer().id,
    +vnode.attrs.zoom,
    +vnode.attrs.lat,
    +vnode.attrs.lon
  );

  vnode.state.map = map;

  function updateHash(){
    const center = map.getCenter();
    const layerId = layerManager.getCurrentLayer().id;

    m.route.set(`/map/${layerId}/${map.getZoom()}/` +
      `${Math.floor(center.lng)}/${Math.floor(center.lat)}`);
  }

  map.on('zoomend', updateHash);
  map.on('moveend', updateHash);
}

export default {
  view(){
    return m("div", { class: "full-screen" });
  },

  oncreate(vnode){
    setupMap(vnode);
  },

  onbeforeupdate(newVnode) {
    const center = newVnode.state.map.getCenter();
    const newAattrs = newVnode.attrs;

    return newAattrs.layerId != layerManager.getCurrentLayer().id ||
      newAattrs.zoom != newVnode.state.map.getZoom() ||
      Math.abs(newAattrs.lat - center.lat) > 2 ||
      Math.abs(newAattrs.lat - center.lat) > 2;
  },

  onupdate(vnode){
    if (vnode.attrs.layerId != layerManager.getCurrentLayer().id){
      //layer changed, recreate map
      vnode.state.map.remove();
      layerManager.setLayerId(vnode.attrs.layerId);
      setupMap(vnode);

    } else {
      //position/zoom change
      vnode.state.map.setView([+vnode.attrs.lat, +vnode.attrs.lon], +vnode.attrs.zoom);

    }
  },

  onremove(vnode){
    vnode.state.map.remove();
  }
};
