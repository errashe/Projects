Debug = function() {
	Meteor.call("debug");
}

Template.parse.onCreated(function() {
	Meteor.subscribe("matches");
});

Template.parse.helpers({
	"matches": function() {
		return Matches.find({}, {sort: { time: -1 }});
	},
	"color": function(stat) {
		if(stat == "Победа") {
			return "success";
		} else {
			return "danger";
		}
	},
	"strToHours": function(str) {
		moment.updateLocale("ru");
		return moment(str).fromNow();
	}
});