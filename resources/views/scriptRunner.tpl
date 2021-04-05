=====
Pagetitle: Site Details
BodyClasses: layout-details
=====
<h1 class="pt-3">Running Script</h1>

<p>Messages:</p>
<div id="messages"></div>

<script>
var ws; // our websocket
var proto = window.parent.location.protocol === "http:" ? "ws": "wss"; // dynamic protocol
var messageHolder = document.getElementById("messages");

ws = new WebSocket(proto + '://' + window.location.host + '/ws');
ws.addEventListener('message', function(e) {
    var msg = JSON.parse(e.data);
    var node = document.createElement("p");
    var text = document.createTextNode(msg.message);
    node.appendChild(text);

    messageHolder.appendChild(node);
});
ws.addEventListener('close', function(e) {
    var a = document.createElement("a");
    var text = document.createTextNode("Complete");
    a.appendChild(text);
    a.href="javascript:window.history.back()";
    a.classList.add("btn");
    a.classList.add("btn-primary");

    messageHolder.insertAdjacentElement('afterend', a);
});

var ofcoServerMessage = {
    "Action": {{ .Action }},
    "Domain": {{ .Domain }}
};
ws.onopen = function (event) {
    ws.send(JSON.stringify(ofcoServerMessage));
};
</script>
