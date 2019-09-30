import layerMgr from '../LayerManager.js';

export default {
  view: function(vnode){
    var result = vnode.attrs.result;

    function getLayer(obj){
      var layer = layerMgr.getLayerByY(obj.y);
      return layer ? layer.name : "<unknown>";
    }

    function getPos(obj){
      var text = obj.x + "/" + obj.y + "/" + obj.z;

      return m("span", {class:"badge badge-success"}, text);
    }

    var rows = result.map(function(obj){

      var row_classes = "";
      var description = obj.type;
      var type = obj.type;

      // train-line result
      if (obj.type == "train"){
        description = [
          m("span", obj.attributes.station),
          " ",
          m("span", {class:"badge badge-info"}, obj.attributes.line)
        ];

        type = m("i", { class: "fa fa-subway" });
      }

      // travelnet
      if (obj.type == "travelnet"){
        description = m("span", obj.attributes.station_name);
        type = m("img", { src: "pics/travelnet_inv.png" });
      }

      // bones
      if (obj.type == "bones"){
        description = m("span", obj.attributes.owner);
        type = m("img", { src: "pics/bones_top.png" });
      }

      // label
      if (obj.type == "label"){
        description = m("span", obj.attributes.text);
        type = m("img", { src: "pics/mapserver_label.png" });
      }

      // digiterm
      if (obj.type == "digiterm"){
        description = m("span", obj.attributes.display_text);
        type = m("img", { src: "pics/digiterms_beige_front.png" });
      }

      // digiline lcd
      if (obj.type == "digilinelcd"){
        description = m("span", obj.attributes.text);
        type = m("img", { src: "pics/lcd_lcd.png" });
      }

      // locator
      if (obj.type == "locator"){
        description = m("span", obj.attributes.name);

        var img = "pics/locator_beacon_level1.png";

        if (obj.attributes.level == "2")
          img = "pics/locator_beacon_level2.png";
        else if (obj.attributes.level == "3")
          img = "pics/locator_beacon_level3.png";

        type = m("img", { src: img });
      }

      // poi marker
      if (obj.type == "poi"){
        description = m("span", obj.attributes.name);

        var color = obj.attributes.color || "blue";
        var icon = obj.attributes.icon || "home";

        type = m("div", { style: "position: relative", class: "awesome-marker awesome-marker-icon-" + color }, [
          m("i", { class: "fa fa-" + icon })
        ]);
      }

      //shop
      if (obj.type == "shop") {
        if (obj.attributes.stock == 0){
          row_classes += "table-warning";
          type = m("img", { src: "pics/shop_empty.png" });
        } else {
          type = m("img", { src: "pics/shop.png" });
        }

        description = m("span", [
          "Shop, trading ",
          m("span", {class:"badge badge-primary"},
            obj.attributes.out_count,
            "x",
            m("i", {class:"fa fa-cart-arrow-down"})
          ),
          m("span", {class:"badge badge-info"}, obj.attributes.out_item),
          " for ",
          m("span", {class:"badge badge-primary"},
            obj.attributes.in_count,
            "x",
            m("i", {class:"fa fa-money-bill"})
          ),
          m("span", {class:"badge badge-info"},  obj.attributes.in_item),
          " Stock: ",
          m("span", {class:"badge badge-info"}, obj.attributes.stock)
        ]);
      }

      function onclick(){
        var layer = layerMgr.getLayerByY(obj.y);
        m.route.set(`/map/${layer.id}/${12}/${obj.x}/${obj.z}`);
      }

      return m("tr", {"class": row_classes}, [
        m("td", type),
        m("td", obj.attributes.owner),
        m("td", getLayer(obj)),
        m("td", getPos(obj)),
        m("td", description),
        m("button[type=button]", {class: "btn btn-secondary", onclick: onclick }, [
          "Goto ",
          m("i", { class: "fas fa-play" })
        ])
      ]);
    });

    return m("table", {class:"table table-striped"}, [
      m("thead", [
        m("tr", [
          m("th", "Type"),
          m("th", "Owner"),
          m("th", "Layer"),
          m("th", "Position"),
          m("th", "Description"),
          m("th", "Action")
        ])
      ]),
      m("tbody", rows)
    ]);
  }
};
