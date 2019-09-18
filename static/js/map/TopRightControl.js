import SearchInput from '../components/SearchInput.js';
import LayerSelector from '../components/LayerSelector.js';
import config from '../config.js';

const Component = {
  view: function(){
    const cfg = config.get();

    return m("div", [
      cfg.enablesearch ? m(SearchInput) : null,
      m(LayerSelector)
    ])
  }
}


export default L.Control.extend({
    initialize: function(wsChannel, opts) {
        L.Control.prototype.initialize.call(this, opts);
    },

    onAdd: function() {
      var div = L.DomUtil.create('div');
      m.mount(div, Component);
      return div;
    }
});
