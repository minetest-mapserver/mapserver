
export getMapObjects(query){
  return m.request({
    method: "POST",
    url: "api/mapobjects/",
    data: query
  });
}

export getConfig(){
  return m.request("api/config");
}
