import cheerio from "cheerio"
import request from "request"

function parse() {
	var options = {
		url: 'http://ru.dotabuff.com/players/92413647',
		headers: {
			"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.79 Safari/537.36"
		}
	};

	request(options, Meteor.bindEnvironment(function(err, resp, body){
		$ = cheerio.load(body);
		var rows = $("div.r-table div.r-row[data-link-to*='/matches/']"); // TODO: create a normal css selector

		rows.each(function(i, e) {
			var match = $(e).attr("data-link-to").split('/')[2];
			var hero = $(e).find("a:nth-child(2)").text();
			var stat = $(e).find("a:nth-child(1)").text();
			var time = $(e).find("time").attr("datetime");

			Matches.upsert({match: match}, { $setOnInsert: {
				match: match,
				hero: hero,
				stat: stat,
				time: time
			}});
		});
	}));
}

Meteor.methods({
	"parse": parse
});

Meteor.startup(() => {
	parse();
	setInterval(Meteor.bindEnvironment(parse), 1000*60);

	Meteor.publish("matches", function() {
		return Matches.find({});
	});
});
