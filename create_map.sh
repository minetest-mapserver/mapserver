#!/bin/sh
# creates an example/test map.sqlite

MTDIR=/tmp/mt
WORLDDIR=${MTDIR}/worlds/world
WORLDMODDIR=${WORLDDIR}/worldmods

rm -rf ${WORLDDIR}
mkdir -p ${WORLDMODDIR}/mapgen


cat <<EOF > ${WORLDMODDIR}/mapgen/init.lua
local function execute_mapgen(callback)
  local pos1 = { x=-100, y=-100, z=-100 }
	local pos2 = { x=100, y=100, z=100 }
	minetest.emerge_area(pos1, pos2, function(blockpos, _, calls_remaining)
		minetest.log("action", "Emerged: " .. minetest.pos_to_string(blockpos))
		if calls_remaining > 0 then
			return
		end

    callback()
  end)
end

minetest.after(1, function()
  execute_mapgen(function()
    minetest.request_shutdown("success")
  end)
end)

EOF

chmod 777 ${MTDIR} -R
docker run --rm -i \
	-v ${CFG}:/etc/minetest/minetest.conf:ro \
  -v ${MTDIR}:/var/lib/minetest/.minetest \
	registry.gitlab.com/minetest/minetest/server:5.0.1

cp ${WORLDDIR}/map.sqlite .
