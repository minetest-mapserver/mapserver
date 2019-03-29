-- mapserver http bridge
local has_advtrains = minetest.get_modpath("advtrains")

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
        arrayoutput = explode("=",arrayoutput[4])
        return arrayoutput[1]
end

local http, url, key

function send_stats()

  local data = {
    time = minetest.get_timeofday() * 24000,
    uptime = minetest.get_server_uptime(),
    max_lag = tonumber(get_max_lag()),
    players = {}
  }

  if has_advtrains then
    -- send trains if 'advtrains' mod installed

    data.trains = {}
    for _, train in pairs(advtrains.trains) do
      --print(dump(train))--XXX

      local t = {
        text_outside = train.text_outside,
        text_inside = train.text_inside,
        line = train.line,
        pos = train.last_pos,
        velocity = train.velocity,
        off_track = train.off_track,
        id = train.id,
        wagons = {}
      }

      for _, part in pairs(train.trainparts) do
        local wagon = advtrains.wagons[part]
        if wagon ~= nil then
          table.insert(t.wagons, {
            id = wagon.id,
            type = wagon.type,
            pos_in_train = wagon.pos_in_train,
          })
        end
      end

      table.insert(data.trains, t)
    end

  end

  for _, player in ipairs(minetest.get_connected_players()) do

    local is_hidden = minetest.check_player_privs(player:get_player_name(), {mapserver_hide_player = true})

    local info = {
      name = player:get_player_name(),
      pos = player:get_pos(),
      hp = player:get_hp(),
      breath = player:get_breath(),
      velocity = player:get_player_velocity()
    }

    if not is_hidden then
      table.insert(data.players, info)
    end
  end

  local json = minetest.write_json(data)
  --print(json)--XXX

  http.fetch({
    url = url .. "/api/minetest",
    extra_headers = { "Content-Type: application/json", "Authorization: " .. key },
    timeout = 1,
    post_data = json
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
