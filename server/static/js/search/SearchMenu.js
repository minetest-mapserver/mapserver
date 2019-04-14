
var SearchResult = {
  view: function(vnode){

    function getLayer(obj){
      var layer = layerMgr.getLayerByY(obj.y);
      return layer ? layer.name : "<unknown>";
    }

    function getPos(obj){
      var layer = layerMgr.getLayerByY(obj.y);
      var link = (layer ? layer.id : "0") + "/" + obj.x + "/" + obj.z + "/" + 12;
      var text = obj.x + "/" + obj.y + "/" + obj.z;

      return m("a", { href: "#" + link }, text);
    }

    var rows = SearchStore.result.map(function(obj){
      return m("tr", [
        m("td", obj.type),
        m("td", obj.attributes.owner),
        m("td", getLayer(obj)),
        m("td", getPos(obj)),
        m("td", "stuff")
      ]);
    });

    return m("table", {class:"table"}, [
      m("thead", [
        m("tr", [
          m("th", "Type"),
          m("th", "Owner"),
          m("th", "Layer"),
          m("th", "Position"),
          m("th", "Description")
        ])
      ]),
      m("tbody", rows)
    ]);
  }
}

var SearchMenu = {
  view: function(){
    var style = {};

    if (!SearchStore.query) {
      style.display = "none";
    }

    function close(){
      SearchStore.clear();
    }

    return m("div", { class: "card", id: "search-menu", style: style }, [
      m("div", { class: "card-header" }, [
        m("i", { class: "fa fa-search"}),
        "Search",
        m("i", { class: "fa fa-times float-right", onclick: close }),
      ]),
      m("div", { class: "card-body", style: {overflow: "auto"} }, m(SearchResult))
    ]);
  }
}
