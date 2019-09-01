
import Map from './components/Map.js';
import Search from './components/Search.js';

var Home = {
    view: function() {
        return "Home";
    }
};

export default {
  "/": Home,
  "/map/:layerId/:zoom/:lon/:lat": Map,
  "/search/:query": Search
};
