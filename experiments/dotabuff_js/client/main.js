Template.main.events({
	"click button": () => {
		Meteor.call("getData", 92413647, (err, data) => console.log(data));
		Meteor.call("getData", 261384156, (err, data) => console.log(data));
	}
});

Template.main.onCreated(() => {
	Meteor.subscribe("lastMatches");
	moment.locale("ru");
})

Template.main.helpers({
	"matches": () => {
		let players = Players.find().fetch();
		return Matches.find({}, {sort: { time: -1, id: -1 }});
	},
	"color": (id) => {
		let color = id%1000000;
		return color < 100000 ? color*10 : color;
	},
	"fromNow": (time) => {
		return moment(time).fromNow();
	},
	"cl": (e) => console.log(e)
});