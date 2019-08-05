minetest.after(2, function()
        local world_path = minetest.get_worldpath()
        local file, err = io.open(world_path .. "advtrains_tcbs.json", "w")

        local data = advtrains.interlocking.db.save()
	local tmp = {}

	for _, entry in pairs(data.tcbs) do
		local tcb = entry[1]
		-- print(dump(tcb))
		if tcb.signal then
			table.insert(tmp, {
				signal = tcb.signal,
				aspect = tcb.aspect,
				signal_name = tcb.signal_name
			})
		end
	end

        local json, err = minetest.write_json(tmp, true)

	if err then
		error(err)
	end

        file:write(json)
        file:close()
end)
