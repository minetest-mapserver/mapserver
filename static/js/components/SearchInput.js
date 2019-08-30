
const state = {
  query: ""
}

function doSearch(){
  m.route.set(`/search/${state.query}`);
}

export default {
  view: function(){

    function handleInput(e){
      state.query = e.target.value;
    }

    function handleKeyDown(e){
      if (e.keyCode == 13){
        doSearch();
      }
    }

    return m("div", { class: "input-group mb-3" }, [
      m("input[type=text]", {
        placeholder: "Search",
        class: "form-control",
        oninput: handleInput,
        onkeydown: handleKeyDown,
        value: state.query
      }),
      m("div", { class: "input-group-append", onclick: doSearch }, [
        m("span", { class: "input-group-text" }, [
          m("i", { class: "fa fa-search"})
        ])
      ])
    ]);
  }
};
