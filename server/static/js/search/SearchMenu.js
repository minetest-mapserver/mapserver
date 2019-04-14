

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
      m("div", { class: "card-body" })
    ]);
  }
}
