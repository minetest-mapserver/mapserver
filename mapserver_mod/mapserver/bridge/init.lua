local MP = minetest.get_modpath("mapserver")
dofile(MP .. "/bridge/defaults.lua")
dofile(MP .. "/bridge/players.lua")
dofile(MP .. "/bridge/advtrains.lua")
dofile(MP .. "/bridge/minecart.lua")


-- mapserver http bridge
local has_advtrains = minetest.get_modpath("advtrains")
local has_minecart = minetest.get_modpath("minecart")


local http, url, key

function send_stats()
  local t0 = minetest.get_us_time()

  -- data to send to mapserver
  local data = {}

  mapserver.bridge.add_players(data)
  mapserver.bridge.add_defaults(data)

  if has_minecart then
    -- send minecarts positions if mod is installed
    mapserver.bridge.add_minecart(data)
  end

  if has_advtrains then
    -- send trains if 'advtrains' mod installed
    mapserver.bridge.add_advtrains(data)
  end


  local json = minetest.write_json(data)
  --print(json)--XXX

  local t1 = minetest.get_us_time()
  local process_time = t1 - t0
  if process_time > 10000 then
    minetest.log("warning", "[mapserver-bridge] processing took " .. process_time .. " us")
  end

  http.fetch({
    url = url .. "/api/minetest",
    extra_headers = { "Content-Type: application/json", "Authorization: " .. key },
    timeout = 1,
    post_data = json
  }, function(res)

    local t2 = minetest.get_us_time()
    local post_time = t2 - t1
    if post_time > 1000000 then -- warn if over a second
      minetest.log("warning", "[mapserver-bridge] post took " .. post_time .. " us")
    end

    -- TODO: error-handling
    minetest.after(2, send_stats)
  end)

end

function mapserver.bridge_init(_http, _url, _key)
  http = _http
  url = _url
  key = _key

  minetest.after(2, send_stats)
end
