/**
 * Created by alivinco on 03/06/16.
 */
var chat
var text
var ws
var name = "Guest" + Math.floor(Math.random() * 1000);

function pushHandler(event){
    url = event.detail.state.url

    if(url.includes("logs")){
        initChat()
    }
    console.dir(url)
}

var now = function () {
        var iso = new Date().toISOString();
        return iso.split("T")[1].split(".")[0];
      };

function initChat(){
         chat = document.getElementById("chat");
         text = document.getElementById("text");
         text.onkeydown = function (e) {
             //console.log("keydown")
                if (e.keyCode === 13 && text.value !== "") {
                  ws.send("<" + name + "> " + text.value);
                  text.value = "";
                }
              };
}

$(function() {
     // Only needed if you want to fire a callback
     window.addEventListener('push', pushHandler);
     console.dir("Connecting")
     var url = "ws://" + window.location.host + "/greenhome/ws";
     ws = new WebSocket(url);
     ws.onmessage = function (msg) {
                jobj = JSON.parse(msg.data)

                var line =  now() + " " + jobj.cls+"."+jobj.subcls+"="+jobj.def.value + "\n";
                chat.innerText += line;
              };
     if (window.location.href.includes("logs"))
     {
        initChat()
      }
});
