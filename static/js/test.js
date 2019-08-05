var camera, scene, renderer;
var colormapping, geometry;

init();
animate();

function getNodePos(x,y,z){ return x + (y * 16) + (z * 256); }

function drawMapblock(posx,posy,posz){
  m.request("api/mapblock/"+posx+"/"+posy+"/"+posz)
  .then(function(mapblock){
    if (!mapblock)
      return;

    for (var x=0; x<16; x++){
      for (var y=0; y<16; y++){
        for (var z=0; z<16; z++){
          var i = getNodePos(x,y,z);
          var contentId = mapblock.mapdata.contentid[i];
          var nodeName = mapblock.blockmapping[contentId]
          var colorObj = colormapping[nodeName];

          if (!colorObj)
            continue;

          var color = new THREE.Color( colorObj.r/256, colorObj.g/256, colorObj.b/256 );
          var material = new THREE.MeshBasicMaterial( { color: color } );

          var mesh = new THREE.Mesh( geometry, material );
          mesh.position.x = (x*3) + (posx*3*16);
          mesh.position.y = (y*3) + (posy*3*16);
          mesh.position.z = (z*3) + (posz*3*16);
        	scene.add( mesh );


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

	geometry = new THREE.BoxGeometry( 3, 3, 3 );

  m.request("api/colormapping")
  .then(function(_colormapping){
    colormapping = _colormapping;
    for (var x=-4; x<4; x++){
      for (var y=0; y<2; y++){
        for (var z=-4; z<4; z++){
          drawMapblock(x,y,z);
        }
      }
    }

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
