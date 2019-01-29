var WebSocketChannel = (function(){
  'use strict';

  function connect(){
    var wsUrl = location.protocol.replace("http", "ws") + "//" + location.host + location.pathname.substring(0, location.pathname.lastIndexOf("/")) + "/api/ws";
    var ws = new WebSocket(wsUrl);

    ws.onmessage = function(e){
      var event = JSON.parse(e.data);

      if (event.type == "rendered-tile"){
        RealtimeTileLayer.update(event.data)
      }
    }
  }

  return {
    connect: connect
  };

}());
