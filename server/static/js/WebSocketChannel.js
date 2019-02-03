'use strict';

function WebSocketChannel(){
  this.wsUrl = location.protocol.replace("http", "ws") + "//" + location.host + location.pathname.substring(0, location.pathname.lastIndexOf("/")) + "/api/ws";
  this.listenerMap = {/* type -> [listeners] */};
}

WebSocketChannel.prototype.addListener = function(type, listener){
  var list = this.listenerMap[type];
  if (!list){
    list = [];
    this.listenerMap[type] = list;
  }

  list.push(listener);
};

WebSocketChannel.prototype.removeListener = function(type, listener){
  var list = this.listenerMap[type];
  if (!list){
    return
  }

  this.listenerMap[type] = list.filter(function(l){
    return l != listener;
  });
};

WebSocketChannel.prototype.connect = function(){
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
  }

  ws.onerror = function(){
    //reconnect after some time
    setTimeout(connect, 1000);
  }
};
