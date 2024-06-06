import AbstractIconOverlay from './AbstractIconOverlay.js';
import {HtmlSanitizer} from '../../lib/HtmlSanitizer.js';

var PhonographIcon = L.icon({
  iconUrl: 'pics/phonograph_node_temp.png',

  iconSize:     [16, 16],
  iconAnchor:   [8, 8],
  popupAnchor:  [0, -16]
});

export default AbstractIconOverlay.extend({
  initialize: function() {
    AbstractIconOverlay.prototype.initialize.call(this, "phonograph", PhonographIcon);
  },

  createPopup: function(obj){
    return "<h4>Phonograph</h4><br>" +
      "<b>Now playing:</b> <i>" + HtmlSanitizer.SanitizeHtml(obj.attributes.song_title) + "</i> by " +
      HtmlSanitizer.SanitizeHtml(obj.attributes.song_artist);
  }
});
