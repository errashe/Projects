Template.layout.events({
	"submit form": function(e, tmpl) {
		e.preventDefault();

		var form = e.currentTarget;

		var login = tmpl.find("input[name=login]").value;
		var password = tmpl.find("input[name=password]").value;

		if(login == "" || password == "") {
			form.reset();
			alert("Форма не заполнена");
			return;
		}

		Meteor.loginWithPassword(login, password, function(error) {
			if(error && _.contains([403, 400], error.error)) {
				alert("Пользователь не найден!");
			}
		});

		form.reset();
	},
	"click #registration": function(e) {
		e.preventDefault();
		Router.go("/registration");
	}
});