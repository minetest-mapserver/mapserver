'use strict';

api.getConfig().then(function(cfg){

  var wsChannel = new WebSocketChannel();
  wsChannel.connect();

  var map = L.map('image-map', {
    minZoom: 2,
    maxZoom: 12,
    center: Hashroute.getCenter(),
    zoom: Hashroute.getZoom(),
    crs: SimpleCRS
  });

  map.attributionControl.addAttribution('<a href="https://github.com/thomasrudin-mt/mapserver">Minetest Mapserver</a>');

  var layers = {};
  var overlays = {}

  var layerMgr = new LayerManager(cfg.layers, map);
  layerMgr.setLayerId( Hashroute.getLayerId() );

  //All layers
  cfg.layers.forEach(function(layer){
    var tileLayer = new RealtimeTileLayer(wsChannel, layer.id, map);
    layers[layer.name] = tileLayer;
  });

  //current layer
  var currentLayer = layerMgr.getCurrentLayer();
  layers[currentLayer.name].addTo(map);

  //All overlays
  Overlaysetup(cfg, map, overlays, wsChannel, layerMgr);

  L.control.layers(layers, overlays).addTo(map);

  new CoordinatesDisplay({ position: 'bottomleft' }).addTo(map);
  new WorldInfoDisplay(wsChannel, { position: 'bottomright' }).addTo(map);

  Hashroute.setup(map, layerMgr);

}).catch(function(e){
  console.error(e);
});
