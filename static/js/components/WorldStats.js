import html from "./html.js";

export default function(info){

  var timeIcon = html`<span class="fa fa-sun style="color: orange;"/>`;

  if (info.time < 5500 || info.time > 19000) //0 - 24'000
    timeIcon = html`<span class="fa fa-moon style="color: blue;"/>`;

  function getHour(){
    return Math.floor(info.time/1000);
  }

  function getMinute(){
    var min = Math.floor((info.time % 1000) / 1000 * 60);
    return min >= 10 ? min : "0" + min;
  }

  function getLag(){
    var color = "green";
    if (info.max_lag > 0.8)
      color = "orange";
    else if (info.max_lag > 1.2)
      color = "red";

    return html`<span class="fa fa-wifi" style="color: ${color}"/> ${parseInt(info.max_lag*1000)} ms`;
  }

  function getPlayers(){
    return html`<span class="fa fa-users"/> ${info.players ? info.player.length : "0"}`;
  }

  return html`<div>
    ${getPlayers()}
    ${getLag()}
    <span class="fa fa-clock">${timeIcon}</span>
    ${getHour(), ":", getMinute()}
  </div>`;
}
