import Map from "./pages/Map.js";


export default [{
	path: "/map/:layerId/:zoom/:lon/:lat", component: Map
},{
	path: "/", redirect: "/map/0/13/0/0"
}];
