import wsChannel from '../../WebSocketChannel.js';
import layerMgr from '../../LayerManager.js';

let planes = [];

let icons = {
  "default": {
    url: "pics/airutils_planes/supercub.png",
    size: 48
  },
  "hidroplane:hidro": {
    url: "pics/airutils_planes/hidro.png",
    size: 48
  },
  "supercub:supercub": {
    url: "pics/airutils_planes/supercub.png",
    size: 48
  },
  "pa28:pa28": {
    url: "pics/airutils_planes/pa28.png",
    size: 64
  },
  "trike:trike": {
    url: "pics/airutils_planes/trike.png",
    size: 40
  },
  "ju52:ju52": {
    url: "pics/airutils_planes/ju52.png",
    size: 72
  },
  "steampunk_blimp:blimp": {
    url: "pics/airutils_planes/blimp.png",
    size: 96
  },
};

// listening for realtime updates
wsChannel.addListener("minetest-info", function(info) {
  planes = info.airutils_planes || [];
});

export default L.LayerGroup.extend({
  initialize: function() {
    L.LayerGroup.prototype.initialize.call(this);

    this.currentObjects = {}; // id => marker

    this.reDraw = this.reDraw.bind(this);
    this.onMinetestUpdate = this.onMinetestUpdate.bind(this);
  },

  createPopup: function(plane) {
    let name = plane.name;
    if (!name) name = plane.entity.substring(plane.entity.indexOf(":")+1);

    let html = "<b>" + name + "</b><br>";
    html += "<hr>";
    html += "<b>Owner:</b> " + plane.owner + "<br>";
    html += "<b>Pilot:</b> " + (plane.driver ? plane.driver : "-") + "<br>";
    html += "<b>Passengers:</b> " + (plane.passenger ? plane.passenger : "-") + "<br>";
    return html;
  },

  createMarker: function(plane) {
    let marker = L.marker([plane.pos.z, plane.pos.x], {icon: this.getIcon(plane)});

    marker.bindPopup(this.createPopup(plane));

    return marker;
  },

  getIcon: function(plane) {
    let icon = icons[plane.entity];
    if (!icon) icon = icons.default;
    return L.divIcon({
      html: `<div style="display:inline-block;width:${icon.size}px;height:${icon.size}px;transform:rotate(${plane.yaw*-1}rad);mask:url(${icon.url}) center/contain;-webkit-mask:url(${icon.url}) center/contain;background:${plane.color}">
          <img src="${icon.url}" style="width:${icon.size}px;height:${icon.size}px;filter:saturate(0%);mix-blend-mode:multiply;" alt="${plane.name}">
        </div>`,
      className: '', // don't use leaflet default of a white block
      iconSize:     [icon.size, icon.size],
      iconAnchor:   [icon.size/2, icon.size/2],
      popupAnchor:  [0, -(icon.size/2)]
    });
  },

  isInCurrentLayer: function(plane) {
    let mapLayer = layerMgr.getCurrentLayer();

    return (
        plane.pos.y >= (mapLayer.from*16) &&
        plane.pos.y <= ((mapLayer.to*16) + 15)
    );
  },

  onMinetestUpdate: function(/*info*/) {

    planes.forEach(plane => {
      let isInLayer = this.isInCurrentLayer(plane);

      if (!isInLayer) {
        if (this.currentObjects[plane.id]) {
          //player is displayed and not on the layer anymore
          //Remove the marker and reference
          this.currentObjects[plane.id].remove();
          delete this.currentObjects[plane.id];
        }

        return;
      }

      if (this.currentObjects[plane.id]) {
        //marker exists
        let marker = this.currentObjects[plane.id];
        marker.setLatLng([plane.pos.z, plane.pos.x]);
        marker.setIcon(this.getIcon(plane));
        marker.setPopupContent(this.createPopup(plane));
      } else {
        //marker does not exist
        let marker = this.createMarker(plane);
        marker.addTo(this);

        this.currentObjects[plane.id] = marker;
      }
    });

    Object.keys(this.currentObjects).forEach(existingId => {
      let planeIsActive = planes.find(function(p) {
        return p.id == existingId;
      });

      if (!planeIsActive) {
        //player
        this.currentObjects[existingId].remove();
        delete this.currentObjects[existingId];
      }
    });
  },

  reDraw: function() {
    this.currentObjects = {};
    this.clearLayers();

    planes.forEach(plane => {
      if (!this.isInCurrentLayer(plane)) {
        //not in current layer
        return;
      }

      let marker = this.createMarker(plane);
      marker.addTo(this);
      this.currentObjects[plane.id] = marker;
    });

  },

  onAdd: function(/*map*/) {
    wsChannel.addListener("minetest-info", this.onMinetestUpdate);
    this.reDraw();
  },

  onRemove: function(/*map*/) {
    this.clearLayers();
    wsChannel.removeListener("minetest-info", this.onMinetestUpdate);
  }
});
