
var PlayerIcon = L.icon({
  iconUrl: 'pics/sam.png',

  iconSize:     [16, 32],
  iconAnchor:   [8, 16],
  popupAnchor:  [0, -16]
});

export default L.LayerGroup.extend({
  initialize: function(wsChannel, layerMgr) {
    L.LayerGroup.prototype.initialize.call(this);

    this.layerMgr = layerMgr;
    this.wsChannel = wsChannel;

    this.currentObjects = {}; // name => marker
    this.players = [];

    this.reDraw = this.reDraw.bind(this);
    this.onMinetestUpdate = this.onMinetestUpdate.bind(this);

    //update players all the time
    this.wsChannel.addListener("minetest-info", function(info){
      this.players = info.players || [];
    }.bind(this));
  },

  createPopup: function(player){
    var html = "<b>" + player.name + "</b>";
    html += "<hr>";

    for (let i=0; i<Math.floor(player.hp / 2); i++)
      html += "<img src='pics/heart.png'>";

    if (player.hp % 2 == 1)
      html += "<img src='pics/heart_half.png'>";

    html += "<br>";

    for (let i=0; i<Math.floor(player.breath / 2); i++)
      html += "<img src='pics/bubble.png'>";

    if (player.breath % 2 == 1)
      html += "<img src='pics/bubble_half.png'>";

    html += `
      <br>
      <b>RTT:</b> ${player.rtt} ms
      <br>
      <b>Protocol version:</b> ${player.protocol_version}
    `;

    return html;
  },

  createMarker: function(player){
    var marker = L.marker([player.pos.z, player.pos.x], {icon: PlayerIcon});

    marker.bindPopup(this.createPopup(player));
    return marker;
  },

  isPlayerInCurrentLayer: function(player){
    var mapLayer = this.layerMgr.getCurrentLayer();

    return (player.pos.y >= (mapLayer.from*16) && player.pos.y <= (mapLayer.to*16));
  },

  onMinetestUpdate: function(/*info*/){

    this.players.forEach(player => {
      var isInLayer = this.isPlayerInCurrentLayer(player);

      if (!isInLayer){
        if (this.currentObjects[player.name]){
          //player is displayed and not on the layer anymore
          //Remove the marker and reference
          this.currentObjects[player.name].remove();
          delete this.currentObjects[player.name];
        }

        return;
      }

      if (this.currentObjects[player.name]){
        //marker exists
        let marker = this.currentObjects[player.name];
        marker.setLatLng([player.pos.z, player.pos.x]);
        marker.setPopupContent(this.createPopup(player));

      } else {
        //marker does not exist
        let marker = this.createMarker(player);
        marker.addTo(this);

        this.currentObjects[player.name] = marker;
      }
    });

    Object.keys(this.currentObjects).forEach(existingName => {
      var playerIsActive = this.players.find(function(p){
        return p.name == existingName;
      });

      if (!playerIsActive){
        //player
        this.currentObjects[existingName].remove();
        delete this.currentObjects[existingName];
      }
    });
  },

  reDraw: function(){
    this.currentObjects = {};
    this.clearLayers();

    var mapLayer = this.layerMgr.getCurrentLayer();

    this.players.forEach(player => {
      if (!this.isPlayerInCurrentLayer(player)){
        //not in current layer
        return;
      }

      var marker = this.createMarker(player);
      marker.addTo(this);
      this.currentObjects[player.name] = marker;
    });

  },

  onAdd: function(/*map*/) {
    this.layerMgr.addListener(this.reDraw);
    this.wsChannel.addListener("minetest-info", this.onMinetestUpdate);
    this.reDraw();
  },

  onRemove: function(/*map*/) {
    this.clearLayers();
    this.layerMgr.removeListener(this.reDraw);
    this.wsChannel.removeListener("minetest-info", this.onMinetestUpdate);
  }
});
