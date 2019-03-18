var api = {

  getMapObjects: function(query){
    return m.request({
      method: "POST",
      url: "api/mapobjects/",
      data: query
    });
  },

  getConfig: function(){
    return m.request("api/config");
  }


};
