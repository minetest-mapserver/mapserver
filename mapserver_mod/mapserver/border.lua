
local last_index = 0
local last_name = ""

local update_formspec = function(meta)
	local name = meta:get_string("name")
	local index = meta:get_string("index")

	meta:set_string("infotext", "Border: Name=" .. name .. ", Index=" .. index)

	meta:set_string("formspec", "size[8,3;]" ..
		-- col 1
		"field[0,1;4,1;name;Name;" .. name .. "]" ..
		"button_exit[4,1;4,1;save;Save]" ..

		-- col 2
		"field[4,2.5;4,1;index;Index;" .. index .. "]" ..
		"")

end


minetest.register_node("mapserver:border", {
	description = "Mapserver Border",
	tiles = {
		"mapserver_border.png"
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

		meta:set_string("name", last_name)
		meta:set_int("index", last_index)

		update_formspec(meta)
	end,

	on_receive_fields = function(pos, formname, fields, sender)
		local meta = minetest.get_meta(pos)
		local playername = sender:get_player_name()

		if playername == meta:get_string("owner") then
			-- owner
			if fields.save then
				last_name = fields.name
				meta:set_string("name", fields.name)
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
