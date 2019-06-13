/* exported LabelOverlay */
/* globals AbstractIconOverlay: true */


var LabelOverlay = AbstractIconOverlay.extend({
  initialize: function(wsChannel, layerMgr) {
    AbstractIconOverlay.prototype.initialize.call(this, wsChannel, layerMgr, "label");
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
      <svg height='${height}' width='${width}' text-anchor='middle'>
        <text x='${width/2}' y='${height/2}'
          font-size='${fontSize}px'
          fill='${lbl.attributes.color}'
          dominant-baseline="central"
          transform="rotate(${lbl.attributes.direction}, 100, 100)">
          ${lbl.attributes.text}
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
