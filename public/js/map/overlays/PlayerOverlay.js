import wsChannel from '../../WebSocketChannel.js';
import layerMgr from '../../LayerManager.js';

let players = [];

//update players all the time
wsChannel.addListener("minetest-info", function(info){
  players = info.players || [];
});

var PlayerIcon = L.icon({
  iconUrl: 'pics/sam.png',

  iconSize:     [16, 32],
  iconAnchor:   [8, 16],
  popupAnchor:  [0, -16]
});

export default L.LayerGroup.extend({
  initialize: function() {
    L.LayerGroup.prototype.initialize.call(this);

    this.currentObjects = {}; // name => marker

    this.reDraw = this.reDraw.bind(this);
    this.onMinetestUpdate = this.onMinetestUpdate.bind(this);
  },

  createPopup: function(player){
    let html = "<b>" + player.name + "</b>";
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
      <b>RTT:</b> ${Math.floor(player.rtt*1000)} ms
      <br>
      <b>Protocol version:</b> ${player.protocol_version}
    `;

    return html;
  },

  createMarker: function(player){
    const marker = L.marker([player.pos.z, player.pos.x], {icon: this.getIcon(player)});

    marker.bindPopup(this.createPopup(player));
    return marker;
  },

  getIcon: function(player) {
    /*
      compatibility with mapserver_mod without `yaw` attribute - value will be 0.
      if a player manages to look exactly north, the indicator will also disappear
      but aligning view at 0.0 is difficult/unlikely during normal gameplay.
    */
    if (player.yaw === 0) return PlayerIcon;

    const icon = 'pics/sam.png';
    const indicator = player.velocity.x !== 0 || player.velocity.z !== 0 ? 'pics/sam_dir_move.png' : 'pics/sam_dir.png';
    return L.divIcon({
      html: `<div style="display:inline-block;width:48px;height:48px">
          <img src="${icon}" style="position:absolute;top:8px;left:16px;width:16px;height:32px;" alt="${player.name}">
          <img src="${indicator}" style="position:absolute;top:0;left:0;width:48px;height:48px;transform:rotate(${player.yaw*-1}rad)" alt="${player.name}">
        </div>`,
      className: '', // don't use leaflet default of a white block
      iconSize:     [48, 48],
      iconAnchor:   [24, 24],
      popupAnchor:  [0, -24]
    });
  },

  isPlayerInCurrentLayer: function(player){
    const mapLayer = layerMgr.getCurrentLayer();

    return (
      player.pos.y >= (mapLayer.from*16) &&
      player.pos.y <= ((mapLayer.to*16) + 15)
    );
  },

  onMinetestUpdate: function(/*info*/){

    players.forEach(player => {
      const isInLayer = this.isPlayerInCurrentLayer(player);

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
        const marker = this.currentObjects[player.name];
        marker.setLatLng([player.pos.z, player.pos.x]);
        marker.setIcon(this.getIcon(player));
        marker.setPopupContent(this.createPopup(player));

      } else {
        //marker does not exist
        const marker = this.createMarker(player);
        marker.addTo(this);

        this.currentObjects[player.name] = marker;
      }
    });

    Object.keys(this.currentObjects).forEach(existingName => {
      const playerIsActive = players.find(function(p){
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

    players.forEach(player => {
      if (!this.isPlayerInCurrentLayer(player)){
        //not in current layer
        return;
      }

      const marker = this.createMarker(player);
      marker.addTo(this);
      this.currentObjects[player.name] = marker;
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
