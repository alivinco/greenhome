/**
 * Created by alivinco on 03/06/16.
 */
var chat
var text
var ws
var name = "Guest" + Math.floor(Math.random() * 1000);

function NewMsg(topic,type,cls,subcls,def,properties){
    payload = {type:type,cls:cls,subcls:subcls,def:def,props:properties}
    msg = {topic:topic,payload:JSON.stringify(payload)}
    return JSON.stringify(msg)
}

function CmdBinarySwitch(topic,value){
    msg = NewMsg(topic,"cmd","binary","switch",{value:value})
    ws.send(msg)
}


// ratchet.js push handler which is called whenever new page is loaded by push.js
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

function wsMessageHandler(msg){
    jobj = JSON.parse(msg.data)

    topic = jobj.topic
    payload = JSON.parse(jobj.payload)

    console.log(topic)
    console.dir(payload)
    $('[topic="'+topic+'"]').html(payload.def.value)
    //var line =  now() + " " + jobj.cls+"."+jobj.subcls+"="+jobj.def.value + "\n";
    //chat.innerText += line;
}

$(function() {
     // Only needed if you want to fire a callback
     window.addEventListener('push', pushHandler);
     console.dir("Connecting")
     var url = "ws://" + window.location.host + "/greenhome/ws?domain="+domain;
     ws = new WebSocket(url);
     ws.onmessage = wsMessageHandler
     if (window.location.href.includes("logs"))
     {
        initChat()
      }
});
