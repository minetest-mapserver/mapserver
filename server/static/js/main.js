
(function(){
    'use strict';

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

    WebSocketChannel.connect();

    map.attributionControl.addAttribution('<a href="https://github.com/thomasrudin-mt/mapserver">Mapserver</a>');

    var layers = {};
    var overlays = {}

    RealtimeTileLayer.init();

    var Layer = RealtimeTileLayer.create(0);
    var tileLayer = new Layer();
    tileLayer.addTo(map);

    layers["Base"] = tileLayer;
    overlays["Travelnet"] = new TravelnetOverlay();

    L.control.layers(layers, overlays).addTo(map);

    var el = CoordinatesDisplay.create({ position: 'bottomleft' });
    el.addTo(map);



})()
