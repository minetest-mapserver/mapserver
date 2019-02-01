var WebSocketChannel = function(){
  'use strict';
  var wsUrl = location.protocol.replace("http", "ws") + "//" + location.host + location.pathname.substring(0, location.pathname.lastIndexOf("/")) + "/api/ws";

  var listenerMap = {/* type -> [listeners] */};

  this.addListener = function(type, listener){
    var list = listenerMap[type];
    if (!list){
      list = [];
      listenerMap[type] = list;
    }

    list.push(listener);
  };

  this.connect = function(){
    var ws = new WebSocket(wsUrl);

    ws.onmessage = function(e){
      var event = JSON.parse(e.data);
      //rendered-tile, mapobject-created, mapobjects-cleared

      var listeners = listenerMap[event.type];
      if (listeners){
        listeners.forEach(function(listener){
          listener(event.data);
        });
      }
    }

    ws.onerror = function(){
      //reconnect after some time
      setTimeout(connect, 1000);
    }
  };

};
