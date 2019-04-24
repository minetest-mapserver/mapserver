/* exported SearchInput */
/* globals SearchService: true */
/* globals SearchStore: true */

var SearchInput = {
  view: function(){
    function handleInput(e){
      SearchStore.query = e.target.value;
    }

    function handleKeyDown(e){
      if (e.keyCode == 13){
        SearchService.search();
      }
    }

    function handleDoSearch(){
      SearchService.search();
    }

    return m("div", { class: "input-group mb-3" }, [
      m("input[type=text]", {
        placeholder: "Search",
        class: "form-control",
        oninput: handleInput,
        onkeydown: handleKeyDown,
        value: SearchStore.query
      }),
      m("div", { class: "input-group-append", onclick: handleDoSearch }, [
        m("span", { class: "input-group-text" }, [
          m("i", { class: "fa fa-search"})
        ])
      ])
    ]);
  }
};
