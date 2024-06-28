import AbstractIconOverlay from './AbstractIconOverlay.js';
import {HtmlSanitizer} from '../../lib/HtmlSanitizer.js';


export default AbstractIconOverlay.extend({
  initialize: function() {
    AbstractIconOverlay.prototype.initialize.call(this, "poi");
  },

  getIcon: function(obj){
    return L.AwesomeMarkers.icon({
      icon: obj.attributes.icon || "home",
      prefix: "fa",
      markerColor: obj.attributes.color || "blue"
    });
  },

  getMaxDisplayedZoom: function(){
    return 5;
  },

  createPopup: function(poi){
    var innerHTML = "";

    if (poi.attributes.url) {
      innerHTML += "<a href=\"" + HtmlSanitizer.SanitizeHtml(poi.attributes.url) + "\">" +
        "<h4>" + HtmlSanitizer.SanitizeHtml(poi.attributes.name) + "</h4></a>";
    } else {
      innerHTML += "<h4>" + HtmlSanitizer.SanitizeHtml(poi.attributes.name) + "</h4>";
    }

    if (poi.attributes.image) {
      innerHTML += "<div class=\"poi_image_container\">" +
        "<img class=\"poi_image\" src=\"" + HtmlSanitizer.SanitizeHtml(poi.attributes.image) +
        "\" crossorigin=\"anonymous\" referrerpolicy=\"origin-when-cross-origin\"></div>";
    }

    innerHTML += "<hr><b>Owner: </b> " + HtmlSanitizer.SanitizeHtml(poi.attributes.owner) + "<br>";

    return innerHTML;
  }


});
