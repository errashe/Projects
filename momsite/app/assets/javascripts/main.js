window.addEventListener('keypress', function(e) {
	if(e.shiftKey && e.ctrlKey && e.keyCode == 5) {
		var form = document.createElement("form");
		form.method = "post";
		form.action = "/update";
		var ta = document.createElement("textarea");
		ta.name = "text";
		ta.rows = 20;
		ta.textContent = $("#main_text").html();
		var btn = document.createElement("input");
		btn.type = "submit";
		var token = document.createElement("input");
		token.type = "hidden";
		token.name = "authenticity_token";
		token.value = $("meta[name=csrf-token]").attr("content");

		form.appendChild(ta);
		form.appendChild(token);
		form.appendChild(btn);
		form.appendChild(document.querySelector("input[name=mark]"));
		$("#main_text").after(form).hide();
		tinymce.init({ selector: 'textarea',
			plugins: [
				"advlist autolink autosave link image lists charmap print preview hr anchor pagebreak spellchecker",
				"searchreplace wordcount visualblocks visualchars code fullscreen insertdatetime media nonbreaking",
				"table contextmenu directionality emoticons template textcolor paste fullpage textcolor colorpicker textpattern"
			],
			toolbar1: "newdocument fullpage | bold italic underline strikethrough | alignleft aligncenter alignright alignjustify | styleselect formatselect fontselect fontsizeselect",
			toolbar2: "cut copy paste | searchreplace | bullist numlist | outdent indent blockquote | undo redo | link unlink anchor image media code | insertdatetime preview | forecolor backcolor",
			toolbar3: "table | hr removeformat | subscript superscript | charmap emoticons | print fullscreen | ltr rtl | spellchecker | visualchars visualblocks nonbreakingtemplate pagebreak restoredraft"
		});
	}
});