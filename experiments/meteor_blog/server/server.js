import { Meteor } from 'meteor/meteor';

Meteor.startup(() => {

	Meteor.publish('articles', function () {
		return Articles.find();
	});
});

Meteor.methods({
	'addArticle': function(title, text) {
		Articles.insert({title: title, text: text});
	}
});