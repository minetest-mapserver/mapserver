
mapserver.can_dig = function(pos, player)
  local meta = minetest.get_meta(pos)
  local owner = meta:get_string("owner")

  return player and player:get_player_name() == owner
end

mapserver.after_place_node = function(pos, placer)
  local meta = minetest.get_meta(pos)
  meta:set_string("owner", placer:get_player_name() or "")
end

mapserver.can_interact = function(pos, player)
  local meta = minetest.get_meta(pos)
  local owner = meta:get_string("owner")
  local playername = player:get_player_name()

  if playername == owner then
    return true
  end

  if minetest.check_player_privs(playername, {protection_bypass = true}) then
    return true
  end

  return false
end
