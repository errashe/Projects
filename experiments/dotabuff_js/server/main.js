getData = (userid) => {
	Meteor.http.get(`http://www.dotabuff.com/players/${userid}/matches`, {
		headers: {"User-Agent": "Meteor: 1.4.0"}
	}, (er, data) => {
		if(er) return;
		let cheerio = Npm.require("cheerio");
		let $ = cheerio.load(data.content);

		let rows = $("article > table > tbody > tr");
		let player = Players.findOne({id: userid})

		rows.map((i, el) => {
			let match = {};
			match.hero = $(el).find("td:nth-child(2) > a").text();
			match.bracket = $(el).find("td:nth-child(2) > div").text();
			match.status = $(el).find("td:nth-child(3) > a").attr("class");
			match.id = $(el).find("td:nth-child(3) > a").attr("href").replace("/matches/", "");
			match.userid = userid;
			match.mapl = `${match.id}-${userid}`
			match.time = $(el).find("td:nth-child(3) > div > time").attr("datetime");
			match.player = player;

			Matches.insert(match, (err) => {
				if(err && err.code != 11000) {
					console.log(err.message)
				}
			});
		});
	});

	return "done";
}

getDataOfAllUsers = () => {
	Players.find().fetch().forEach(el => getData(el.id));
}

Meteor.startup(() => {
	setInterval(Meteor.bindEnvironment(getDataOfAllUsers), 60000);

	Meteor.publish("lastMatches", () => {
		return Matches.find({}, {limit: 15, sort: { time: -1, id: -1 }});
	})
});

Meteor.methods({
	"getData": getData
});
