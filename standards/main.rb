require 'csv'
require 'mechanize'

a = Mechanize.new

(1..12).each do |i|
	a.get("http://docs.cntd.ru/search/gostlastyear/year/2016/month/#{i}")

	CSV.open("месяц#{i}.csv", "wb") do |csv|
		(0..a.page.search(".number_doc>strong").text.to_i).step(20) do |j|
			a.get("http://docs.cntd.ru/search/gostlastyear/year/2016/month/#{i}/offset/#{j}")
			a.page.search(".content>ul>li").each do |line|
				csv << [line.search("a").text.strip.gsub(/\n\s+/, ""), line.search("span").text.strip.gsub(/\n\s+/, "")]
			end
		end
	end

	# CSV.open("месяц#{i}.csv", "wb") do |csv|
	# 	a.get("http://docs.cntd.ru/search/gostlastyear/year/2016/month/#{i}/offset/0")

	# 	a.page.search(".content>ul>li").each do |line|
	# 		csv << [line.search("a").text.strip, line.search("span").text.strip]
	# 	end
	# end
end
