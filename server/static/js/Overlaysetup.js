/* jshint undef: false */
/* exported Overlaysetup */

function Overlaysetup(cfg, map, overlays, wsChannel, layerMgr){

    if (cfg.mapobjects.mapserver_player) {
      overlays.Player = new PlayerOverlay(wsChannel, layerMgr);
      if (cfg.defaultoverlays.contains("mapserver_player")) {
        map.addLayer(overlays.Player);
      }
    }

    if (cfg.mapobjects.mapserver_poi) {
      overlays.POI = new PoiOverlay(wsChannel, layerMgr);
      if (cfg.defaultoverlays.contains("mapserver_poi")) {
        map.addLayer(overlays.POI);
      }
    }

    if (cfg.mapobjects.mapserver_label) {
      overlays.Label = new LabelOverlay(wsChannel, layerMgr);
      if (cfg.defaultoverlays.contains("mapserver_label")) {
        map.addLayer(overlays.Label);
      }
    }

    if (cfg.mapobjects.mapserver_trainline) {
      overlays.Trainlines = new TrainlineOverlay(wsChannel, layerMgr);
      if (cfg.defaultoverlays.contains("mapserver_trainline")) {
        map.addLayer(overlays.Trainlines);
      }
    }

    if (cfg.mapobjects.mapserver_border) {
      overlays.Border = new BorderOverlay(wsChannel, layerMgr);
      if (cfg.defaultoverlays.contains("mapserver_border")) {
        map.addLayer(overlays.Border);
      }
    }

    if (cfg.mapobjects.travelnet) {
      overlays.Travelnet = new TravelnetOverlay(wsChannel, layerMgr);
      if (cfg.defaultoverlays.contains("travelnet")) {
        map.addLayer(overlays.Travelnet);
      }
    }

    if (cfg.mapobjects.bones) {
      overlays.Bones = new BonesOverlay(wsChannel, layerMgr);
      if (cfg.defaultoverlays.contains("bones")) {
        map.addLayer(overlays.Bones);
      }
    }

    if (cfg.mapobjects.digilines) {
      overlays["Digilines LCD"] = new LcdOverlay(wsChannel, layerMgr);
      if (cfg.defaultoverlays.contains("digilines")) {
        map.addLayer(overlays["Digilines LCD"]);
      }
    }

    if (cfg.mapobjects.digiterms) {
      overlays.Digiterms = new DigitermOverlay(wsChannel, layerMgr);
      if (cfg.defaultoverlays.contains("digiterms")) {
        map.addLayer(overlays.Digiterms);
      }
    }

    if (cfg.mapobjects.luacontroller) {
      overlays["Lua Controller"] = new LuacontrollerOverlay(wsChannel, layerMgr);
      if (cfg.defaultoverlays.contains("luacontroller")) {
        map.addLayer(overlays["Lua Controller"]);
      }
    }

    if (cfg.mapobjects.technic_anchor) {
      overlays["Technic Anchor"] = new TechnicAnchorOverlay(wsChannel, layerMgr);
      if (cfg.defaultoverlays.contains("technic_anchor")) {
        map.addLayer(overlays["Technic Anchor"]);
      }
    }

    if (cfg.mapobjects.technic_quarry) {
      overlays["Technic Quarry"] = new TechnicQuarryOverlay(wsChannel, layerMgr);
      if (cfg.defaultoverlays.contains("technic_quarry")) {
        map.addLayer(overlays["Technic Quarry"]);
      }
    }

    if (cfg.mapobjects.technic_switch) {
      overlays["Technic Switching station"] = new TechnicSwitchOverlay(wsChannel, layerMgr);
      if (cfg.defaultoverlays.contains("technic_switch")) {
        map.addLayer(overlays["Technic Switching station"]);
      }
    }

    if (cfg.mapobjects.protector) {
      overlays.Protector = new ProtectorOverlay(wsChannel, layerMgr);
      if (cfg.defaultoverlays.contains("protector")) {
        map.addLayer(overlays.Protector);
      }
    }

    if (cfg.mapobjects.xpprotector) {
      overlays["XP Protector"] = new XPProtectorOverlay(wsChannel, layerMgr);
      if (cfg.defaultoverlays.contains("xpprotector")) {
        map.addLayer(overlays["XP Protector"]);
      }
    }

    if (cfg.mapobjects.privprotector) {
      overlays["Priv Protector"] = new PrivProtectorOverlay(wsChannel, layerMgr);
      if (cfg.defaultoverlays.contains("privprotector")) {
        map.addLayer(overlays["Priv Protector"]);
      }
    }

    if (cfg.mapobjects.mission) {
      overlays.Missions = new MissionOverlay(wsChannel, layerMgr);
      if (cfg.defaultoverlays.contains("mission")) {
        map.addLayer(overlays.Missions);
      }
    }

    if (cfg.mapobjects.train) {
      overlays.Trains = new TrainOverlay(wsChannel, layerMgr);
      if (cfg.defaultoverlays.contains("train")) {
        map.addLayer(overlays.Trains);
      }
    }
}
