
local update_formspec = function(meta)
	local text = meta:get_string("text")
	local size = meta:get_string("size")
	local direction = meta:get_string("direction")
	local color = meta:get_string("color") or "rgb(0,0,0)"

	meta:set_string("infotext", "Label, Text:" .. text .. ", Size:" .. size .. ", Direction:" .. direction)

	meta:set_string("formspec", "size[8,6;]" ..
		-- col 1
		"field[0,1;4,1;text;Text;" .. text .. "]" ..
		"button_exit[4,1;4,1;save;Save]" ..

		-- col 2
		"field[0,2.5;4,1;size;Size (1-10);" .. size .. "]" ..

		-- col 3
		"field[0,3.5;8,1;direction;Direction (0-360);" .. direction .. "]" ..

		-- col 4
		"field[0,4.5;8,1;color;Color;" .. color .. "]" ..

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
		meta:set_string("color", "rgb(0,0,0)")

		update_formspec(meta)
	end,

	on_receive_fields = function(pos, formname, fields, sender)

		if not mapserver.can_interact(pos, sender) then
			return
		end

		local meta = minetest.get_meta(pos)

		if fields.save then
			meta:set_string("color", fields.color)
			meta:set_string("text", fields.text)
			meta:set_string("direction", fields.direction)
			meta:set_string("size", fields.size)
		end

		update_formspec(meta)
	end
})

if mapserver.enable_crafting then
	minetest.register_craft({
	    output = 'mapserver:label',
	    recipe = {
				{"", "default:paper", ""},
				{"default:paper", "default:goldblock", "default:paper"},
				{"", "default:glass", ""}
			}
	})
end
