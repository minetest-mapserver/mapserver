
export function getMapObjects(query){
  return fetch("api/mapobjects/", {
    method: "POST",
    body: JSON.stringify(query)
  })
  .then(r => this.json());
}

export function getConfig(){
  return fetch("api/config")
    .then(r => r.json());
}

export function getStats(){
	return fetch("api/stats")
    .then(r => r.json());
}
