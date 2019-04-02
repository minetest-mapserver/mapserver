
local update_formspec = function(meta)
	local name = meta:get_string("name")
	local category = meta:get_string("category")
	local url = meta:get_string("url") or ""

	meta:set_string("infotext", "POI: " .. name .. ", " .. category)

	meta:set_string("formspec", "size[8,5;]" ..
		-- col 1
		"field[0,1;4,1;name;Name;" .. name .. "]" ..
		"button_exit[4,1;4,1;save;Save]" ..

		-- col 2
		"field[0,2.5;4,1;category;Category;" .. category .. "]" ..

		-- col 3
		"field[0,3.5;8,1;url;URL;" .. url .. "]" ..
		"")

end


minetest.register_node("mapserver:poi", {
	description = "Mapserver POI",
	tiles = {
		"mapserver_poi.png"
	},
	groups = {cracky=3,oddly_breakable_by_hand=3},
	sounds = default.node_sound_glass_defaults(),
	can_dig = mapserver.can_dig,
	after_place_node = mapserver.after_place_node,

	on_construct = function(pos)
		local meta = minetest.get_meta(pos)

		meta:set_string("name", "<unconfigured>")
		meta:set_string("category", "main")
		meta:set_string("url", "")

		update_formspec(meta)
	end,

	on_receive_fields = function(pos, formname, fields, sender)

		if not mapserver.can_interact(pos, sender) then
			return
		end

		local meta = minetest.get_meta(pos)

		if fields.save then
			meta:set_string("name", fields.name)
			meta:set_string("url", fields.url)
			meta:set_string("category", fields.category)
		end

		update_formspec(meta)
	end
})


if mapserver.enable_crafting then
	minetest.register_craft({
	    output = 'mapserver:poi',
	    recipe = {
				{"", "dye:blue", ""},
				{"default:paper", "default:goldblock", "default:paper"},
				{"", "default:glass", ""}
			}
	})
end
