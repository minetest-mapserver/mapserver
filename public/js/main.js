import App from './app.js';
import routes from './routes.js';
import { get as getConfig } from './api/config.js';
import configStore from './store/config.js';
import ws from './service/ws.js';

function start() {
	// create router instance
	const router = VueRouter.createRouter({
		history: VueRouter.createWebHashHistory(),
		routes: routes
	});

	// start vue
	const app = Vue.createApp(App);
	app.use(router);
	app.mount("#app");
}

// fetch config from server first
getConfig().then(cfg => {
	// copy config to store
	Object.keys(cfg).forEach(k => configStore[k] = cfg[k]);

	// start websocket/polling
	ws.connect();

	// start app
	start();
});