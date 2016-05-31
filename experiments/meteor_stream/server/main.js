Meteor.startup(() => {
	Streamer.allowRead('all');
	Streamer.allowWrite('all');
});
