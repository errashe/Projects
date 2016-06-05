Template.parse.rendered = function() {
	Meteor.subscribe("matches");
}

Template.parse.helpers({
	"matches": function() {
		return Matches.find({});
	},
	"color": function(stat) {
		if(stat == "Победа") {
			return "green";
		} else {
			return "red";
		}
	}
});
