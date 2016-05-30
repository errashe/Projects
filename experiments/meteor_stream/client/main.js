import { Meteor } from "meteor/meteor"

import "/libs/main.js"

sendMessage = function(message) {
	Streamer.emit('message', message);
	console.log('me: ' + message);
};

Streamer.on('message', function(message) {
	console.log('user: ' + message);
});