
local update_formspec = function(meta)
	local inv = meta:get_inventory()

	local active = meta:get_int("active") == 1
	local state = "Inactive"

	if active then
		state = "Active"
	end

	local name = meta:get_string("name")
	local category = meta:get_string("category")
	local url = meta:get_string("url") or ""

	meta:set_string("infotext", "POI: " .. name .. ", " .. category .. " (" .. state .. ")")

	meta:set_string("formspec", "size[8,5;]" ..
		-- col 1
		"field[0,1;4,1;name;Name;" .. name .. "]" ..
		"button_exit[4,1;4,1;save;Save]" ..

		-- col 2
		"field[0,2.5;4,1;category;Category;" .. category .. "]" ..
		"button_exit[4,2;4,1;toggle;Toggle]" ..

		-- col 3
		"field[0,3.5;8,1;url;URL;" .. url .. "]" ..
		"")

end


minetest.register_node("mapserver:poi", {
	description = "Mapserver POI",
	tiles = {
		"mapserver_poi.png",
		"mapserver_poi.png",
		"mapserver_poi.png",
		"mapserver_poi.png",
		"mapserver_poi.png",
		"mapserver_poi.png"
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

		meta:set_string("name", "<unconfigured>")
		meta:set_string("category", "main")
		meta:set_string("url", "")
		meta:set_int("active", 0)

		update_formspec(meta)
	end,

	on_receive_fields = function(pos, formname, fields, sender)
		local meta = minetest.get_meta(pos)
		local playername = sender:get_player_name()

		if playername == meta:get_string("owner") then
			-- owner
			if fields.save then
				meta:set_string("name", fields.name)
				meta:set_string("url", fields.url)
				meta:set_string("category", fields.category)
			end

			if fields.toggle then
				if meta:get_int("active") == 1 then
					meta:set_int("active", 0)
				else
					meta:set_int("active", 1)
				end
			end

		else
			-- non-owner
		end


		update_formspec(meta)
	end


})
