var request = require("request");
var cheerio = require("cheerio");

exports.go = function(url, next) {
	var options = {
		url: url,
		headers: {
			"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.79 Safari/537.36"
		}
	};

	request(options, function(err, resp, body){
		var $ = cheerio.load(body);
		var user = $("div.header-content-title h1");
		user = user.clone().children().remove().end().text();
			var rows = $("div.r-table div.r-row[data-link-to*='/matches/']"); // TODO: create a normal css selector

			next($, rows, user);

			// rows.each(function(i, e) {
			// 	var match = $(e).attr("data-link-to").split('/')[2];
			// 	var hero = $(e).find("a:nth-child(2)").text();
			// 	var stat = $(e).find("a:nth-child(1)").text();
			// 	var time = $(e).find("time").attr("datetime");

			// 	Matches.upsert({match: match, user: user}, { $setOnInsert: {
			// 		user: user,
			// 		match: match,
			// 		hero: hero,
			// 		stat: stat,
			// 		time: time
			// 	}});
			// });
		});
}
