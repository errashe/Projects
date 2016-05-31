var keyState = {};
var users = {};
var name;

sendMessage = function(message) {
	Streamer.emit('message', message);
	console.log('me: ' + message);
};

startMadness = function() {
	setInterval(function() {
		sendMessage("OMG");
	}, 1000/60);
}

Streamer.on('message', function(message) {

});

window.addEventListener('keydown',function(e){
	keyState[e.keyCode || e.which] = true;
},true);
window.addEventListener('keyup',function(e){
	keyState[e.keyCode || e.which] = false;
},true);

Template.hello.rendered = function() {
	name = Date.now();
	var canvas = document.querySelector("#game");
	var ctx = canvas.getContext("2d");

	ctx.fillRect(10, 10, 10, 10);
}