
local update_formspec = function(meta)
	local inv = meta:get_inventory()

	local active = meta:get_int("active") == 1
	local state = "Inactive"

	if active then
		state = "Active"
	end

	local channel = meta:get_string("channel")
	local message = meta:get_string("message")

	meta:set_string("infotext", "Digimessage: Channel=" .. channel .. ", Message=" .. message .. " (" .. state .. ")")

	meta:set_string("formspec", "size[8,2;]" ..
		-- col 1
		"field[0,1;8,1;channel;Channel;" .. channel .. "]" ..

		-- col 3
		"button_exit[0,2;4,1;save;Save]" ..
		"button_exit[4,2;4,1;toggle;Toggle]" ..
		"")

end


minetest.register_node("mapserver:digimessage", {
	description = "Mapserver Digiline Message",
	tiles = {
		"tileserver_digimessage.png",
		"tileserver_digimessage.png",
		"tileserver_digimessage.png",
		"tileserver_digimessage.png",
		"tileserver_digimessage.png",
		"tileserver_digimessage.png"
	},
	groups = {cracky=3,oddly_breakable_by_hand=3},
	sounds = default.node_sound_glass_defaults(),

	digiline = {
		receptor = {action = function() end},
		effector = {
			action = function(pos, _, channel, msg)
				local meta = minetest.env:get_meta(pos)
				local set_channel = meta:get_string("channel")

				if channel == set_channel and type(msg) == "string" then
					meta:set_string("message", msg)
				end
			end
		},
	},

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

		meta:set_string("channel", "digimessage")
		meta:set_string("message", "")
		meta:set_int("active", 1)

		update_formspec(meta)
	end,

	on_receive_fields = function(pos, formname, fields, sender)
		local meta = minetest.get_meta(pos)
		local playername = sender:get_player_name()

		if playername == meta:get_string("owner") then
			-- owner
			if fields.save then
				last_line = fields.line
				meta:set_string("channel", fields.channel)
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
