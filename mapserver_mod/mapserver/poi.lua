
-- possible icons: https://fontawesome.com/icons?d=gallery&s=brands,regular,solid&m=free
-- default: "home"

local update_formspec = function(meta)
	local name = meta:get_string("name")
	local icon = meta:get_string("icon") or "home"
	local url = meta:get_string("url") or ""

	meta:set_string("infotext", "POI, name:" .. name .. ", icon:" .. icon)

	meta:set_string("formspec", "size[8,5;]" ..
		-- col 1
		"field[0,1;4,1;name;Name;" .. name .. "]" ..
		"button_exit[4,1;4,1;save;Save]" ..

		-- col 2
		"field[0,2.5;4,1;icon;Icon;" .. icon .. "]" ..

		-- col 3
		"field[0,3.5;8,1;url;URL;" .. url .. "]" ..
		"")

end

local on_receive_fields = function(pos, formname, fields, sender)

	if not mapserver.can_interact(pos, sender) then
		return
	end

	local meta = minetest.get_meta(pos)

	if fields.save then
		meta:set_string("name", fields.name)
		meta:set_string("url", fields.url)
		meta:set_string("icon", fields.icon or "home")
	end

	update_formspec(meta)
end

local register_poi = function(color, dye)
	minetest.register_node("mapserver:poi_" .. color, {
		description = "Mapserver POI (" .. color .. ")",
		tiles = {
			"[combine:16x16:0,0=mapserver_gold_block.png:3,2=mapserver_poi_" .. color .. ".png"
		},
		groups = {cracky=3,oddly_breakable_by_hand=3},
		sounds = default.node_sound_glass_defaults(),
		can_dig = mapserver.can_dig,
		after_place_node = mapserver.after_place_node,

		on_construct = function(pos)
			local meta = minetest.get_meta(pos)

			meta:set_string("name", "<unconfigured>")
			meta:set_string("icon", "home")
			meta:set_string("url", "")

			update_formspec(meta)
		end,

		on_receive_fields = on_receive_fields
	})


	if mapserver.enable_crafting then
		minetest.register_craft({
		    output = 'mapserver:poi_' .. color,
		    recipe = {
					{"", "dye:" .. dye, ""},
					{"default:paper", "default:goldblock", "default:paper"},
					{"", "default:glass", ""}
				}
		})
	end
end

register_poi("blue", "blue")
register_poi("green", "green")
register_poi("orange", "orange")
register_poi("red", "red")
register_poi("purple", "violet")

-- default poi was always blue
minetest.register_alias("mapserver:poi", "mapserver:poi_blue")
