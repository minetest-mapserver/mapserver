
var SearchInput = {
  view: function(){
    function handleInput(e){
      SearchStore.search = e.target.value;
    }

    return m("input[type=text]", { placeholder: "Search", class: "form-control", oninput: handleInput });
  }
}
