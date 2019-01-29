var WebSocketChannel = (function(){
  'use strict';

  var wsUrl = location.protocol.replace("http", "ws") + "//" + location.host + location.pathname.substring(0, location.pathname.lastIndexOf("/")) + "/api/ws";

  function connect(){
    var ws = new WebSocket(wsUrl);

    ws.onmessage = function(e){
      var event = JSON.parse(e.data);

      if (event.type == "rendered-tile"){
        //Update tiles
        RealtimeTileLayer.update(event.data)

      } else if (event.type == "mapobject-created"){
        //TODO
        console.log(event);

      } else if (event.type == "mapobjects-cleared"){
        //TODO
        console.log(event);

      }
    }

    ws.onerror = function(){
      //reconnect after some time
      setTimeout(connect, 1000);
    }
  }

  return {
    connect: connect
  };

}());
