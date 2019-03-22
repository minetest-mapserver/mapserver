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
        local arrayoutput = explode("=",arrayoutput[4])
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

  --[[
  "trains":[
      {"id":"690096","off_track":false,"pos":{"x":-125,"y":6,"z":-46},"velocity":0},
      {"id":"973462","off_track":false,"pos":{"x":-125,"y":5,"z":-64},"velocity":0}
    ],
  "uptime":2.0000000298023224,
  "wagons":[
    {"id":"980189","pos_in_train":2,"train_id":"690096","type":"advtrains:subway_wagon"},
    {"id":"243215","pos_in_train":2.5,"train_id":"973462","type":"advtrains:engine_japan"}
  ]}

  {"max_lag":0.38148100000000001,"players":[{"breath":11,"hp":20,"name":"BuckarooBanzai","pos":{"x":-114.66400146484375,"y":3.5,"z":-66.122001647949219},"velocity":{"x":0,"y":0,"z":0}}],"time":18530.00020980835,"trains":[{"id":"261557","off_track":true,"pos":{"x":-125,"y":5,"z":-56},"velocity":0,"wagons":[{"id":"423954","pos_in_train":6,"type":"advtrains:subway_wagon"},{"id":"635309","pos_in_train":10,"type":"advtrains:subway_wagon"},{"id":"253216","pos_in_train":2,"type":"advtrains:subway_wagon"}]},{"id":"690096","off_track":false,"pos":{"x":-125,"y":6,"z":-46},"velocity":0,"wagons":[{"id":"980189","pos_in_train":2,"type":"advtrains:subway_wagon"}]}],"uptime":235.100003503263}


  --]]

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

    local is_hidden = minetest.check_player_privs(player:get_player_name(), {mapserver_hide_player = true}) then

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
