
minetest.after(5, function()
    minetest.log("action", "[mapserver_emerge] emerging area")
    local pos1 = { x=0, y=-50, z=0 }
    local pos2 = { x=50, y=50, z=0 }
    minetest.emerge_area(pos1, pos2)
end)