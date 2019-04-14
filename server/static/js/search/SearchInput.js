
var SearchInput = {
  view: function(){
    function handleInput(e){
      SearchStore.search(e.target.value);
    }

    return m("div", { class: "input-group mb-3" }, [
      m("div", { class: "input-group-prepend" }, [
        m("span", { class: "input-group-text" }, [
          m("i", { class: "fa fa-search"})
        ])
      ]),
      m("input[type=text]", { placeholder: "Search", class: "form-control", oninput: handleInput, value: SearchStore.query })
    ]);
  }
}
