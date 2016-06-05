import cheerio from "cheerio"
import request from "request"

function parse(url) {
	var options = {
		url: url,
		headers: {
			"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.79 Safari/537.36"
		}
	};

	request(options, Meteor.bindEnvironment(function(err, resp, body){
		console.log(err);
		$ = cheerio.load(body);
		var user = $("div.header-content-title h1");
		user = user.clone().children().remove().end().text();
		var rows = $("div.r-table div.r-row[data-link-to*='/matches/']"); // TODO: create a normal css selector

		rows.each(function(i, e) {
			var match = $(e).attr("data-link-to").split('/')[2];
			var hero = $(e).find("a:nth-child(2)").text();
			var stat = $(e).find("a:nth-child(1)").text();
			var time = $(e).find("time").attr("datetime");

			Matches.upsert({match: match, user: user}, { $setOnInsert: {
				user: user,
				match: match,
				hero: hero,
				stat: stat,
				time: time
			}});
		});
	}));
}

var list = [
"http://ru.dotabuff.com/players/87094975",
"http://ru.dotabuff.com/players/92413647",
"http://ru.dotabuff.com/players/36981197",
"http://ru.dotabuff.com/players/23509620",
"http://ru.dotabuff.com/players/107020823",
"http://ru.dotabuff.com/players/130181018"
];

function parseAll() {
	for(var i in list) {
		Meteor.bindEnvironment(parse(list[i]));
	}
}

Meteor.startup(() => {
	setInterval(Meteor.bindEnvironment(parseAll), 10*1000);

	Meteor.publish("matches", function() {
		return Matches.find({}, {sort: { time: -1 }, limit: 20});
	});
});
