import layerManager from '../map/LayerManager.js';
import { createMap } from '../map/MapFactory.js';

export default {
  view(){
    return m("div", { class: "full-screen" });
  },

  oncreate(vnode){

    const map = createMap(
      vnode.dom,
      +vnode.attrs.layerId,
      +vnode.attrs.zoom,
      +vnode.attrs.lat,
      +vnode.attrs.lon
    );

    vnode.state.map = map;

    function updateHash(){
      const center = map.getCenter();
      const layerId = layerManager.getCurrentLayer().id;

      m.route.set(`/map/${layerId}/${map.getZoom()}/${center.lng}/${center.lat}`);
    }

    map.on('zoomend', updateHash);
    map.on('moveend', updateHash);
    map.on('baselayerchange', updateHash);
  },

  onbeforeupdate(newVnode) {
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
};
