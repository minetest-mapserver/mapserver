/* exported WorldInfoDisplay */

var worldInfoRender = function(info){

  var timeIcon = m("span", { class: "fa fa-sun", style: "color: orange;" });

  if (info.time < 5500 || info.time > 19000) //0 - 24'000
    timeIcon = m("span", { class: "fa fa-moon", style: "color: blue;" });

  function getHour(){
    return Math.floor(info.time/1000);
  }

  function getMinute(){
    var min = Math.floor((info.time % 1000) / 1000 * 60);
    return min > 10 ? min : "0" + min;
  }

  function getLag(){
    var color = "green";
    if (info.max_lag > 0.8)
      color = "orange";
    else if (info.max_lag > 1.2)
      color = "red";

    return [
      m("span", { class: "fa fa-wifi", style: "color: " + color }),
      parseInt(info.max_lag*1000),
      " ms"
    ];
  }

  function getPlayers(){
    return [
      m("span", { class: "fa fa-users" }),
      info.players ? info.players.length : "0"
    ];
  }

  return m("div", [
    getPlayers(),
    " ",
    getLag(),
    " ",
    m("span", { class: "fa fa-clock" }),
    timeIcon,
    getHour(), ":", getMinute()
  ]);

};

// coord display
var WorldInfoDisplay = L.Control.extend({
    initialize: function(wsChannel, opts) {
        L.Control.prototype.initialize.call(this, opts);
        this.wsChannel = wsChannel;
    },

    onAdd: function() {
      var div = L.DomUtil.create('div', 'leaflet-bar leaflet-custom-display');

      this.wsChannel.addListener("minetest-info", function(info){
        m.render(div, worldInfoRender(info));
      });

      return div;
    }
});
