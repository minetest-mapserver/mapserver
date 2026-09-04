import debounce from '../../util/debounce.js';
import layerMgr from '../../LayerManager.js';

function escapeHtml(value) {
  return String(value)
    .replaceAll("&", "&amp;")
    .replaceAll("<", "&lt;")
    .replaceAll(">", "&gt;")
    .replaceAll('"', "&quot;")
    .replaceAll("'", "&#039;");
}

const get_places = async () => {
  const response = await fetch("api/respawnplaces", {
    headers: {
      "Accept": "application/json"
    }
  });

  if (!response.ok) {
    throw new Error(
      `Unable to load respawn places: ${response.status} ${response.statusText}`
    );
  }
  const places = await response.json();
  return places || {};
};

export default L.LayerGroup.extend({
  initialize: function(cfg) {
    L.LayerGroup.prototype.initialize.call(this);
    this.placescolor = "green";
    if (cfg.respawn.placescolor)
      this.placescolor = cfg.respawn.placescolor;
    this.currentObjects = {};
    this.reDraw = this.reDraw.bind(this);
    this.onMapMove = debounce(this.onMapMove.bind(this), 50);
    // no onMinetestUpdate needed, like AirUtilsPlanesOverlay.js. Places do not change very often.
  },

  createPopup: function(place, name) {
    if (!name) name = "unnamed";
    let html = "<b>" + escapeHtml(name) + "</b>";
    if (place.full_name)
      html += "<br/>" + escapeHtml(place.full_name);
    return html;
  },

  getIcon: function(color){
    return L.AwesomeMarkers.icon({
      // person-arrow-up-from-line is my first choice, but not present
      // in the bundled fontawesome file.
      icon: "user-plus",
      prefix: "fa",
      markerColor: color || this.placescolor || "green"
    })
  },

  createMarker: function(place, name) {
    let marker = L.marker([place.pos.z + 0.5, place.pos.x + 0.5]);
    marker.bindPopup(this.createPopup(place, name));
    marker.setIcon(this.getIcon());
    return marker;
  },

  isInCurrentLayer: function(plane) {
    let mapLayer = layerMgr.getCurrentLayer();

    return (
      plane.pos.y >= (mapLayer.from*16) &&
      plane.pos.y <= ((mapLayer.to*16) + 15)
    );
  },

  getMaxDisplayedZoom: function(){
    return 5;
  },

  reDraw: async function() {
    if (this.map.getZoom() < this.getMaxDisplayedZoom()) {
      this.currentObjects = {};
      this.clearLayers();
      return;
    }

    const places = await get_places();
    this.currentObjects = {};
    this.clearLayers();

    Object.entries(places).forEach(([name,place]) => {
      if (!this.isInCurrentLayer(place)) {
        // not in current layer
        return;
      }
      let marker = this.createMarker(place, name);
      marker.addTo(this);
      this.currentObjects[name] = marker;
    });
  },

  onMapMove: function(){
    this.reDraw(false);
  },

  onLayerChange: function(/*layer*/){
    this.reDraw(true);
  },

  onAdd: function(map) {
    this.map = map;
    map.on("zoomend", this.onMapMove);
    this.reDraw();
  },

  onRemove: function(map) {
    map.off("zoomend", this.onMapMove);
    this.clearLayers();
  }
});
