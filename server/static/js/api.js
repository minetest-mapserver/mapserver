
export function getMapObjects(query){
  return m.request({
    method: "POST",
    url: "api/mapobjects/",
    data: query
  });
}

export function getConfig(){
  return m.request("api/config");
}
