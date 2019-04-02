
local update_formspec = function(meta)
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
		"mapserver_label.png"
	},
	groups = {cracky=3,oddly_breakable_by_hand=3},
	sounds = default.node_sound_glass_defaults(),
	can_dig = mapserver.can_dig,
	after_place_node = mapserver.after_place_node,

	on_construct = function(pos)
		local meta = minetest.get_meta(pos)

		meta:set_string("text", "")
		meta:set_string("direction", "0")
		meta:set_string("size", "1")

		update_formspec(meta)
	end,

	on_receive_fields = function(pos, formname, fields, sender)

		if not mapserver.can_interact(pos, sender) then
			return
		end

		local meta = minetest.get_meta(pos)

		if fields.save then
			meta:set_string("text", fields.text)
			meta:set_string("direction", fields.direction)
			meta:set_string("size", fields.size)
		end

		update_formspec(meta)
	end


})
