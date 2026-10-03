
function json(res){
  if (!res.ok){
    throw new Error(`${res.status} ${res.statusText}: ${res.url}`);
  }
  return res.json();
}

export function getMapObjects(query){
  return fetch("api/mapobjects/", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(query)
  }).then(json);
}

export function getConfig(){
  return fetch("api/config").then(json);
}

export function getStats(){
  return fetch("api/stats").then(json);
}
