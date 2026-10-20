
export function getMapObjects(query){
  return m.request({
    method: "POST",
    url: "api/mapobjects/",
    data: query
  });
}

export function getRespawnPlaces() {
  return m.request({
    method: "GET",
    url: "api/respawnplaces"
  });
}

export function getConfig(){
  return m.request("api/config");
}

export function getStats(){
	return m.request("api/stats");
}
