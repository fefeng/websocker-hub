var sock = null;
var wsuri = "ws://127.0.0.1:1234/socket"
window.onload = function () {
    console.log("window on load");
    sock = new WebSocket(wsuri);
    sock.onopen = function () {
        console.log("connected to " + wsuri);
    }
    sock.onclose = function (e) {
        console.log("connection close ", e.code)
    }
    sock.onmessage = function (e) {
        // document.getElementById("message1").value = e.data;
        document.getElementById("message1").value = e.data
        console.log("message received : " + e.data)
    }
}

function send() {
    var msg = document.getElementById("message").value;
    sock.send(msg);
}