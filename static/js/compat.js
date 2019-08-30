
export function hashCompat(){

  if (window.location.hash) {
    let match = window.location.hash.match(/^#\/(\d*)\/(\d*)\/(\d*)$/m);
    if (match) {
      window.location.hash = `#!/map/0/${match[1]}/${match[2]}/${match[3]}`
    }

    match = window.location.hash.match(/^#\/(\d*)\/(\d*)\/(\d*)\/(\d*)$/m);
    if (match) {
      window.location.hash = `#!/map/${match[1]}/${match[2]}/${match[3]}/${match[4]}`
    }
  }
}
