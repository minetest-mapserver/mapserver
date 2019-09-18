
const Component = {
  view: function(){
    return m("select", {
      class: "form-control"
        },[
          m("option", { value: "Ground" }, "Ground"),
          m("option", { value: "Sky" }, "Sky")
        ]
    );
  }
};

export default L.Control.extend({
    onAdd: function() {
      var div = L.DomUtil.create('div');
      m.mount(div, Component);
      return div;
    }
});
