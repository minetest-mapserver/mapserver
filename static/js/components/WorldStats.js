import html from "./html.js";

export default function(info){
  const is_night = info.time < 5500 || info.time > 19000;

  function getMinute(){
    var min = Math.floor((info.time % 1000) / 1000 * 60);
    return min >= 10 ? min : "0" + min;
  }

  var color = "green";
  if (info.max_lag > 0.8)
    color = "orange";
  else if (info.max_lag > 1.2)
    color = "red";

  return html`<div>
    <span class="fa fa-users"/> ${info.players ? info.players.length : "0"}
    <span class="fa fa-wifi" style="color: ${color}"/>
    ${parseInt(info.max_lag*1000)} ms
    <span class="fa fa-clock"/>
    <span class="fa fa-${is_night ? "moon" : "sun"}" style="color: ${is_night ? "blue" : "orange"};"/>
    ${Math.floor(info.time/1000)}:${getMinute()}
  </div>`;
}
