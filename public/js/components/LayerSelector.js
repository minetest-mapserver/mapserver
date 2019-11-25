import LayerManager from '../LayerManager.js';

function onchange(e){
  const params = m.route.param();
  params.layerId = e.target.value;

  m.route.set("/map/:layerId/:zoom/:lon/:lat", params);
}

export default {
  view: function(){
    // Display layer selector only if there is choice
    if (LayerManager.layers.length <= 1)
      return null;

    const layers = LayerManager.layers.map(layer => m(
      "option",
      { value: layer.id, selected: layer.id == LayerManager.getCurrentLayer().id },
      layer.name
    ));

    return m("select", { class: "form-control", onchange: onchange },layers);
  }
};
