

var SearchMenu = {
  view: function(){
    var style = {};

    if (!SearchStore.search) {
      style.display = "none";
    }

    return m("div", { class: "card", id: "search-menu", style: style }, [
      m("div", { class: "card-header" }, "Search"),
      m("div", { class: "card-body" })
    ]);
  }
}
