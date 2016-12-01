FlowRouter.route('/', {
	name: 'root',
	action(params, queryParams) {
		BlazeLayout.render('App_body', {main: 'root'});
	}
});

FlowRouter.route('/subsection/:id', {
	name: 'section',
	action(params, queryParams) {
		BlazeLayout.render('App_body', {main: 'subsection'});
	}
});

Template.root.onCreated(function() {
	Meteor.subscribe("secWsubsec");
});

Template.root.helpers({
	sections() {
		return Section.find();
	},
	subsections() {
		return Subsection.find({sectionId: this._id});
	}
});

Template.subsection.helpers({
	id() {
		return FlowRouter.current().params.id;
	}
});