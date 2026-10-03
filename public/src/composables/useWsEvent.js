import { shallowRef, onMounted, onUnmounted } from 'vue';
import wsChannel from '../WebSocketChannel.js';

// latest payload of a websocket event type, listener is bound to the component lifecycle
export function useWsEvent(type){
  const data = shallowRef(null);
  const listener = d => { data.value = d; };

  onMounted(() => wsChannel.addListener(type, listener));
  onUnmounted(() => wsChannel.removeListener(type, listener));

  return data;
}
