
mapserver.bridge.add_advtrains = function(data)
  data.trains = {}
  for _, train in pairs(advtrains.trains) do

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
