import r from "request"
import c from "cheerio"

Meteor.methods({
	"makePage": function(url) {
		console.log(url);
		var options = {
			url: encodeURI(url),
			headers: {
				"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.79 Safari/537.36"
			}
		};

		r.get(options, Meteor.bindEnvironment(function(e, res) {
			if(e == null) {
				$ = c.load(res.body);
				var text = $("div#toc").text();
				Temp.insert({text: text})
			} else {
				console.log(e);
			}
		}));
	}
});