-- mapserver http bridge


local function explode(sep, input)
        local t={}
        local i=0
        for k in string.gmatch(input,"([^"..sep.."]+)") do
            t[i]=k
            i=i+1
        end
        return t
end
local function get_max_lag()
        local arrayoutput = explode(", ",minetest.get_server_status())
        local arrayoutput = explode("=",arrayoutput[4])
        return arrayoutput[1]
end

local http, url, key

function send_stats()

  local data = {
    time = minetest.get_timeofday() * 24000,
    uptime = minetest.get_server_uptime(),
    max_lag = get_max_lag(),
    players = {}
  }

  for _, player in ipairs(minetest.get_connected_players()) do
    local info = {
      name = player:get_player_name(),
      pos = player:get_pos(),
      hp = player:get_hp(),
      breath = player:get_breath(),
      velocity = player:get_player_velocity()
    }

    table.insert(data.players, player)
  end

  http.fetch({
    url = url,
    extra_headers = { "Content-Type: application/json", "Authorization: " .. key },
    timeout = 1,
    post_data = minetest.write_json(data)
  }, function(res)
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
