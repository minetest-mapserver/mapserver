-- mapserver http bridge

local http, url, key

function send_stats()

  local data = {}

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
