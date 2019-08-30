
var customOverlays = {};

try {
  customOverlays = JSON.parse(localStorage["mapserver-customOverlays"])
} catch (e){}

function save(){
  localStorage["mapserver-customOverlays"] = JSON.stringify(customOverlays);
}

export default function(map, overlays){

  Object.keys(customOverlays)
  .filter(name => overlays[name])
  .forEach(name => {
    const layer = overlays[name];

    if (customOverlays[name] && !map.hasLayer(layer)){
      //Add
      map.addLayer(layer);
    }

    if (!customOverlays[name] && map.hasLayer(layer)){
      //Remove
      map.removeLayer(layer);
    }
  });

  map.on('overlayadd', e => {
    customOverlays[e.name] = true;
    save();
  });

  map.on('overlayremove', e => {
    customOverlays[e.name] = false;
    save();
  });

}
