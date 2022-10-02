import wsChannel from '../../WebSocketChannel.js';
import layerMgr from '../../LayerManager.js';

function getTrainImageUrlForType(type){
  switch(type){
    case "advtrains:subway_wagon":
      return "pics/advtrains/advtrains_subway_wagon_inv.png";
    case "advtrains:engine_japan":
      return  "pics/advtrains/advtrains_engine_japan_inv.png";
    case "advtrains:wagon_japan":
      return  "pics/advtrains/advtrains_wagon_japan_inv.png";
    case "advtrains:engine_steam":
      return  "pics/advtrains/advtrains_engine_steam_inv.png";
    case "advtrains:detailed_steam_engine":
      return  "pics/advtrains/advtrains_detailed_engine_steam_inv.png";
    case "advtrains:engine_industrial":
      return  "pics/advtrains/advtrains_engine_industrial_inv.png";
    case "advtrains:wagon_wood":
      return  "pics/advtrains/advtrains_wagon_wood_inv.png";
    case "advtrains:wagon_box":
      return  "pics/advtrains/advtrains_wagon_box_inv.png";
    case "advtrains:wagon_default":
      return  "pics/advtrains/advtrains_wagon_inv.png";

    case "advtrains:subway_wagon_blue":
      return "pics/advtrains/advtrains_subway_wagon_inv_blue.png";
    case "advtrains:subway_wagon_red":
      return "pics/advtrains/advtrains_subway_wagon_inv_red.png";
    case "advtrains:subway_wagon_green":
      return "pics/advtrains/advtrains_subway_wagon_inv_green.png";

    default:
      //TODO: fallback image
      return "pics/advtrains/advtrains_subway_wagon_inv.png";
  }
}

let trains = [];

//update trains all the time
wsChannel.addListener("minetest-info", function(info){
  trains = info.trains || [];
});

export default L.LayerGroup.extend({
  initialize: function() {
    L.LayerGroup.prototype.initialize.call(this);

    this.currentObjects = {}; // name => marker

    this.onMinetestUpdate = this.onMinetestUpdate.bind(this);
  },

  createPopup: function(train){
    var html = "<b>Train</b><hr>";

    html += "<b>Name:</b> " + train.text_outside + "<br>";
    html += "<b>Line:</b> " + train.line + "<br>";
    html += "<b>Velocity:</b> "+ Math.floor(train.velocity*10)/10 + "<br>";

    if (train.wagons){
	    html += "<b>Composition: </b>";
	    train.wagons.forEach(function(w){
	      var iconUrl =  getTrainImageUrlForType(w.type);
	      html += "<img src='"+iconUrl+"'>";
	    });
    }

    return html;
  },


  getMaxDisplayedZoom: function(){
    return 10;
  },

  createMarker: function(train){

    //search for wagon in front (whatever "front" is...)
    var type;
    var lowest_pos = 100;
    if (train.wagons){
    	train.wagons.forEach(function(w){
      		if (w.pos_in_train < lowest_pos){
       			lowest_pos = w.pos_in_train;
       			type = w.type;
      		}
    	});
    }

    var Icon = L.icon({
      iconUrl: getTrainImageUrlForType(type),

      iconSize:     [16, 16],
      iconAnchor:   [8, 8],
      popupAnchor:  [0, -16]
    });

    var marker = L.marker([train.pos.z, train.pos.x], {icon: Icon});
    marker.bindPopup(this.createPopup(train));

    return marker;
  },

  isTrainInCurrentLayer: function(train){
    var mapLayer = layerMgr.getCurrentLayer();

    return (train.pos.y >= (mapLayer.from*16) && train.pos.y <= (mapLayer.to*16));
  },


  onMinetestUpdate: function(/*info*/){

    if (this.map.getZoom() < this.getMaxDisplayedZoom()) {
      this.clearLayers();
      this.currentObjects = {};
      return;
    }

    trains.forEach(train => {
      var isInLayer = this.isTrainInCurrentLayer(train);

      if (!isInLayer){
        if (this.currentObjects[train.id]){
          //train is displayed and not on the layer anymore
          //Remove the marker and reference
          this.currentObjects[train.id].remove();
          delete this.currentObjects[train.id];
        }

        return;
      }

      if (this.currentObjects[train.id]){
        //marker exists
        let marker = this.currentObjects[train.id];
        marker.setLatLng([train.pos.z, train.pos.x]);
        marker.setPopupContent(this.createPopup(train));

      } else {
        //marker does not exist
        let marker = this.createMarker(train);
        marker.addTo(this);

        this.currentObjects[train.id] = marker;
      }
    });

    Object.keys(this.currentObjects).forEach(existingId => {
      var trainIsActive = trains.find(function(t){
        return t.id == existingId;
      });

      if (!trainIsActive){
        //train
        this.currentObjects[existingId].remove();
        delete this.currentObjects[existingId];
      }
    });
  },

  reDraw: function(){
    this.currentObjects = {};
    this.clearLayers();

    if (this.map.getZoom() < this.getMaxDisplayedZoom()) {
      return;
    }

    trains.forEach(train => {
      if (!this.isTrainInCurrentLayer(train)){
        //not in current layer
        return;
      }

      var marker = this.createMarker(train);
      marker.addTo(this);
      this.currentObjects[train.id] = marker;
    });

  },

  onAdd: function(map) {
    this.map = map;
    wsChannel.addListener("minetest-info", this.onMinetestUpdate);
    this.reDraw();
  },

  onRemove: function(/*map*/) {
    this.clearLayers();
    wsChannel.removeListener("minetest-info", this.onMinetestUpdate);
  }
});
