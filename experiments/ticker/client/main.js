Template.hello.events({
	'click button'(event, instance) {
		Meteor.call("setTicker");
	},
});
