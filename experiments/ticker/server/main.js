import { Meteor } from 'meteor/meteor';

var i = 0;

Meteor.startup(() => {
	setInterval(function() {
		console.log(`Hello, world!${i}`);
		i++;
	}, 500)
});
