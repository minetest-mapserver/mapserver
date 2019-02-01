'use strict';

api.getConfig().then(function(cfg){
  var layerMgr = new LayerManager(cfg.layers);

  var wsChannel = new WebSocketChannel();
  wsChannel.connect();

  var rtTiles = new RealtimeTileLayer(wsChannel);


  var crs = L.Util.extend({}, L.CRS.Simple, {
      //transformation: L.transformation(0.001, 0, -0.001, 0),
      scale: function (zoom) {
          return Math.pow(2, zoom-9);
      }
  });

  var initialZoom = 11;
  var initialCenter = [0, 0];

  var map = L.map('image-map', {
    minZoom: 2,
    maxZoom: 12,
    center: initialCenter,
    zoom: initialZoom,
    crs: crs
  });

  map.attributionControl.addAttribution('<a href="https://github.com/thomasrudin-mt/mapserver">Mapserver</a>');

  var layers = {};
  var overlays = {}

  var Layer = rtTiles.createLayer(0);
  var tileLayer = new Layer();
  tileLayer.addTo(map);

  layers["Base"] = tileLayer;
  overlays["Travelnet"] = new TravelnetOverlay();

  L.control.layers(layers, overlays).addTo(map);

  var el = CoordinatesDisplay.create({ position: 'bottomleft' });
  el.addTo(map);

});
