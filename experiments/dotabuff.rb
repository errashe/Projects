require 'mechanize'

Game = Struct.new(:hero, :stats, :link)

def parse(list)
	a = Mechanize.new
	ret = []

	list.each do |url|
		a.get(url)
		rows = a.page.search("div.r-row[data-link-to*='/matches/']")
		rows.each do |row|
			info = row.search("div.r-body>a")

			hero = info[0].text
			stats = info[1].text
			link = info[0].attr("href")

			ret << Game.new(hero, stats, link)
		end
	end

	return ret
end

list = []
list << "http://www.dotabuff.com/players/92413647"
list << "http://www.dotabuff.com/players/261384156"
list << "http://www.dotabuff.com/players/138747075"
list << "http://www.dotabuff.com/players/130181018"
list << "http://www.dotabuff.com/players/87094975"
list << "http://www.dotabuff.com/players/98900816"
list << "http://www.dotabuff.com/players/23509620"
list << "http://www.dotabuff.com/players/149733512"

p parse(list)
