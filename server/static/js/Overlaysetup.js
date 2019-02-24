'use strict';

function Overlaysetup(cfg, map, overlays, wsChannel, layerMgr){

    if (cfg.mapobjects.mapserver) {
      overlays["Player"] = new PlayerOverlay(wsChannel, layerMgr);
      overlays["POI"] = new PoiOverlay(wsChannel, layerMgr);
      overlays["Labels"] = new LabelOverlay(wsChannel, layerMgr);

      map.addLayer(overlays["Player"]);
      map.addLayer(overlays["POI"]);
      map.addLayer(overlays["Labels"]);
    }

    if (cfg.mapobjects.travelnet) {
      overlays["Travelnet"] = new TravelnetOverlay(wsChannel, layerMgr);
    }

    if (cfg.mapobjects.bones) {
      overlays["Bones"] = new BonesOverlay(wsChannel, layerMgr);
    }

    if (cfg.mapobjects.digilines) {
      overlays["Digilines LCD"] = new LcdOverlay(wsChannel, layerMgr);
    }

    if (cfg.mapobjects.digiterms) {
      overlays["Digiterms"] = new DigitermOverlay(wsChannel, layerMgr);
    }

    if (cfg.mapobjects.luacontroller) {
      overlays["Lua Controller"] = new LuacontrollerOverlay(wsChannel, layerMgr);
    }

    if (cfg.mapobjects.technic) {
      overlays["Technic Anchor"] = new TechnicAnchorOverlay(wsChannel, layerMgr);
      overlays["Technic Quarry"] = new TechnicQuarryOverlay(wsChannel, layerMgr);
      overlays["Technic Switching station"] = new TechnicSwitchOverlay(wsChannel, layerMgr);
    }

    if (cfg.mapobjects.protector) {
      overlays["Protector"] = new ProtectorOverlay(wsChannel, layerMgr);
    }

    if (cfg.mapobjects.mission) {
      overlays["Missions"] = new MissionOverlay(wsChannel, layerMgr);
    }

    if (cfg.mapobjects.train) {
      overlays["Trains"] = new TrainOverlay(wsChannel, layerMgr);
    }
}
