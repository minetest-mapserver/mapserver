
export function hashCompat(){

  if (window.location.hash) {

    const parts = window.location.hash.replace("#", "").split("/");
    if (parts.length == 0){
      //invalid
      return;
    }

    if (parts[0] == "map"){
      //new link
      return;
    }

    if (isNaN(+parts[0])){
      //NaN
      return;
    }

    if (parts.length == 3){
      // #1799.5/399/10
      window.location.hash = `#!/map/0/${parts[0]}/${parts[1]}/${parts[2]}`;
    }

    if (parts.length == 4) {
      // #0/-1799.5/399/10
      // #0/5405.875/11148/12
      window.location.hash = `#!/map/${parts[0]}/${parts[3]}/${parts[1]}/${parts[2]}`;
    }
  }
}
