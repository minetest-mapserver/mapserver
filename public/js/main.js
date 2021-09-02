import { getConfig } from './api.js';
import routes from './routes.js';

getConfig()
.then(cfg => {
  console.log(cfg);

  	// create router instance
	const router = new VueRouter({
	  routes: routes
	});

  	// start vue
	new Vue({
	  el: "#app",
	  router: router
	});
})
.catch(e => {
  document.getElementById("app").innerHTML = e;
});
