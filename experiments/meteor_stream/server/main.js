import { Meteor } from "meteor/meteor"

import '/libs/main.js'

Meteor.startup(() => {
	Streamer.allowRead('all');
	Streamer.allowWrite('all');
});
