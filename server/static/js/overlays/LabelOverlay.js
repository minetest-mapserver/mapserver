/* exported LabelOverlay */
/* globals AbstractIconOverlay: true */

var LabelIcon = L.Icon.extend({
  initialize: function(options) {
    L.Icon.prototype.initialize.call(this, options);
  },

  createIcon: function() {
		var div = document.createElement('div'),
		    options = this.options;

		div.innerHTML = options.html || "";
    div.style.width = "200px";
    div.style.height = "200px";
    div.style.marginLeft = "-100px";
    div.style.marginTop = "-175px";

		return div;
  }
});

var LabelOverlay = AbstractIconOverlay.extend({
  initialize: function(wsChannel, layerMgr) {
    AbstractIconOverlay.prototype.initialize.call(this, wsChannel, layerMgr, "label");
  },

  getIcon: function(lbl){

    var height = 200;
    var width = 200;
    var fontSize = Math.min(lbl.attributes.size, 20);

    const html = `
      <svg height='${height}' width='${width}' text-anchor='middle' font-size='${fontSize}px'>
        <text x='${width/2}' y='${height/2}'
          fill='${lbl.attributes.color}'
          dominant-baseline="central"
          transform="rotate(${lbl.attributes.direction}, 100, 100)">
          ${lbl.attributes.text}
        </text>
      </svg>
    `;

    return new LabelIcon({
      html: html,
      height: height,
      width: width
    });
  },

  createPopup: function(){}
});
