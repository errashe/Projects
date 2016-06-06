Template.parse.rendered = function() {
	Meteor.subscribe("matches");
}

Template.parse.helpers({
	"matches": function() {
		return Matches.find({}, {sort: { time: -1 }});
	},
	"color": function(stat) {
		if(stat == "Победа") {
			return "green";
		} else {
			return "red";
		}
	}
});
