import "/libs/main.js"

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
	console.log('user: ' + message);
});