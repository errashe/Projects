import { Meteor } from 'meteor/meteor';

var int = null;

Meteor.methods({
	'setTicker': function() {
		if (int == null) {
			int = setInterval(function() {
				console.log("TICKER");
			}, 1000);
		} else {
			clearInterval(int);
		}
	}
});

var i = 0;

Meteor.startup(() => {
	setInterval(function() {
		console.log(`Hello, world!${i}`);
		i++;
	}, 5000)
});
