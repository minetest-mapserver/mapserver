
export default L.LayerGroup.extend({
  initialize: function(wsChannel, layerMgr) {
    L.LayerGroup.prototype.initialize.call(this);

    this.layerMgr = layerMgr;
    this.wsChannel = wsChannel;

    this.currentObjects = {}; // name => marker
    this.signals = [];

    this.reDraw = this.reDraw.bind(this);
    this.onMinetestUpdate = this.onMinetestUpdate.bind(this);

    //update players all the time
    this.wsChannel.addListener("minetest-info", function(info){
      this.signals = info.signals || [];
    }.bind(this));
  },

  createPopup: function(signal){
    var html = "<b>Signal</b><hr>";
    html += "<b>State:</b> " + signal.green + "<br>";

    return html;
  },

  hashPos: function(x,y,z){
    return x + "/" + y + "/" + z;
  },

  getMaxDisplayedZoom: function(){
    return 10;
  },

  createMarker: function(signal){

    var Icon = L.icon({
      iconUrl: "TODO",

      iconSize:     [16, 16],
      iconAnchor:   [8, 8],
      popupAnchor:  [0, -16]
    });

    var marker = L.marker([signal.pos.z, signal.pos.x], {icon: Icon});
    marker.bindPopup(this.createPopup(signal));

    return marker;
  },

  isSignalInCurrentLayer: function(signal){
    var mapLayer = this.layerMgr.getCurrentLayer();

    return (signal.pos.y >= (mapLayer.from*16) && signal.pos.y <= (mapLayer.to*16));
  },


  onMinetestUpdate: function(/*info*/){

    if (this.map.getZoom() < this.getMaxDisplayedZoom()) {
      this.clearLayers();
      this.currentObjects = {};
      return;
    }

    this.signals.forEach(signal => {
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

      } else {
        //marker does not exist
        let marker = this.createMarker(signal);
        marker.addTo(this);

        this.currentObjects[signalId] = marker;
      }
    });

    Object.keys(this.currentObjects).forEach(existingId => {
      var signalIsActive = this.signals.find((t) => {
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

    var mapLayer = this.layerMgr.getCurrentLayer();

    this.signals.forEach(signal => {
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
    this.layerMgr.addListener(() => this.reDraw());
    this.wsChannel.addListener("minetest-info", () => this.onMinetestUpdate());
    this.reDraw();
  },

  onRemove: function(/*map*/) {
    this.clearLayers();
    this.layerMgr.removeListener(() => this.reDraw());
    this.wsChannel.removeListener("minetest-info", () => this.onMinetestUpdate());
  }
});
