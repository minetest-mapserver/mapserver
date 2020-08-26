import layerMgr from '../LayerManager.js';
import html from "./html.js";

export default {
  view: function(vnode){
    var result = vnode.attrs.result;

    function getLayer(obj){
      var layer = layerMgr.getLayerByY(obj.y);
      return layer ? layer.name : "<unknown>";
    }

    function getPos(obj){
      var text = obj.x + "/" + obj.y + "/" + obj.z;

      return html`<span class="badge badge-success">${text}</span>`;
    }

    var rows = result.map(function(obj){

      var row_classes = "";
      var description = obj.type;
      var type = obj.type;

      // train-line result
      if (obj.type == "train"){
        description = html`
          <span>${obj.attributes.station}</span>
          <span class="badge badge-info">${obj.attributes.line}</span>
        `;

        type = html`<i class="fa fa-subway"/>`;
      }

      // travelnet
      if (obj.type == "travelnet"){
        description = html`<span>${obj.attributes.station_name}</span>`;
        type = html`<img src="pics/travelnet_inv.png"/>`;
      }

      // bones
      if (obj.type == "bones"){
        description = html`<span>${obj.attributes.owner}</span>`;
        type = html`<img src="pics/bones_top.png"/>`;
      }

      // label
      if (obj.type == "label"){
        description = html`<span>${obj.attributes.text}</span>`;
        type = html`<img src="pics/mapserver_label.png"/>`;
      }

      // digiterm
      if (obj.type == "digiterm"){
        description = html`<span>${obj.attributes.display_text}</span>`;
        type = html`<img src="pics/digiterms_beige_front.png"/>`;
      }

      // digiline lcd
      if (obj.type == "digilinelcd"){
        description = html`<span>${obj.attributes.text}</span>`;
        type = html`<img src="pics/lcd_lcd.png"/>`;
      }

      // locator
      if (obj.type == "locator"){
        description = html`<span>${obj.attributes.name}</span>`;

        var img = "pics/locator_beacon_level1.png";

        if (obj.attributes.level == "2")
          img = "pics/locator_beacon_level2.png";
        else if (obj.attributes.level == "3")
          img = "pics/locator_beacon_level3.png";

          type = html`<img src=${img}/>`;
      }

      // poi marker
      if (obj.type == "poi"){
        description = html`<span>${obj.attributes.name}</span>`;

        var color = obj.attributes.color || "blue";
        var icon = obj.attributes.icon || "home";

        type = html`<div class="awesome-marker awesome-marker-icon-${color}" style="position: relative;">
          <i class="fa fa-${icon}"/>
        </div>`;
      }

      //shop
      if (obj.type == "shop") {
        if (obj.attributes.stock == 0){
          row_classes += "table-warning";
          type = html`<img src="pics/shop_empty.png"/>`;
        } else {
          type = html`<img src="pics/shop.png"/>`;
        }

        description = html`<span>
          Shop, trading
          <span class="badge badge-primary">
            ${obj.attributes.out_count}
            x
            <i class="fa fa-cart-arrow-down"/>
          </span>
          for
          <span class="badge badge-primary">
            ${obj.attributes.in_count}
            x
            <i class="fa fa-money-bill"/>
          </span>
          <span class="badge badge-info">${obj.attributes.in_item}</span>
          Stock:
          <span class="badge badge-info">${obj.attributes.stock}</span>
        </span>`;
      }

      function onclick(){
        var layer = layerMgr.getLayerByY(obj.y);
        m.route.set(`/map/${layer.id}/${12}/${obj.x}/${obj.z}`);
      }

      return html`<tr class="${row_classes}">
        <td>${type}</td>
        <td>${obj.attributes.owner}</td>
        <td>${getLayer(obj)}</td>
        <td>${getPos(obj)}</td>
        <td>${description}</td>
        <td>
          <button type="button" class="btn btn-secondary" onclick=${onclick}>
            Goto <i class="fas fa-play"/>
          </button>
        </td>
        <td></td>
      </tr>`;
    });

    return html`<table class="table table-striped">
      <thead>
        <tr>
          <th>Type</th>
          <th>Owner</th>
          <th>Layer</th>
          <th>Position</th>
          <th>Description</th>
          <th>Action</th>
        </tr>
      </thead>
      <tbody>
        ${rows}
      </tbody>
    </table>`;
  }
};
