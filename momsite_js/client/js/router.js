Router.configure({
	layoutTemplate: 'layout',
	loadingTemplate: 'loading',
	notFoundTemplate: 'loading'
});

Router.route('/', function() {
	this.render("main");
});

Router.route('/logout', function() {
	Meteor.logout();
	this.redirect("/");
});

Router.route('/registration', function() {
	if(Meteor.user()) this.redirect("/");
	this.render("registration");
});