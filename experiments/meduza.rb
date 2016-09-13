#!/usr/bin/env ruby

require 'unirest'
require 'pp'

res = Unirest.get("https://meduza.io/api/v3/search?chrono=news&page=0&per_page=10&locale=ru")

res.body["documents"].each do |article|
	p article[1]["title"]
end