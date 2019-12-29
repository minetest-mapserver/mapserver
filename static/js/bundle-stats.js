(function(f){typeof define==='function'&&define.amd?define(f):f();}((function(){'use strict';function WorldStats(info){

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

}class WebSocketChannel {
  constructor(){
    this.wsUrl = window.location.protocol.replace("http", "ws") +
      "//" + window.location.host +
      window.location.pathname.substring(0, window.location.pathname.lastIndexOf("/")) +
      "/api/ws";

    this.listenerMap = {/* type -> [listeners] */};
  }

  addListener(type, listener){
    var list = this.listenerMap[type];
    if (!list){
      list = [];
      this.listenerMap[type] = list;
    }

    list.push(listener);
  }

  removeListener(type, listener){
    var list = this.listenerMap[type];
    if (!list){
      return;
    }

    this.listenerMap[type] = list.filter(l => l != listener);
  }

  connect(){
    var ws = new WebSocket(this.wsUrl);
    var self = this;

    ws.onmessage = function(e){
      var event = JSON.parse(e.data);
      //rendered-tile, mapobject-created, mapobjects-cleared

      var listeners = self.listenerMap[event.type];
      if (listeners){
        listeners.forEach(function(listener){
          listener(event.data);
        });
      }
    };

    ws.onerror = function(){
      //reconnect after some time
      setTimeout(self.connect.bind(self), 1000);
    };
  }
}

var wsChannel = new WebSocketChannel();wsChannel.connect();

wsChannel.addListener("minetest-info", function(info){
	m.render(document.getElementById("app"), WorldStats(info));
});})));//# sourceMappingURL=bundle-stats.js.map
