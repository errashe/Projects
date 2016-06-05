Template.parse.rendered = function() {
	Meteor.subscribe("matches");
}

Template.parse.helpers({
	"matches": function() {
		return Matches.find({}, {sort: { time: -1 }, limit: 10});
	},
	"color": function(stat) {
		if(stat == "Победа") {
			return "green";
		} else {
			return "red";
		}
	}
});

Template.parse.events({
	'click button'(event, instance) {
		Meteor.call("parse");
	},
});
