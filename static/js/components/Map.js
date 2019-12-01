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

  return map;
}

export default {
  view(){
    return m("div", { class: "full-screen" });
  },

  oncreate(vnode){
    this.map = setupMap(vnode);
  },

  onupdate(vnode){
    if (vnode.attrs.layerId != layerManager.getCurrentLayer().id){
      //layer changed, recreate map
      this.map.remove();
      layerManager.setLayerId(vnode.attrs.layerId);
      this.map = setupMap(vnode);

    } else {
      //position/zoom change
      //this.map.setView([+vnode.attrs.lat, +vnode.attrs.lon], +vnode.attrs.zoom);

    }
    return false;
  },

  onremove(){
    this.map.remove();
  }
};
