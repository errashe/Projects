import r from 'request'
var token = "18125f3dea048c0dfc3b75b906691ee20a603e56ba0cdaaefcdca43733602d1e96b07c7a199e3077a175f";

Meteor.startup(() => {
	r.get("http://yandex.ru/", function(e, r) {
		console.log(r.body);
	});
});
