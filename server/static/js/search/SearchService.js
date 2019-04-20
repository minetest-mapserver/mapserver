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

      function searchFor(q){
        q.pos1 = { x:-2048, y:-2048, z:-2048 };
        q.pos2 = { x:2048, y:2048, z:2048 };
        return api.getMapObjects(q);
      }

      var shop_prom = searchFor({
        type: "shop",
        attributelike: {
          key: "out_item",
          value: "%" + SearchStore.query + "%"
        }
      });

      var poi_prom = searchFor({
        type: "poi",
        attributelike: {
          key: "name",
          value: "%" + SearchStore.query + "%"
        }
      });

      Promise.all([shop_prom, poi_prom]).then(function(results){
        SearchStore.result = results[1].concat(results[0]);
      });

    }, 400),

    clear: function(){
      SearchStore.query = "";
      SearchStore.result = [];
    }
};
