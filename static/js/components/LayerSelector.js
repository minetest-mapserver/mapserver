import LayerManager from '../LayerManager.js';
import html from "./html.js";

function onchange(e){
  const params = m.route.param();
  params.layerId = e.target.value;

  m.route.set("/map/:layerId/:zoom/:lon/:lat", params);
}


const layerOption = layer => html`
  <option value=${layer.id} selected=${layer.id == LayerManager.getCurrentLayer()}>
    ${layer.name}
  </option>
`;

export default {
  view: () => html`<select class="form-control" onchange=${onchange}>
      ${LayerManager.layers.map(layerOption)}
    </select>`
};
