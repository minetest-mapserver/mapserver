
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

// texture or model as blob, null if it does not exist
export async function getViewMedia(filename){
  const res = await fetch(`api/view/media/${encodeURIComponent(filename)}`);
  return res.ok ? res.blob() : null;
}
