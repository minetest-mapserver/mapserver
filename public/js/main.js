import { getConfig } from './api.js';

getConfig()
.then(cfg => {
  console.log(cfg);
})
.catch(e => {
  document.getElementById("app").innerHTML = e;
});
