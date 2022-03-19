import App from './app.js';
import routes from './routes.js';

// create router instance
const router = VueRouter.createRouter({
	history: VueRouter.createWebHashHistory(),
	routes: routes
});

// start vue
const app = Vue.createApp(App);
app.use(router);
app.mount("#app");
