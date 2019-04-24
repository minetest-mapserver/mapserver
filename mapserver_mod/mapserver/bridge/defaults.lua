
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


mapserver.bridge.add_defaults = function(data)
  data.time = minetest.get_timeofday() * 24000
  data.uptime = minetest.get_server_uptime()
  data.max_lag = tonumber(get_max_lag())

end
