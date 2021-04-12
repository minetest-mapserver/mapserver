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

    if (poi.attributes.url)
    {
      return "<a href=\"" + HtmlSanitizer.SanitizeHtml(poi.attributes.url) + "\">" +
      "<h4>" + HtmlSanitizer.SanitizeHtml(poi.attributes.name) + "</h4></a><hr>" +
      "<b>Owner: </b> " + HtmlSanitizer.SanitizeHtml(poi.attributes.owner) + "<br>";
    }
    else
    {
      return "<h4>" + HtmlSanitizer.SanitizeHtml(poi.attributes.name) + "</h4><hr>" +
      "<b>Owner: </b> " + HtmlSanitizer.SanitizeHtml(poi.attributes.owner) + "<br>";
    }
  }


});
