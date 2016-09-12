$(function() {
	$("#registration-form").on("submit", function(e) {
		e.preventDefault();
		var form = e.currentTarget;
		var pass = form.querySelector("#password-input").value;
		var conf = form.querySelector("#password-input-confirmation").value;

		if(pass == conf) {
			form.submit();
		} else {
			form.reset();
			alert("Пароли не совпадают");
		}
	});
});