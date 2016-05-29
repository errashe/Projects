import { Meteor } from 'meteor/meteor';

Meteor.methods({
	'addPlayer': function(name) {
		Players.insert({name: name, x:0, y:0});
	},
	'move': function(name, pos) {
		Players.update({name: name}, {$inc: {x: pos}})
	}
});

Meteor.startup(() => {
	Meteor.publish("players", function(name) {
		return Players.find({name: name});
	});
});
