// expose leaflet as a global, the map overlays and plugins reference `L` directly
import L from 'leaflet';
import 'leaflet/dist/leaflet.css';

window.L = L;

export default L;
