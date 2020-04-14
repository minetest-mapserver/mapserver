#!/bin/bash

# build
docker build . -t mapserver

docker run --name mapserver --rm \
 --network host \
 -v $(pwd)/integration-test-world:/app \
 mapserver &

function cleanup {
	# cleanup
	docker stop mapserver
}

trap cleanup EXIT

bash -c 'while !</dev/tcp/localhost/8080; do sleep 1; done;'

curl http://127.0.0.1:8080/api/tile/0/0/0/0 > tile.png
file tile.png | grep "PNG image data" || exit 1
