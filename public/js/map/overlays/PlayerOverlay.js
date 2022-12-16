import wsChannel from '../../WebSocketChannel.js';
import layerMgr from '../../LayerManager.js';

const defaultSkin = "pics/sam.png";

let players = [];
let playerSkins = {};

//update players all the time
wsChannel.addListener("minetest-info", function(info){
  players = info.players || [];
});

export default L.LayerGroup.extend({
  initialize: function() {
    L.LayerGroup.prototype.initialize.call(this);

    this.currentObjects = {}; // name => marker

    this.reDraw = this.reDraw.bind(this);
    this.onMinetestUpdate = this.onMinetestUpdate.bind(this);
  },

  createPopup: function(player) {
    // moderators get a small crown icon
    let moderator = player.moderator ? `<img src="pics/crown.png">` : "";

    let info = `<b>${moderator} ${player.name}</b>`;
    info += "<hr>";

    for (let i = 0; i < Math.floor(player.hp / 2); i++)
      info += "<img src='pics/heart.png'>";

    if (player.hp % 2 === 1)
      info += "<img src='pics/heart_half.png'>";

    info += "<br>";

    for (let i = 0; i < Math.floor(player.breath / 2); i++)
      info += "<img src='pics/bubble.png'>";

    if (player.breath % 2 === 1)
      info += "<img src='pics/bubble_half.png'>";

    info += `
      <br>
      <b>RTT:</b> ${Math.floor(player.rtt*1000)} ms
      <br>
      <b>Protocol version:</b> ${player.protocol_version}
    `;

    info = `<div class="info">${info}</div>`;

    let portrait = `<img class="portrait" src="${this.getSkin(player)}" alt="${player.name}">`;

    return `<div class="player-popup">${portrait}${info}</div>`;
  },

  createMarker: function(player) {
    const marker = L.marker([player.pos.z, player.pos.x], {icon: this.getIcon(player)});

    marker.bindPopup(this.createPopup(player), {minWidth: 220});
    return marker;
  },

  getIcon: function(player) {
    const icon = this.getSkin(player);
    const indicator = player.yaw === 0 // compatibility with mapserver_mod without `yaw` attribute - value will be 0.
        ? false
        : player.velocity.x !== 0 || player.velocity.z !== 0 ? 'pics/sam_dir_move.png' : 'pics/sam_dir.png';
    return L.divIcon({
      html: `<div style="display:inline-block;width:48px;height:48px">
          <img src="${icon}" style="position:absolute;top:8px;left:16px;width:16px;height:32px;" alt="${player.name}">
          ${indicator 
          ? `<img src="${indicator}" style="position:absolute;top:0;left:0;width:48px;height:48px;transform:rotate(${player.yaw*-1}rad)" alt="${player.name}">`
          : ''}
        </div>`,
      className: '', // don't use leaflet default of a white block
      iconSize:     [48, 48],
      iconAnchor:   [24, 24],
      popupAnchor:  [0, -24]
    });
  },

  getSkin: function(player) {
    if (!player.skin || player.skin === "" || player.skin === "character.png") return defaultSkin;

    let skin = `api/skins/${player.skin}`;

    if (playerSkins[skin]) return playerSkins[skin];

    // no cached skin, we need to build the image
    let img = new Image();
    img.onload = function() {
      const canvas = document.createElement("canvas");
      const ctx = canvas.getContext("2d");
      canvas.width = 16;
      canvas.height = 32;

      // head
      ctx.drawImage(img, 8, 8, 8, 8, 4, 0, 8, 8);
      // chest
      ctx.drawImage(img, 20, 20, 8, 12, 4, 8, 8, 12);
      // leg left
      ctx.drawImage(img, 4, 20, 4, 12, 4, 20, 4, 12);
      // leg right
      if (img.height === 64) {
        ctx.drawImage(img, 20, 52, 4, 12, 8, 20, 4, 12);
      } else {
        ctx.drawImage(img, 4, 20, 4, 12, 8, 20, 4, 12);
      }
      // arm left
      ctx.drawImage(img, 44, 20, 4, 12, 0, 8, 4, 12);
      // arm right
      if (img.height === 64) {
        ctx.drawImage(img, 36, 52, 4, 12, 12, 8, 4, 12);
      } else {
        ctx.drawImage(img, 44, 20, 4, 12, 12, 8, 4, 12);
      }

      // store the skin, so it gets used on next update
      playerSkins[skin] = canvas.toDataURL("image/png");
    }

    // trigger source image load
    img.src = skin;

    // return the default skin while the replacement loads
    return defaultSkin;
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
