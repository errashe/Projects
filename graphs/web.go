package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var page string = `
	<html>
		<head><title>Hello, world!</title></head>
		<body>%s</body>
	</html>
`

var form string = `
	<script type="text/javascript">
		function add() {
			event.preventDefault();
			
			var div = document.createElement("div");
			div.classList.add("text");
			var input = document.createElement("input");
			input.name = "path";
			input.type = "text";
			div.appendChild(input);
			var del = document.createElement("button");
			del.textContent = "Удалить точку";
			del.onclick = function() {
				event.preventDefault();
				this.parentElement.remove();
			}
			div.appendChild(del);
			var where = document.querySelector("#params");
			where.appendChild(div);
			console.log("Hello, world!")
		}

		function debug() {
			console.log(countOfInputs());
		}
	</script>
	<form method="post" action="/graph">
		<div id="params"></div>
		<input type="submit">
		<button onclick="add()">Добавить параметр поиска</button>
	</form>
`

var result string = `
	<a href="%s">Карта</a><br />
	%s
`

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, page, form)
}

func ghandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	res := StartSearch(r.Form["path"])
	link := fmt.Sprintf("https://yandex.ru/maps/53/kurgan/?rtext=%s", strings.Join(res, "~"))
	var f string
	if len(res) == 0 {
		fmt.Fprintf(w, page, "Не достаточно аргументов или не найдены объекты по критериям")
		return
	}
	for _, item := range res {
		f += fmt.Sprintf("%s<br />", item)
	}

	fmt.Fprintf(w, page, fmt.Sprintf(result, link, f))
}

func StartServer() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/graph", ghandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
