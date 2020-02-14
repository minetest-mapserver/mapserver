var camera, scene, renderer;
var geometry = new THREE.BufferGeometry()
	.fromGeometry(new THREE.BoxGeometry(1,1,1));

init();
animate();

function getNodePos(x,y,z){ return x + (y * 16) + (z * 256); }

var colormapping, controls;

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

		if (nodeName == "default:water_source"){
			material.transparent = true;
			material.opacity = 0.5;
		}

    materialCache[nodeName] = material;
  }

  return material;
}

function isNodeHidden(mapblock,x,y,z){
  if (x<=1 || x>=14 || y<=1 || y>=14 || z<=1 || z>=14){
    // not sure, may be visible
    return false;
  }

  function isTransparent(contentId){
    var nodeName = mapblock.blockmapping[contentId];
    return nodeName == "air" || nodeName == "default:water_source";
  }

  if (isTransparent(mapblock.contentid[getNodePos(x-1,y,z)]))
    return false;
  if (isTransparent(mapblock.contentid[getNodePos(x,y-1,z)]))
    return false;
  if (isTransparent(mapblock.contentid[getNodePos(x,y,z-1)]))
    return false;
  if (isTransparent(mapblock.contentid[getNodePos(x+1,y,z)]))
    return false;
  if (isTransparent(mapblock.contentid[getNodePos(x,y+1,z)]))
    return false;
  if (isTransparent(mapblock.contentid[getNodePos(x,y,z+1)]))
    return false;

  return true;
}

function drawMapblock(posx,posy,posz){
  return m.request("api/viewblock/"+posx+"/"+posy+"/"+posz)
  .then(function(mapblock){
    if (!mapblock)
      return;

    if (mapblock.blockmapping.length == 1 && mapblock.blockmapping[0] == "air"){
      return;
    }

    var nodenameGeometriesMap = {}; // nodeName => [geo, geo, ...]

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

          var geo = geometry.clone();
          var matrix = new THREE.Matrix4()
            .makeTranslation(
              x + (posx*16),
              y + (posy*16),
              z + (posz*16)
            );
          geo.applyMatrix4(matrix);

          var list = nodenameGeometriesMap[nodeName];
          if (!list){
            list = [];
            nodenameGeometriesMap[nodeName] = list;
          }

          list.push(geo);
        }
      }
    }

    Object.keys(nodenameGeometriesMap).forEach(function(nodeName){
      var material = getMaterial(nodeName);

      if (material){
        var list = THREE.BufferGeometryUtils.mergeBufferGeometries(nodenameGeometriesMap[nodeName]);
        var mesh = new THREE.Mesh(list, material);
        scene.add( mesh );
      }
    });
  });
}

function init() {

	camera = new THREE.PerspectiveCamera( 45, window.innerWidth / window.innerHeight, 2, 2000 );
	camera.position.z = -150;
	camera.position.x = -150;
	camera.position.y = 100;

	scene = new THREE.Scene();

  var min = -7, max = 7;
  var x = min, y = -1, z = min;

  function increment(){
    x++;
    if (x > max){
      z++;
      x = min;
    }
    if (z > max){
      y++;
      z = min;
    }
  }

  var drawLoop = function(){
    if (y >= 3){
      return;
    }

    drawMapblock(x,y,z)
    .then(function(){
      render();
      increment();
      setTimeout(drawLoop, 50);
    });
  };

  m.request("api/colormapping")
  .then(function(_colormapping){
    colormapping = _colormapping;
    drawLoop();
  });

	renderer = new THREE.WebGLRenderer();
	renderer.setSize( window.innerWidth, window.innerHeight );
	document.body.appendChild( renderer.domElement );

	controls = new THREE.TrackballControls( camera, renderer.domElement );
	controls.rotateSpeed = 2.0;
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
