#!/usr/bin/env ruby

require 'active_support'
require 'active_support/core_ext'
require 'action_view'
require 'action_view/helpers'
require 'mechanize'
require 'thread'
require 'thwait'
require 'json'

include ActionView::Helpers::DateHelper

Game = Struct.new(:user, :skill, :hero, :stats, :link, :time)

def constructor
	@threads = []
	@return = []
	@list = []

	@list << "http://www.dotabuff.com/players/92413647/matches"
	@list << "http://www.dotabuff.com/players/261384156/matches"
	@list << "http://www.dotabuff.com/players/138747075/matches"
	@list << "http://www.dotabuff.com/players/130181018/matches"
	@list << "http://www.dotabuff.com/players/87094975/matches"
	@list << "http://www.dotabuff.com/players/98900816/matches"
	@list << "http://www.dotabuff.com/players/23509620/matches"
	@list << "http://www.dotabuff.com/players/149733512/matches"
	@list << "http://www.dotabuff.com/players/241084305/matches"
	@list << "http://www.dotabuff.com/players/92033022/matches"
end

def destructor
	ThreadsWait.all_waits(*@threads)

	@new = @return.sort_by{|e| [e["time"], e["user"]]}.each{|e| e["time"] = time_ago_in_words Time.parse(e["time"])}.reverse
	@new = @new.map{|e| e.to_h.to_json }
	File.write("test.txt", @new.join("\n"))
	p @new.count
end

def parse
	while !@list.nil?
		a = Mechanize.new
		a.get(@list.pop)
		user = a.page.title.split(" - ").first
		rows = a.page.search("tbody>tr")
		rows.each do |row|
			td = row.search("td")

			hero = td[1].search("a").text
			skill = td[1].search("div").text
			link = td[1].search("a").attr("href").value
			stats = td[2].search("a").attr("class").value
			time = td[2].search("time").attr("datetime").text

			game = Game.new(user, skill, hero, stats, link, time)
			# @file.write("#{game.to_h.to_json}\n")
			@return << game
		end
	end
	Thread::exit()
end

while true
	constructor()
	(0..5).each do |i|
		@threads << Thread.new{ parse }
	end
	destructor()

	sleep 10.seconds
end