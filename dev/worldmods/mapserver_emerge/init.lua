

local function place_mapobjects()
    minetest.log("action", "[mapserver_emerge] placing mapobjects")

    local pos = vector.new(0,20,0)
    minetest.set_node(pos, { name = "mapserver:poi_blue" })
    local meta = minetest.get_meta(pos)
    meta:set_string("name", "Blue POI")
    meta:set_string("icon", "home")
    meta:set_string("url", "https://github.com/luanti-org/luanti/blob/master/doc/lua_api.md")
    meta:set_string("image", "")
end

minetest.after(5, function()
    minetest.log("action", "[mapserver_emerge] emerging area")
    local pos1 = { x=0, y=-50, z=0 }
    local pos2 = { x=50, y=50, z=0 }
    minetest.emerge_area(pos1, pos2, function(_, _, calls_remaining)
        if calls_remaining == 0 then
            place_mapobjects()
        end
    end)
end)