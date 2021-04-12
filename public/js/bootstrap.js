(function(){
  var s = document.createElement("script");

  if (location.host === "127.0.0.1:8080") {
    //dev
    s.setAttribute("src", "js/main.js");
    s.setAttribute("type", "module");

  } else {
    //prod
    s.setAttribute("src", "js/bundle.js");

  }

  document.body.appendChild(s);
})();
