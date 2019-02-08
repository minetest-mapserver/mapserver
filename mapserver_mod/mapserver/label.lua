
local update_formspec = function(meta)
	local inv = meta:get_inventory()

	local text = meta:get_string("text")
	local size = meta:get_string("size")
	local direction = meta:get_string("direction")

	meta:set_string("infotext", "Label, Text:" .. text .. ", Size:" .. size .. ", Direction:" .. direction)

	meta:set_string("formspec", "size[8,5;]" ..
		-- col 1
		"field[0,1;4,1;text;Text;" .. text .. "]" ..
		"button_exit[4,1;4,1;save;Save]" ..

		-- col 2
		"field[0,2.5;4,1;size;Size (1-10);" .. size .. "]" ..

		-- col 3
		"field[0,3.5;8,1;direction;Direction (0-360);" .. direction .. "]" ..
		"")

end


minetest.register_node("mapserver:label", {
	description = "Mapserver Label",
	tiles = {
		"mapserver_label.png",
		"mapserver_label.png",
		"mapserver_label.png",
		"mapserver_label.png",
		"mapserver_label.png",
		"mapserver_label.png"
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

		meta:set_string("text", "")
		meta:set_string("direction", "0")
		meta:set_string("size", "1")

		update_formspec(meta)
	end,

	on_receive_fields = function(pos, formname, fields, sender)
		local meta = minetest.get_meta(pos)
		local playername = sender:get_player_name()

		if playername == meta:get_string("owner") then
			-- owner
			if fields.save then
				meta:set_string("text", fields.text)
				meta:set_string("direction", fields.direction)
				meta:set_string("size", fields.size)
			end
		else
			-- non-owner
		end


		update_formspec(meta)
	end


})
