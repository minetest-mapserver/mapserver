import wsChannel from '../../WebSocketChannel.js';
import layerMgr from '../../LayerManager.js';

var IconOn = L.icon({
  iconUrl: "pics/advtrains/advtrains_signal_on.png",
  iconSize:     [16, 16],
  iconAnchor:   [8, 8],
  popupAnchor:  [0, -16]
});

var IconOff = L.icon({
  iconUrl: "pics/advtrains/advtrains_signal_off.png",
  iconSize:     [16, 16],
  iconAnchor:   [8, 8],
  popupAnchor:  [0, -16]
});

let signals = [];

//update signals all the time
wsChannel.addListener("minetest-info", function(info){
  signals = info.signals || [];
});

export default L.LayerGroup.extend({
  initialize: function() {
    L.LayerGroup.prototype.initialize.call(this);

    this.currentObjects = {}; // name => marker
    this.onMinetestUpdate = this.onMinetestUpdate.bind(this);
  },

  createPopup: function(signal){
    var html = "<b>Signal</b><hr>";
    html += "<b>State:</b> " +
      (signal.green ? "Green" : "Red") +
      "<br>";

    return html;
  },

  hashPos: function(x,y,z){
    return x + "/" + y + "/" + z;
  },

  getMaxDisplayedZoom: function(){
    return 10;
  },

  createMarker: function(signal){

    var Icon = signal.green ? IconOn : IconOff;
    var marker = L.marker([signal.pos.z, signal.pos.x], {icon: Icon});
    marker.bindPopup(this.createPopup(signal));

    return marker;
  },

  isSignalInCurrentLayer: function(signal){
    var mapLayer = layerMgr.getCurrentLayer();

    return (signal.pos.y >= (mapLayer.from*16) && signal.pos.y <= (mapLayer.to*16));
  },


  onMinetestUpdate: function(/*info*/){

    if (this.map.getZoom() < this.getMaxDisplayedZoom()) {
      this.clearLayers();
      this.currentObjects = {};
      return;
    }

    signals.forEach(signal => {
      var isInLayer = this.isSignalInCurrentLayer(signal);
      var signalId = this.hashPos(signal.pos.x, signal.pos.y, signal.pos.z);

      if (!isInLayer){
        if (this.currentObjects[signalId]){
          //signal is displayed and not on the layer anymore
          //Remove the marker and reference
          this.currentObjects[signalId].remove();
          delete this.currentObjects[signalId];
        }

        return;
      }

      if (this.currentObjects[signalId]){
        //marker exists
        let marker = this.currentObjects[signalId];
        marker.setLatLng([signal.pos.z, signal.pos.x]);
        marker.setPopupContent(this.createPopup(signal));
        marker.setIcon(signal.green ? IconOn : IconOff);

      } else {
        //marker does not exist
        let marker = this.createMarker(signal);
        marker.addTo(this);

        this.currentObjects[signalId] = marker;
      }
    });

    Object.keys(this.currentObjects).forEach(existingId => {
      var signalIsActive = signals.find((t) => {
	var hash = this.hashPos(t.pos.x, t.pos.y, t.pos.z);
        return hash == existingId;
      });

      if (!signalIsActive){
        this.currentObjects[existingId].remove();
        delete this.currentObjects[existingId];
      }
    });
  },

  reDraw: function(){
    this.currentObjects = {};
    this.clearLayers();

    if (this.map.getZoom() < this.getMaxDisplayedZoom()) {
      return;
    }

    var mapLayer = layerMgr.getCurrentLayer();

    signals.forEach(signal => {
      if (!this.isSignalInCurrentLayer(signal)){
        //not in current layer
        return;
      }

      var marker = this.createMarker(signal);
      marker.addTo(this);
      var hash = this.hashPos(signal.pos.x, signal.pos.y, signal.pos.z);
      this.currentObjects[hash] = marker;
    });

  },

  onAdd: function(map) {
    this.map = map;
    wsChannel.addListener("minetest-info", this.onMinetestUpdate);
    this.reDraw();
  },

  onRemove: function(/*map*/) {
    this.clearLayers();
    wsChannel.removeListener("minetest-info", this.onMinetestUpdate);
  }
});
