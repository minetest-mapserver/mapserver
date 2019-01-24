
mapserver = {
}

local MP = minetest.get_modpath("mapserver")
local has_digiline_mod = minetest.get_modpath("digiline")

dofile(MP.."/poi.lua")
dofile(MP.."/train.lua")

if has_digiline_mod then
	dofile(MP.."/digimessage.lua")
end


print("[OK] Mapserver")
