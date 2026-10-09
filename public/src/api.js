
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
