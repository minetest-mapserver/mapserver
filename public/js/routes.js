import MapPage from "./pages/MapPage.js";


export default [{
	path: "/map/:layerId/:zoom/:lon/:lat", component: MapPage
},{
	path: "/", redirect: "/map/0/13/0/0"
}];
