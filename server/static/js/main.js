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

  //All layers
  cfg.layers.forEach(function(layer){
    var tileLayer = new RealtimeTileLayer(wsChannel, layer.id);
    tileLayer.addTo(map);

    layers[layer.name] = tileLayer;
  });

  //All overlays
  Overlaysetup(cfg, map, overlays, wsChannel, layerMgr);

  L.control.layers(layers, overlays).addTo(map);

  new CoordinatesDisplay({ position: 'bottomleft' }).addTo(map);
  new WorldInfoDisplay(wsChannel, { position: 'bottomright' }).addTo(map);

}).catch(function(e){
  console.error(e);
});
