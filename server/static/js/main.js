'use strict';

api.getConfig().then(function(cfg){

  var wsChannel = new WebSocketChannel();
  wsChannel.connect();

  wsChannel.addListener("minetest-info", function(e){
    console.log(e); //XXX
  });

  var rtTiles = new RealtimeTileLayer(wsChannel);

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

  layers["Base"] = tileLayer;
  overlays["Players"] = new PlayerOverlay(wsChannel, layerMgr);
  overlays["Travelnet"] = new TravelnetOverlay(wsChannel, layerMgr);

  L.control.layers(layers, overlays).addTo(map);

  var el = new CoordinatesDisplay({ position: 'bottomleft' });
  el.addTo(map);

});
