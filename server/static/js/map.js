
export function setup(cfg){
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

  var overlays = {};

  window.layerMgr = new LayerManager(wsChannel, cfg.layers, map, Hashroute.getLayerId());

  //All overlays
  Overlaysetup(cfg, map, overlays, wsChannel, layerMgr);


  new CoordinatesDisplay({ position: 'bottomleft' }).addTo(map);
  new WorldInfoDisplay(wsChannel, { position: 'bottomright' }).addTo(map);

  if (cfg.enablesearch){
    new SearchControl(wsChannel, { position: 'topright' }).addTo(map);
  }

  //layer control
  L.control.layers(layerMgr.layerObjects, overlays, { position: "topright" }).addTo(map);

  Hashroute.setup(map, layerMgr);
}
