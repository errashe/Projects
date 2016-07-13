#!/usr/bin/env ruby

require 'active_support'
require 'active_support/core_ext'
require 'mechanize'
require 'thread'
require 'thwait'
require 'json'

Game = Struct.new(:hero, :stats, :link)

def fill_list
	@threads = []
	@return = []
	@list = []

	@list << "http://www.dotabuff.com/players/92413647"
	@list << "http://www.dotabuff.com/players/261384156"
	@list << "http://www.dotabuff.com/players/138747075"
	@list << "http://www.dotabuff.com/players/130181018"
	@list << "http://www.dotabuff.com/players/87094975"
	@list << "http://www.dotabuff.com/players/98900816"
	@list << "http://www.dotabuff.com/players/23509620"
	@list << "http://www.dotabuff.com/players/149733512"
end

def parse
	while !@list.nil?
		a = Mechanize.new
		a.get(@list.pop)
		rows = a.page.search("div.r-row[data-link-to*='/matches/']")
		rows.each do |row|
			info = row.search("div.r-body>a")

			hero = info[0].text
			stats = info[1].text
			link = info[0].attr("href")

			game = Game.new(hero, stats, link)
			@return << game
		end
	end
	Thread::exit()
end

while true
	fill_list()
	(0..5).each do |i|
		@threads << Thread.new{ parse }
	end
	ThreadsWait.all_waits(*@threads)
	p @return.count
	sleep 10.seconds
end