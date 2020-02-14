var camera, scene, renderer;
var geometry;

init();
animate();

function getNodePos(x,y,z){ return x + (y * 16) + (z * 256); }

var colormapping, controls;
var nodeCount = 0;

var materialCache = {};
function getMaterial(nodeName){
  var material = materialCache[nodeName];
  if (!material) {
    var colorObj = colormapping[nodeName];

    if (!colorObj){
      return;
    }

    var color = new THREE.Color( colorObj.r/256, colorObj.g/256, colorObj.b/256 );
    material = new THREE.MeshBasicMaterial( { color: color } );
    materialCache[nodeName] = material;
  }

  return material;
}

function isNodeHidden(mapblock,x,y,z){
  if (x==0 && x>=15 && y==0 && y>=15 && z==0 && z>=15){
    // not sure, may be visible
    return false;
  }

  function isTransparent(contentId){
    var nodeName = mapblock.blockmapping[contentId];
    return nodeName == "air";
  }

  if (isTransparent(mapblock[getNodePos(x-1,y,z)]))
    return false;
  if (isTransparent(mapblock[getNodePos(x,y-1,z)]))
    return false;
  if (isTransparent(mapblock[getNodePos(x,y,z-1)]))
    return false;
  if (isTransparent(mapblock[getNodePos(x+1,y,z)]))
    return false;
  if (isTransparent(mapblock[getNodePos(x,y+1,z)]))
    return false;
  if (isTransparent(mapblock[getNodePos(x,y,z+1)]))
    return false;

}

function drawMapblock(posx,posy,posz){
  return m.request("api/viewblock/"+posx+"/"+posy+"/"+posz)
  .then(function(mapblock){
    if (!mapblock)
      return;

    for (var x=0; x<16; x++){
      for (var y=0; y<16; y++){
        for (var z=0; z<16; z++){
          if (isNodeHidden(mapblock, x,y,z)){
            //skip hidden node
            continue;
          }

          var i = getNodePos(x,y,z);
          var contentId = mapblock.contentid[i];
          var nodeName = mapblock.blockmapping[contentId];

          var material = getMaterial(nodeName);

          if (material) {
            var mesh = new THREE.Mesh( geometry, material );
            mesh.position.x = (x*1) + (posx*1*16);
            mesh.position.y = (y*1) + (posy*1*16);
            mesh.position.z = (z*1) + (posz*1*16);
          	scene.add( mesh );
            nodeCount++;
          }
        }
      }
    }
  });
}

function init() {

	camera = new THREE.PerspectiveCamera( 75, window.innerWidth / window.innerHeight, 2, 2000 );
	camera.position.z = 30;
	camera.position.y = 10;

	scene = new THREE.Scene();

	geometry = new THREE.BoxGeometry( 1, 1, 1 );

  m.request("api/colormapping")
  .then(function(_colormapping){
    colormapping = _colormapping;
    var drawPromises = [];

    for (var x=-12; x<3; x++){
      for (var y=0; y<2; y++){
        for (var z=-4; z<2; z++){
          drawPromises.push(drawMapblock(x,y,z));
        }
      }
    }

    return Promise.all(drawPromises);
  })
  .then(function(){
    console.log("Node-count: " + nodeCount);
  });

	renderer = new THREE.WebGLRenderer();
	renderer.setSize( window.innerWidth, window.innerHeight );
	document.body.appendChild( renderer.domElement );

	controls = new THREE.TrackballControls( camera, renderer.domElement );
	controls.rotateSpeed = 1.0;
	controls.zoomSpeed = 1.2;
	controls.panSpeed = 0.8;

	controls.noZoom = false;
	controls.noPan = false;

	controls.staticMoving = true;
	controls.dynamicDampingFactor = 0.3;

	controls.keys = [ 65, 83, 68 ];

	controls.addEventListener( 'change', render );

	render();
}

function render(){
	renderer.render( scene, camera );
}

function animate() {
	requestAnimationFrame( animate );
	controls.update();
}
