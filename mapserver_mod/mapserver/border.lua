
local last_index = 0
local last_name = ""

local update_formspec = function(meta)
	local name = meta:get_string("name")
	local index = meta:get_string("index")
	local color = meta:get_string("color") or "rgb(10,10,200)"

	meta:set_string("infotext", "Border: Name=" .. name .. ", Index=" .. index)

	meta:set_string("formspec", "size[8,4;]" ..
		-- col 1
		"field[0,1;4,1;name;Name;" .. name .. "]" ..
		"button_exit[4,1;4,1;save;Save]" ..

		-- col 2
		"field[4,2.5;4,1;index;Index;" .. index .. "]" ..

		-- col 3
		"field[4,3.5;4,1;color;Color;" .. color .. "]" ..
		"")

end


minetest.register_node("mapserver:border", {
	description = "Mapserver Border",
	tiles = {
		"mapserver_border.png"
	},
	groups = {cracky=3,oddly_breakable_by_hand=3},
	sounds = default.node_sound_glass_defaults(),
	can_dig = mapserver.can_dig,
	after_place_node = mapserver.after_place_node,

	on_construct = function(pos)
		local meta = minetest.get_meta(pos)

		last_index = last_index + 5

		meta:set_string("color", "rgb(10,10,200)")
		meta:set_string("name", last_name)
		meta:set_int("index", last_index)

		update_formspec(meta)
	end,

	on_receive_fields = function(pos, formname, fields, sender)

		if not mapserver.can_interact(pos, sender) then
			return
		end

		local meta = minetest.get_meta(pos)

		if fields.save then
			last_name = fields.name
			meta:set_string("name", fields.name)
			meta:set_string("color", fields.color)
			local index = tonumber(fields.index)
			if index ~= nil then
				last_index = index
				meta:set_int("index", index)
			end
		end

		update_formspec(meta)
	end
})

if mapserver.enable_crafting then
	minetest.register_craft({
	    output = 'mapserver:border',
	    recipe = {
				{"", "default:steel_ingot", ""},
				{"default:paper", "default:goldblock", "default:paper"},
				{"", "default:glass", ""}
			}
	})
end
