import SearchResult from './SearchResult.js';
import { getMapObjects } from '../api.js';

const state = {
  busy: false,
  result: []
}

function searchFor(type, key, valuelike){
  return getMapObjects({
    pos1: { x:-2048, y:-2048, z:-2048 },
    pos2: { x:2048, y:2048, z:2048 },
    type: type,
    attributelike: {
      key: key,
      value: "%" + valuelike +"%"
    }
  });
}

function search(query){
  state.result = [];

  var prom_list = [
    searchFor("shop", "out_item", query),
    searchFor("poi", "name", query),
    searchFor("train", "station", query),
    searchFor("travelnet", "station_name", query),
    searchFor("bones", "owner", query),
    searchFor("locator", "name", query),
    searchFor("label", "text", query),
    searchFor("digiterm", "display_text", query),
    searchFor("digilinelcd", "text", query)
  ];

  Promise.all(prom_list)
  .then(function(results){

    var arr = [];
    results.forEach(function(r) {
      arr = arr.concat(r);
    });

    state.result = arr;
    state.busy = false;
  });

}

export default {
  oncreate(vnode){
    console.log("oncreate", vnode);
    search(vnode.attrs.query);
  },

  view(vnode){
    if (state.result.length == 0) {
      return m("div", vnode.attrs.query);
    } else {
      return m(SearchResult, { result: state.result })
    }
  }
}
