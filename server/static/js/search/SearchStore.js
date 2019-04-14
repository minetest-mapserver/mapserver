
var SearchStore = {
  query: "",

  search: function(q){
    this.query = q;

    this.fetchData();
  },

  fetchData: debounce(function(){
    console.log(this.query);
    if (!this.query){
      return;
    }

    api.getMapObjects({
      pos1: { x:-2048, y:-2048, z:-2048 },
      pos2: { x:2048, y:2048, z:2048 },
      type: "shop"
    })
    .then(function(result){
      console.log(result);
    });

  }, 400),

  clear: function(){
    this.query = "";
  }
};
