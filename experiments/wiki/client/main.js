Template.main.events({
	"keydown #wikipage": function(e) {
		if(e.keyCode == 13) {
			Meteor.call("makePage", e.target.value);
		}
	}
});