import AbstractIconOverlay from './AbstractIconOverlay.js';
import { HtmlSanitizer } from '../../lib/HtmlSanitizer.js';

export default AbstractIconOverlay.extend({
  initialize: function() {
    AbstractIconOverlay.prototype.initialize.call(this, "label");
  },

  getMaxDisplayedZoom: function(){
    return 5;
  },

  getIcon: function(lbl){

    var zoom = this.map.getZoom();
    var factor = Math.pow(2, 12-zoom);

    var height = 200;
    var width = 200;
    var fontSize = Math.min(lbl.attributes.size, 200) / factor;

    var notVisible = (fontSize < 2 || fontSize > 50);

    const html = `
      <svg height='${height}' width='${width}' text-anchor='middle' style='pointer-events: none;'>
        <text x='${width/2}' y='${height/2}'
          font-size='${fontSize}px'
          fill='${lbl.attributes.color}'
          dominant-baseline="central"
          transform="rotate(${lbl.attributes.direction}, 100, 100)">
          ${HtmlSanitizer.SanitizeHtml(lbl.attributes.text)}
        </text>
      </svg>
    `;

    return L.divIcon({
      html: notVisible ? "" : html,
      bgPos: L.point(100,100),
      className: "mapserver-label-icon"
    });
  },

  createPopup: function(){}
});
