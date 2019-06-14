
import { getConfig } from './api.js';
import { setup } from './map.js';

getConfig().then(setup);
