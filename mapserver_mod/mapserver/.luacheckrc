unused_args = false
allow_defined_top = true

globals = {
	"mapserver"
}

read_globals = {
	-- Stdlib
	string = {fields = {"split"}},
	table = {fields = {"copy", "getn"}},

	-- Minetest
	"minetest",
	"vector", "ItemStack",
	"dump",

	-- Deps
	"unified_inventory", "default", "advtrains",
	"minecart",

	-- optional mods
	"xban"
}
