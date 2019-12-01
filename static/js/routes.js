
import Map from './components/Map.js';
import Search from './components/Search.js';

export default {
  "/map/:layerId/:zoom/:lon/:lat": Map,
  "/search/:query": Search
};
