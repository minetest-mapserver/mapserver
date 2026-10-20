import SearchResult from './SearchResult.js';
import { getMapObjects, getRespawnPlaces } from '../api.js';

const state = {
  busy: false,
  result: []
};

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

function searchRespawnFor(type, key, valuelike) {
  const needle = valuelike.toLowerCase();
  return getRespawnPlaces().then(function(places) {
    return Object.entries(places)
      .filter(function([placeKey, place]) {
        return (
          placeKey.toLowerCase().includes(needle) ||
          (place.full_name || "").toLowerCase().includes(needle)
        );
      })
      .map(function([placeKey, place]) {
        const pos = place.pos;
        return {
          mapblock: {
            x: Math.floor(pos.x / 16),
            y: Math.floor(pos.y / 16),
            z: Math.floor(pos.z / 16)
          },
          x: pos.x,
          y: pos.y,
          z: pos.z,
          type: type,
          attributes: {
            // Use the requested key so the result has a name field.
            [key]: place.full_name || placeKey,
            full_name: place.full_name,
            key: placeKey,
            color: place.color,
            icon: place.icon
          }
        };
      });
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
    searchFor("digilinelcd", "text", query),
    searchFor("um_area_forsale", "description", query),
    searchRespawnFor("respawn_place", "name", query),
    searchFor("sign", "display_text", query)
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
    search(vnode.attrs.query);
  },

  view(vnode){
    if (state.result.length == 0) {
      return m("div", vnode.attrs.query);
    } else {
      return m(SearchResult, { result: state.result });
    }
  }
};
