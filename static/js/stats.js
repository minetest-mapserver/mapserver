
import WorldStats from './components/WorldStats.js';
import wsChannel from './WebSocketChannel.js';

wsChannel.connect();

wsChannel.addListener("minetest-info", function(info){
	m.render(document.getElementById("app"), WorldStats(info));
});
