import AbstractIconOverlay from './AbstractIconOverlay.js';
import {HtmlSanitizer} from '../../lib/HtmlSanitizer.js';

export default AbstractIconOverlay.extend({
  initialize: function () {
    AbstractIconOverlay.prototype.initialize.call(this, "um_area_forsale");
  },

  getMaxDisplayedZoom: function () {
    return 8;
  },

  getIcon: function (obj) {
    return L.icon({
      iconUrl: "pics/um_area_forsale_sign_alpha.png",
      iconSize: [32, 32],
      iconAnchor: [16, 16],
      popupAnchor: [0, -16]
    });
  },

  createPopup: function (obj) {
    return "<h4>Area for sale</h4>" +
      "<h5>" + HtmlSanitizer.SanitizeHtml(obj.attributes.description || "No Description") + "</h5><hr>" +
      "<b>Owner:</b> " + HtmlSanitizer.SanitizeHtml(obj.attributes.owner) + "<br>" +
      "<b>Area IDs:</b> " + HtmlSanitizer.SanitizeHtml(obj.attributes.id) + "<br>" +
      "<b>Price:</b> $" + HtmlSanitizer.SanitizeHtml(obj.attributes.price) + "<br>";
  }
});
