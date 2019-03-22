
minetest.register_lbm({
	label = "Tileserver->Mapserver poi upgrade",
	name = "tileserver:poi",
	nodenames = {"tileserver:poi"},
	run_at_every_load = true,
	action = function(pos, node)
    minetest.swap_node(pos, { name="mapserver:poi" })
	end
})

minetest.register_lbm({
	label = "Tileserver->Mapserver train upgrade",
	name = "tileserver:train",
	nodenames = {"tileserver:train"},
	run_at_every_load = true,
	action = function(pos, node)
    minetest.swap_node(pos, { name="mapserver:train" })
	end
})
