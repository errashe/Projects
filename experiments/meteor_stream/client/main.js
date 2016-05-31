var keyState = {};
var users = {};
var name, canvas, ctx;

sendMessage = function(who, message) {
	Streamer.emit('message', who, message);
};

Streamer.on('message', function(who, message) {
	users[who] = message;
});

debug = function() {
	return users;
}

function gameLoop() {
	ctx.clearRect(0, 0, 500, 500);
	if (keyState[37]) {
		users[name].x -= 5
		sendMessage(name, {x: users[name].x-5, y: users[name].y});
	}
	if (keyState[39]) {
		users[name].x += 5
		sendMessage(name, {x: users[name].x+5, y: users[name].y});
	}

	for(var id in users) {
		var user = users[id];
		ctx.fillRect(user.x, user.y, 10, 10);
	}
}


window.addEventListener('keydown',function(e){
	keyState[e.keyCode || e.which] = true;
},true);
window.addEventListener('keyup',function(e){
	keyState[e.keyCode || e.which] = false;
},true);

Template.hello.rendered = function() {
	name = Date.now();
	users[name] = {x: 0, y: 0};
	canvas = document.querySelector("#game");
	ctx = canvas.getContext("2d");

	// ctx.fillRect(10, 10, 10, 10);
	setInterval(gameLoop, 1000/60);
}