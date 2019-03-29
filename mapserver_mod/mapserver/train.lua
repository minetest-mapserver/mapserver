
local last_index = 0
local last_line = ""

local update_formspec = function(meta)
	local line = meta:get_string("line")
	local station = meta:get_string("station")
	local index = meta:get_string("index")

	meta:set_string("infotext", "Train: Line=" .. line .. ", Station=" .. station)

	meta:set_string("formspec", "size[8,3;]" ..
		-- col 1
		"field[0,1;4,1;line;Line;" .. line .. "]" ..
		"button_exit[4,1;4,1;save;Save]" ..

		-- col 2
		"field[0,2.5;4,1;station;Station;" .. station .. "]" ..
		"field[4,2.5;4,1;index;Index;" .. index .. "]"
	)

end


minetest.register_node("mapserver:train", {
	description = "Mapserver Train",
	tiles = {
		"mapserver_train.png"
	},
	groups = {cracky=3,oddly_breakable_by_hand=3},
	sounds = default.node_sound_glass_defaults(),

	can_dig = function(pos, player)
		local meta = minetest.env:get_meta(pos)
		local owner = meta:get_string("owner")

		return player and player:get_player_name() == owner
	end,

	after_place_node = function(pos, placer)
		local meta = minetest.get_meta(pos)
		meta:set_string("owner", placer:get_player_name() or "")
	end,

	on_construct = function(pos)
		local meta = minetest.get_meta(pos)

		last_index = last_index + 5

		meta:set_string("station", "")
		meta:set_string("line", last_line)
		meta:set_int("index", last_index)

		update_formspec(meta)
	end,

	on_receive_fields = function(pos, formname, fields, sender)
		local meta = minetest.get_meta(pos)
		local playername = sender:get_player_name()

		if playername == meta:get_string("owner") then
			-- owner
			if fields.save then
				last_line = fields.line
				meta:set_string("line", fields.line)
				meta:set_string("station", fields.station)
				local index = tonumber(fields.index)
				if index ~= nil then
					last_index = index
					meta:set_int("index", index)
				end
			end
		end


		update_formspec(meta)
	end


})
