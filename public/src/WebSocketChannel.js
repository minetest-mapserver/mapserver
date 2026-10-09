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

    this.listenerMap[type] = list.filter(l => l !== listener);
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

    this.pollingHandle = setInterval(async () => {
        const stats = await getStats();
        if (stats){
          this.emit("minetest-info", stats);
        }
    }, 2000);
  }

  stopPolling(){
    if (this.pollingHandle){
      clearInterval(this.pollingHandle);
      this.pollingHandle = null;
    }
  }

  connect(){
    var ws = new WebSocket(this.wsUrl);

    ws.onopen = () => {
      this.stopPolling();
    };

    ws.onmessage = e => {
      var event;
      try {
        event = JSON.parse(e.data);
      } catch (err) {
        console.error("invalid websocket message", err);
        return;
      }
      //rendered-tile, mapobject-created, mapobjects-cleared, minetest-info
      this.emit(event.type, event.data);
    };

    ws.onerror = () => {
      //fallback to polling stats
      this.startPolling();
    };

    ws.onclose = () => {
      //fallback to polling stats and try to reconnect
      this.startPolling();
      setTimeout(() => this.connect(), 5000);
    };
  }
}

export default new WebSocketChannel();
