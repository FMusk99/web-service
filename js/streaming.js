var player = videojs('my-player');
var options = {};
var sourcce = "https://rr2---sn-jhjup-i5ol.googlevideo.com/videoplayback?expire=1640078250&ei=SkfBYfz8MpqqgQPs347YDA&ip=103.89.87.230&id=f2TT5oQhpec.1&itag=244&aitags=133%2C134%2C135%2C136%2C137%2C160%2C242%2C243%2C244%2C247%2C248%2C278&source=yt_live_broadcast&requiressl=yes&mh=tw&mm=44%2C26&mn=sn-jhjup-i5ol%2Csn-30a7yne7&ms=lva%2Conr&mv=u&mvi=2&pl=25&vprv=1&live=1&hang=1&noclen=1&mime=video%2Fwebm&ns=Nx9aFFjG8NOOHYW0JadbM-sG&gir=yes&mt=1640055778&fvip=6&keepalive=yes&fexp=24001373%2C24007246&c=WEB&n=zVhwraCxKjq8AQ&sparams=expire%2Cei%2Cip%2Cid%2Caitags%2Csource%2Crequiressl%2Cvprv%2Clive%2Chang%2Cnoclen%2Cmime%2Cns%2Cgir&sig=AOq0QJ8wRAIgUu2lw77Pgu5c3jLH-sb-ZWSZTEa1C5t5Awo37xPYuXICIErq9CEnu-pbB6YbMCDfVNRWMO_fQldERXYrsBwnQ0ov&lsparams=mh%2Cmm%2Cmn%2Cms%2Cmv%2Cmvi%2Cpl&lsig=AG3C_xAwRgIhAJo-15OACt85_csYtyTpVKYRM3a-mcwdIqOho6QyDH4NAiEAv-ZxWk_df_ckA-EZQLmUwZj4N31jo2ADhV6U8IiJmcU%3D&alr=yes&cpn=V_0eHj7s4NsSG-U9&cver=2.20211220.01.00&sq=3081&rn=28&rbuf=27496"
var playerDOM =  document.getElementById("my-player_html5_api")
var sourceTag = document.createElement("SOURCE")
sourceTag.setAttribute("src",sourcce);
sourceTag.setAttribute("type", "video/webm");
{/* <source src="" type="video/mp4"></source>
<source src="" type="video/webm"></source>
<source src="" type="video/ogg"></source> */}
playerDOM.appendChild(sourceTag);
var player = videojs('my-player', options, function onPlayerReady() {
  videojs.log('Your player is ready!');

  // In this context, `this` is the player that was created by Video.js.
  this.play();

  // How about an event listener?
  this.on('ended', function() {
    videojs.log('Awww...over so soon?!');
  });
});