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
  //overlays["Protector"] = new ProtectorOverlay(wsChannel, layerMgr);

  map.addLayer(overlays["Player"]);

  L.control.layers(layers, overlays).addTo(map);

  new CoordinatesDisplay({ position: 'bottomleft' }).addTo(map);
  new WorldInfoDisplay(wsChannel, { position: 'bottomright' }).addTo(map);

}).catch(function(e){
  console.error(e);
});
