import { Template } from 'meteor/templating';
import { ReactiveVar } from 'meteor/reactive-var';

Meteor.subscribe('articles');

Template.main.events({
	'submit #article': function(event) {
		event.preventDefault();
		form = event.target;
		title = form.querySelector("#title");
		text = form.querySelector("#text");

		Meteor.call("addArticle", title.value, text.value);
		title.value = "";
		text.value = "";
	}
});

Template.test.helpers({
	articles: function() {
		return Articles.find();
	}
});

Router.configure({
	layoutTemplate: 'AppL'
});

Router.route('/', function () {
	this.render('main');
});

Router.route('/test', function() {
	this.render('test');
});