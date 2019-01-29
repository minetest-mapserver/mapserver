
(function(){

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


    var wsUrl = location.protocol.replace("http", "ws") + "//" + location.host + location.pathname.substring(0, location.pathname.lastIndexOf("/")) + "/api/ws";
    var ws = new WebSocket(wsUrl);

    ws.onmessage = function(e){
      var event = JSON.parse(e.data);

      if (event.type == "rendered-tile"){
        realtimelayer.update(event.data)
      }
    }

    var layers = {};

    var Layer = realtimelayer.create(0);
    var tileLayer = new Layer();
    tileLayer.addTo(map);

    L.control.layers(layers, {}).addTo(map);

    var el = L.control.coordinatesDisplay({ position: 'bottomleft' });
    el.addTo(map);



})()
