/* exported SearchInput */
/* globals SearchService: true */
/* globals SearchStore: true */

var SearchInput = {
  view: function(){
    function handleInput(e){
      SearchService.search(e.target.value);
    }

    return m("div", { class: "input-group mb-3" }, [
      m("div", { class: "input-group-prepend" }, [
        m("span", { class: "input-group-text" }, [
          m("i", { class: "fa fa-search"})
        ])
      ]),
      m("input[type=text]", {
        placeholder: "Search",
        class: "form-control",
        oninput: handleInput,
        value: SearchStore.query
      })
    ]);
  }
};
