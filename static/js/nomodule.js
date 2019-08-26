
m.mount(document.getElementById("app"), {
  view: function(){
    return m("div", "I'm sorry, your browser is just too old ;)");
  }
});
