
mapserver = {
	enable_crafting = minetest.settings:get("mapserver.enable_crafting")
}

local MP = minetest.get_modpath("mapserver")
dofile(MP.."/common.lua")
dofile(MP.."/poi.lua")
dofile(MP.."/train.lua")
dofile(MP.."/label.lua")
dofile(MP.."/border.lua")
dofile(MP.."/legacy.lua")
dofile(MP.."/privs.lua")


-- optional mapserver-bridge stuff below

--[[ minetest.conf
secure.http_mods = mapserver
mapserver.url = http://127.0.0.1:8080
mapserver.key = myserverkey
--]]

local http = minetest.request_http_api()

if http then
	local mapserver_url = minetest.settings:get("mapserver.url")
	local mapserver_key = minetest.settings:get("mapserver.key")

	if not mapserver_url then error("mapserver.url is not defined") end
	if not mapserver_key then error("mapserver.key is not defined") end

	print("[Mapserver] starting mapserver-bridge with endpoint: " .. mapserver_url)
	dofile(MP .. "/bridge.lua")
	mapserver.bridge_init(http, mapserver_url, mapserver_key)

else
	print("[Mapserver] bridge not active, additional infos will not be visible on the map")

end


print("[OK] Mapserver")
