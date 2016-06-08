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
		var d = new Date();
		d.setTime(Date.parse(str));
		var res = Date.now() - d.getTime();
		res = (res/1000).toFixed(0);
		res = (res/(60*60)).toFixed(2);
		res = Math.round(res);
		return res == 0 ? "Только что" : `${res} часа(ов) назад`;
	}
});