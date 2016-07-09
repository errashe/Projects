class Match < ApplicationRecord
	belongs_to :gamer, foreign_key: "puid", primary_key: "uid"
	validates :uhash, presence: true, uniqueness: true
	after_create_commit { BroadcastChangesJob.perform_later }

	def self.parse_dotabuff
		players = Gamer.all
		work(players)
		Match.count
	end

	private
	def self.work(list)
		a = Mechanize.new

		list.each do |p|
			a.get("http://www.dotabuff.com/players/#{p.uid}")
			rows = a.page.search("div.r-row[data-link-to*='/matches/']")
			rows.each do |row|
				info = row.search("div.r-body>a")
				time = row.search("time").attr("datetime").text

				hero = info[0].text
				stats = info[1].text
				uid = row.attr("data-link-to").split("/")[2]

				Match.create(
					hero: hero,
					stats: (stats =~ /Lost/).nil?,
					uid: uid,
					puid: p.uid,
					uhash: "#{p.uid}:#{uid}",
					match_time: time
					)
			end
		end
	end
end
