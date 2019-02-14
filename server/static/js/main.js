'use strict';

api.getConfig().then(function(cfg){

  var wsChannel = new WebSocketChannel();
  wsChannel.connect();

  var initialZoom = 11;
  var initialCenter = [0, 0];

  var map = L.map('image-map', {
    minZoom: 2,
    maxZoom: 12,
    center: initialCenter,
    zoom: initialZoom,
    crs: SimpleCRS
  });

  map.attributionControl.addAttribution('<a href="https://github.com/thomasrudin-mt/mapserver">Minetest Mapserver</a>');

  var layers = {};
  var overlays = {}

  var layerMgr = new LayerManager(cfg.layers, map);

  var tileLayer = new RealtimeTileLayer(wsChannel, 0);
  tileLayer.addTo(map);

  //TODO: all layers
  layers["Base"] = tileLayer;

  if (cfg.mapobjects.mapserver) {
    overlays["Player"] = new PlayerOverlay(wsChannel, layerMgr);
    overlays["POI"] = new PoiOverlay(wsChannel, layerMgr);
    overlays["Labels"] = new LabelOverlay(wsChannel, layerMgr);

    map.addLayer(overlays["Player"]);
    map.addLayer(overlays["POI"]);
    map.addLayer(overlays["Labels"]);
  }

  if (cfg.mapobjects.travelnet) {
    overlays["Travelnet"] = new TravelnetOverlay(wsChannel, layerMgr);
  }

  if (cfg.mapobjects.bones) {
    overlays["Bones"] = new BonesOverlay(wsChannel, layerMgr);
  }

  if (cfg.mapobjects.digilines) {
    overlays["Digilines LCD"] = new LcdOverlay(wsChannel, layerMgr);
  }

  if (cfg.mapobjects.digiterms) {
    overlays["Digiterms"] = new DigitermOverlay(wsChannel, layerMgr);
  }

  if (cfg.mapobjects.luacontroller) {
    overlays["Lua Controller"] = new LuacontrollerOverlay(wsChannel, layerMgr);
  }

  if (cfg.mapobjects.technic) {
    overlays["Technic Anchor"] = new TechnicAnchorOverlay(wsChannel, layerMgr);
    overlays["Technic Quarry"] = new TechnicQuarryOverlay(wsChannel, layerMgr);
    overlays["Technic Switching station"] = new TechnicSwitchOverlay(wsChannel, layerMgr);
  }

  if (cfg.mapobjects.protector) {
    overlays["Protector"] = new ProtectorOverlay(wsChannel, layerMgr);
  }

  if (cfg.mapobjects.mission) {
    overlays["Missions"] = new MissionOverlay(wsChannel, layerMgr);
  }

  L.control.layers(layers, overlays).addTo(map);

  new CoordinatesDisplay({ position: 'bottomleft' }).addTo(map);
  new WorldInfoDisplay(wsChannel, { position: 'bottomright' }).addTo(map);

}).catch(function(e){
  console.error(e);
});
