Meteor.startup(() => {
	Meteor.publishComposite('secWsubsec', {
		find: function() {
				return Section.find({});
		},
		children: [
			{
				find: function(section) {
					return Subsection.find({sectionId: section._id});
				},
			},
		]
	});
});
