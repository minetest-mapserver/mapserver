#!/bin/sh

docker run --rm -it \
	-u root:root \
	-v $(pwd)/minetest.conf:/etc/minetest/minetest.conf \
	-v $(pwd)/worldmods:/root/.minetest/worlds/world/worldmods \
	-v $(pwd)/data:/root/.minetest/worlds/world \
	--network host \
	registry.gitlab.com/minetest/minetest/server:5.2.0
