/* exported SearchService */
/* globals SearchStore: true */

var SearchService = {

    search: function(q){
      SearchStore.query = q;

      this.fetchData();
    },

    fetchData: debounce(function(){
      SearchStore.result = [];

      if (!SearchStore.query){
        return;
      }

      api.getMapObjects({
        pos1: { x:-2048, y:-2048, z:-2048 },
        pos2: { x:2048, y:2048, z:2048 },
        type: "shop",
        attributelike: {
          key: "out_item",
          value: "%" + SearchStore.query + "%"
        }
      })
      .then(function(result){
        SearchStore.result = result;
        //console.log(result); //XXX
      });

    }, 400),

    clear: function(){
      SearchStore.query = "";
      SearchStore.result = [];
    }
};
