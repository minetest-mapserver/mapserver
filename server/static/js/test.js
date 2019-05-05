var camera, scene, renderer;

init();
animate();

function init() {

	camera = new THREE.PerspectiveCamera( 75, window.innerWidth / window.innerHeight, 2, 2000 );
	camera.position.z = 30;
	camera.position.y = 10;

	scene = new THREE.Scene();

	var geometry = new THREE.BoxGeometry( 3, 3, 3 );

	//var texture = new THREE.TextureLoader().load( 'textures/technic_water_mill_top_active.png', render );
	//var material = new THREE.MeshBasicMaterial( { map: texture } );
  var material = new THREE.MeshBasicMaterial( { color: 0xff0000, opacity: 0.5, transparent: true } );
  var colormapping;

  m.request("api/colormapping")
  .then(function(_colormapping){
    colormapping = _colormapping;
    console.log(colormapping);
    return m.request("api/mapblock/0/0/0")
  })
  .then(function(mapblock){
    //console.log(mapblock);
    function getNodePos(x,y,z){ return x + (y * 16) + (z * 256); }

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
          mesh.position.x = x*3;
          mesh.position.y = y*3;
          mesh.position.z = z*3;
        	scene.add( mesh );


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
