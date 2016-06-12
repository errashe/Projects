import Parser from '../modules/parser.js';

var list = [
"http://ru.dotabuff.com/players/92413647"
];

function parseAll() {
	for(var i in list) {
		Parser.go(list[i], Meteor.bindEnvironment(function($, rows, user) {
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
}

function publish() {
	Meteor.publish("matches", function() {
		return Matches.find({}, {limit: 15, sort: { time: -1 }});
	});
}

Meteor.startup(() => {
	Meteor.setInterval(parseAll, 10*1000);
	publish();
});

Meteor.methods({
	"debug": function() {
		Matches.insert({user: "test", time: "2016-06-11T23:19:08+00:00", hero: "test", match: "100500", stat: "Победа"})
		setTimeout(Meteor.bindEnvironment(function() {
			Matches.remove({user: "test"});
		}), 5000);
	}
});
