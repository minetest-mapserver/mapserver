var api = {

  getMapObjects: function(x1,y1,z1,x2,y2,z2,type){
    return m.request("/api/mapobjects/" +
      x1 + "/" + y1 + "/" + z1 + "/" +
      x2 + "/" + y2 + "/" + z2 + "/" +
      type
    );
  },

  getConfig: function(){
    return m.request("/api/config");
  }


};
