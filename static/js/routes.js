
import Map from './components/Map.js';

var Home = {
    view: function() {
        return "Home"
    }
}

export default {
  "/": Home,
  "/map/:layerId/:zoom/:lon/:lat": Map
}
