import PlayerOverlay from './overlays/PlayerOverlay.js';
import PoiOverlay from './overlays/PoiOverlay.js';
import ShopOverlay from './overlays/ShopOverlay.js';
import LabelOverlay from './overlays/LabelOverlay.js';
import TrainlineOverlay from './overlays/TrainlineOverlay.js';
import TravelnetOverlay from './overlays/TravelnetOverlay.js';
import BonesOverlay from './overlays/BonesOverlay.js';
import LcdOverlay from './overlays/LcdOverlay.js';
import DigitermOverlay from './overlays/DigitermOverlay.js';
import LuacontrollerOverlay from './overlays/LuacontrollerOverlay.js';
import TechnicAnchorOverlay from './overlays/TechnicAnchorOverlay.js';
import TechnicQuarryOverlay from './overlays/TechnicQuarryOverlay.js';
import TechnicSwitchOverlay from './overlays/TechnicSwitchOverlay.js';
import ProtectorOverlay from './overlays/ProtectorOverlay.js';
import XPProtectorOverlay from './overlays/XPProtectorOverlay.js';
import PrivProtectorOverlay from './overlays/PrivProtectorOverlay.js';
import MissionOverlay from './overlays/MissionOverlay.js';
import MinecartOverlay from './overlays/MinecartOverlay.js';
import ATMOverlay from './overlays/ATMOverlay.js';
import LocatorOverlay from './overlays/LocatorOverlay.js';
import BorderOverlay from './overlays/BorderOverlay.js';
import TrainOverlay from './overlays/TrainOverlay.js';
import TrainsignalOverlay from './overlays/TrainsignalOverlay.js';
import SignOverlay from './overlays/SignOverlay.js';
import AirUtilsPlanesOverlay from "./overlays/AirUtilsPlanesOverlay.js";
import UnifiedMoneyAreaForSaleOverlay from './overlays/UnifiedMoneyAreaForSaleOverlay.js';

export default function(cfg, map, overlays, wsChannel){

  function isDefault(key){
    return cfg.defaultoverlays.indexOf(key) >= 0;
  }

  if (cfg.mapobjects.mapserver_player) {
    overlays.Player = new PlayerOverlay();
    if (isDefault("mapserver_player")) {
      map.addLayer(overlays.Player);
    }
  }

  if (cfg.mapobjects.mapserver_poi) {
    overlays.POI = new PoiOverlay(wsChannel);
    if (isDefault("mapserver_poi")) {
      map.addLayer(overlays.POI);
    }
  }

  if (cfg.mapobjects.smartshop || cfg.mapobjects.fancyvend) {
    overlays.Shop = new ShopOverlay();
    if (isDefault("smartshop") || isDefault("fancyvend")) {
      map.addLayer(overlays.Shop);
    }
  }

  if (cfg.mapobjects.mapserver_label) {
    overlays.Label = new LabelOverlay();
    if (isDefault("mapserver_label")) {
      map.addLayer(overlays.Label);
    }
  }

  if (cfg.mapobjects.mapserver_trainline) {
    overlays.Trainlines = new TrainlineOverlay();
    if (isDefault("mapserver_trainline")) {
      map.addLayer(overlays.Trainlines);
    }
  }

  if (cfg.mapobjects.mapserver_border) {
    overlays.Border = new BorderOverlay();
    if (isDefault("mapserver_border")) {
      map.addLayer(overlays.Border);
    }
  }

  if (cfg.mapobjects.travelnet) {
    overlays.Travelnet = new TravelnetOverlay();
    if (isDefault("travelnet")) {
      map.addLayer(overlays.Travelnet);
    }
  }

  if (cfg.mapobjects.bones) {
    overlays.Bones = new BonesOverlay();
    if (isDefault("bones")) {
      map.addLayer(overlays.Bones);
    }
  }

  if (cfg.mapobjects.digilines) {
    overlays["Digilines LCD"] = new LcdOverlay();
    if (isDefault("digilines")) {
      map.addLayer(overlays["Digilines LCD"]);
    }
  }

  if (cfg.mapobjects.digiterms) {
    overlays.Digiterms = new DigitermOverlay();
    if (isDefault("digiterms")) {
      map.addLayer(overlays.Digiterms);
    }
  }

  if (cfg.mapobjects.luacontroller) {
    overlays["Lua Controller"] = new LuacontrollerOverlay();
    if (isDefault("luacontroller")) {
      map.addLayer(overlays["Lua Controller"]);
    }
  }

  if (cfg.mapobjects.technic_anchor) {
    overlays["Technic Anchor"] = new TechnicAnchorOverlay();
    if (isDefault("technic_anchor")) {
      map.addLayer(overlays["Technic Anchor"]);
    }
  }

  if (cfg.mapobjects.technic_quarry) {
    overlays["Technic Quarry"] = new TechnicQuarryOverlay();
    if (isDefault("technic_quarry")) {
      map.addLayer(overlays["Technic Quarry"]);
    }
  }

  if (cfg.mapobjects.technic_switch) {
    overlays["Technic Switching station"] = new TechnicSwitchOverlay();
    if (isDefault("technic_switch")) {
      map.addLayer(overlays["Technic Switching station"]);
    }
  }

  if (cfg.mapobjects.protector) {
    overlays.Protector = new ProtectorOverlay();
    if (isDefault("protector")) {
      map.addLayer(overlays.Protector);
    }
  }

  if (cfg.mapobjects.xpprotector) {
    overlays["XP Protector"] = new XPProtectorOverlay();
    if (isDefault("xpprotector")) {
      map.addLayer(overlays["XP Protector"]);
    }
  }

  if (cfg.mapobjects.privprotector) {
    overlays["Priv Protector"] = new PrivProtectorOverlay();
    if (isDefault("privprotector")) {
      map.addLayer(overlays["Priv Protector"]);
    }
  }

  if (cfg.mapobjects.mission) {
    overlays.Missions = new MissionOverlay();
    if (isDefault("mission")) {
      map.addLayer(overlays.Missions);
    }
  }

  if (cfg.mapobjects.train) {
    overlays.Trains = new TrainOverlay();

    if (isDefault("train")) {
      map.addLayer(overlays.Trains);
    }
  }

  if (cfg.mapobjects.trainsignal) {
    overlays.Trainsignals = new TrainsignalOverlay();

    if (isDefault("trainsignal")) {
      map.addLayer(overlays.Trainsignals);
    }
  }

  if (cfg.mapobjects.minecart) {
    overlays.Minecart = new MinecartOverlay();
    if (isDefault("minecart")) {
      map.addLayer(overlays.Minecart);
    }
  }

  if (cfg.mapobjects.atm) {
    overlays.ATM = new ATMOverlay();
    if (isDefault("atm")) {
      map.addLayer(overlays.ATM);
    }
  }

  if (cfg.mapobjects.locator) {
    overlays.Locator = new LocatorOverlay();
    if (isDefault("locator")) {
      map.addLayer(overlays.Locator);
    }
  }

  if (cfg.mapobjects.signs) {
    overlays.Signs = new SignOverlay();
    if (isDefault("signs")) {
      map.addLayer(overlays.Signs);
    }
  }

  if (cfg.mapobjects.mapserver_airutils) {
    overlays.Planes = new AirUtilsPlanesOverlay();
    if (isDefault("mapserver_airutils")) {
      map.addLayer(overlays.Planes);
    }
  }

  if (cfg.mapobjects.um_area_forsale) {
    overlays["Area For Sale"] = new UnifiedMoneyAreaForSaleOverlay();
    if (isDefault("um_area_forsale")) {
      map.addLayer(overlays["Area For Sale"]);
    }
  }

}
