import html from "./html.js";

const state = {
  query: ""
};

function doSearch(){
  m.route.set(`/search/${state.query}`);
}

function handleInput(e){
  state.query = e.target.value;
}

function handleKeyDown(e){
  if (e.keyCode == 13){
    doSearch();
  }
}

export default {
  view: () => html`<div class="input-group mb-3">
    <input type="text"
      class="form-control"
      placeholder="Search"
      oninput=${handleInput}
      onkeydown=${handleKeyDown}
      value=${state.query}
    />
    <div class="input-group-append" onclick=${doSearch}>
      <span class="input-group-text">
        <i class="fa fa-search"/>
      </span>
    </div>
  </div>`
};
