package main

import (
	"fmt"
	"log"
	"net/http"
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
		}
	</script>
	<form method="post" action="/graph">
		<div id="params"></div>
		<input type="submit">
		<button onclick="add()">Добавить параметр поиска</button>
	</form>
`

var result string = `
	<script src="https://api-maps.yandex.ru/2.1/?lang=en_US" type="text/javascript"></script>
	<script type="text/javascript">
		function str2xy(str) {
			var arrContent = str.split(",");
			var x = parseFloat(arrContent[0]);
			var y = parseFloat(arrContent[1]);
			return [x,y]
		}
	
		ymaps.ready(init);
		var myMap;

		function init(){     
			myMap = new ymaps.Map("map", {
				center: [55.76, 37.64], 
				zoom: 7
			});

			var points = document.querySelectorAll(".point");
			var poly = [];
			for(var i=0; i<points.length; i++) {
				var temp = str2xy(points[i].textContent);
				myPlacemark = new ymaps.Placemark(temp, {
					iconContent: i,
				});
				myMap.geoObjects.add(myPlacemark);
				poly.push(temp);
			}
			var polyline = new ymaps.Polyline(poly);

			myMap.geoObjects.add(polyline);
			myMap.setBounds(polyline.geometry.getBounds());

			// ymaps.route(poly, {
			// 	multiRoute: true
			// }).done(function (route) {
			// 	route.getWayPoints().options.set("visible", false)
			// 	myMap.geoObjects.add(route);
			// 	myMap.setBounds(route.getBounds());
			// }, function (err) {
			// 	throw err;
			// }, this);
		}


	</script>
	<div id="map" style="width: 1000px; height: 1000px; float: left;"></div>
	<div style="margin-left: 1020px;" id="points">%s</div>
`

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, page, form)
}

func ghandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	res := StartSearch(r.Form["path"])
	var f string
	if len(res) == 0 {
		fmt.Fprintf(w, page, "Не достаточно аргументов или не найдены объекты по критериям")
		return
	}
	for _, item := range res {
		f += fmt.Sprintf("<span class='point'>%s</span><br />", item)
	}

	fmt.Fprintf(w, page, fmt.Sprintf(result, f))
}

func StartServer() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/graph", ghandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
