import { getStats } from './api.js';

class WebSocketChannel {
  constructor(){
    this.wsUrl = window.location.protocol.replace("http", "ws") +
      "//" + window.location.host +
      window.location.pathname.substring(0, window.location.pathname.lastIndexOf("/")) +
      "/api/ws";

    this.listenerMap = {/* type -> [listeners] */};
    this.pollingHandle = null;
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

  emit(type, data){
    var listeners = this.listenerMap[type];
    if (listeners){
      listeners.forEach(listener => listener(data));
    }
  }

  startPolling(){
    if (this.pollingHandle){
      // already polling
      return;
    }

    this.pollingHandle = setInterval(() => {
      getStats()
      .then(stats => {
        if (stats){
          this.emit("minetest-info", stats);
        }
      })
      .catch(() => {
        // ignore, retry on next interval
      });
    }, 2000);
  }

  connect(){
    var ws = new WebSocket(this.wsUrl);

    ws.onmessage = e => {
      var event = JSON.parse(e.data);
      //rendered-tile, mapobject-created, mapobjects-cleared, minetest-info
      this.emit(event.type, event.data);
    };

    ws.onerror = () => {
      //fallback to polling stats
      this.startPolling();
    };
  }
}

export default new WebSocketChannel();
