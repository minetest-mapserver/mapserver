
import MapComponent from './map/MapComponent.js';

var Home = {
    view: function() {
        return "Home"
    }
}

export default {
  "/": Home,
  "/map/:layerId/:zoom/:lon/:lat": MapComponent
}
