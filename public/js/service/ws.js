
import { get } from '../api/stats.js';

class WebSocketChannel {
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

		function fallbackPolling(){
			get().then(function(stats){
				if (!stats){
					// no stats (yet)
					return;
				}

				var listeners = self.listenerMap["minetest-info"];
				if (listeners){
					listeners.forEach(function(listener){
						listener(stats);
					});
				}
			});
		}

    ws.onerror = function(){
      //fallback to polling stats
      setInterval(fallbackPolling, 2000);
    };
  }
}

export default new WebSocketChannel();
