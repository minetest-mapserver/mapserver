
function json(res){
  if (!res.ok){
    throw new Error(`${res.status} ${res.statusText}: ${res.url}`);
  }
  return res.json();
}

export async function getMapObjects(query){
  const res = await fetch("api/mapobjects/", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(query)
  });
  return json(res);
}

// place key -> { pos, full_name, color, icon }
export async function getRespawnPlaces(){
  return json(await fetch("api/respawnplaces"));
}

// topmost non-air position of a layer at x/z, null if there is none
export async function getStartHeight(layerId, x, z){
  const res = await fetch(`api/startheight/${layerId}/${x}/${z}`);
  if (res.status == 404){
    return null;
  }
  return json(res);
}

export async function getConfig(){
  return json(await fetch("api/config"));
}

export async function getStats(){
  return json(await fetch("api/stats"));
}

// batch of mapblocks: chunk position = floor(mapblock position / 5)
export async function getViewChunk(x, y, z){
  return json(await fetch(`api/viewchunk/${x}/${y}/${z}`));
}

// node name -> node definition
export async function getViewNodedefs(){
  return json(await fetch("api/view/nodedefs"));
}

// batch of textures and models: list of { filename, data (base64 without padding), type }
// files that do not exist are omitted
export async function getViewMediaBatch(filenames){
  const res = await fetch("api/view/media", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(filenames)
  });
  return json(res);
}
