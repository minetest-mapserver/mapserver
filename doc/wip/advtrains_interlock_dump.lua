minetest.after(2, function()
        local world_path = minetest.get_worldpath()
        local file, err = io.open(world_path .. "advtrains.json", "w")

        local data = advtrains.interlocking.db.save()
        local json = minetest.write_json(data.ts, true)

        file:write(json)
        file:close()
end)
