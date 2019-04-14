
var SearchStore = {
  query: "",
  result: [],

  search: function(q){
    this.query = q;

    this.fetchData();
  },

  fetchData: debounce(function(){
    var self = this;
    this.result = [];

    if (!this.query){
      return;
    }

    api.getMapObjects({
      pos1: { x:-2048, y:-2048, z:-2048 },
      pos2: { x:2048, y:2048, z:2048 },
      type: "shop",
      attributelike: {
        key: "out_item",
        value: "%" + this.query + "%"
      }
    })
    .then(function(result){
      self.result = result;
      console.log(result); //XXX
    });

  }, 400),

  clear: function(){
    this.query = "";
    this.result = [];
  }
};
