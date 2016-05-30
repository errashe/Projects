Meteor.methods({
	'regPlayer': function(name) {
		Players.upsert({ name: name }, {
			$setOnInsert: { name: name, x: 0, y:0 }
		})
	},
	'movePlayer': function(name, pos) {
		Players.update({ name: name }, { $inc: pos });
	}
});

Meteor.startup(() => {
	Meteor.publish("players", function() {
		return Players.find();
	});
});
