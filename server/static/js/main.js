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

  overlays["Player"] = new PlayerOverlay(wsChannel, layerMgr);
  overlays["POI"] = new PoiOverlay(wsChannel, layerMgr);
  overlays["Travelnet"] = new TravelnetOverlay(wsChannel, layerMgr);
  overlays["Bones"] = new BonesOverlay(wsChannel, layerMgr);
  overlays["Digilines LCD"] = new LcdOverlay(wsChannel, layerMgr);
  overlays["Digiterms"] = new DigitermOverlay(wsChannel, layerMgr);
  overlays["Lua Controller"] = new LuacontrollerOverlay(wsChannel, layerMgr);
  overlays["Technic Anchor"] = new TechnicAnchorOverlay(wsChannel, layerMgr);
  overlays["Technic Quarry"] = new TechnicQuarryOverlay(wsChannel, layerMgr);
  //overlays["Protector"] = new ProtectorOverlay(wsChannel, layerMgr);

  //Default enabled overlays
  map.addLayer(overlays["Player"]);
  map.addLayer(overlays["POI"]);

  L.control.layers(layers, overlays).addTo(map);

  new CoordinatesDisplay({ position: 'bottomleft' }).addTo(map);
  new WorldInfoDisplay(wsChannel, { position: 'bottomright' }).addTo(map);

}).catch(function(e){
  console.error(e);
});
