require 'sinatra'

get '/' do
	erb :index
end

get '/draw' do
	draw
end

get '/:file' do
	send_file "#{params[:file]}"
end

def flush_graph(graph, points, last=false)
	File.open("#{graph}.txt", "a") do |f|
		points.each do |point|
			f.puts "#{point.join(' ')}\n"
		end
		f.puts "\n\n" if !last
	end
end

def gnuplot(commands)
	IO.popen("gnuplot", "w") { |io| io.puts commands }
end

def draw
	@filenames = ["xy"]
	@filenames.each{ |e| File.delete("#{e}.txt") if File.exists? "#{e}.txt" }

	width = 1200
	height = 800

	vars = ["wx", "dx", "wy", "dy", "a", "b"]

	wx = params["wx"].to_f || 1
	dx = params["dx"].to_f || 0
	wy = params["wy"].to_f || 1
	dy = params["dy"].to_f || 0
	a1 = params["a"].to_f || 1
	b1 = params["b"].to_f || 1

	a = a1 / [a1, b1].min
	b = b1 / [a1, b1].min

	if ((wx + wy) > 30)
		itlen = 0.005
	elsif ((wx + wy) > 12)
		itlen = 0.01
	else
		itlen = 0.02
	end

	xlast = (Math.cos(wx * 0 + dx) * (width / 2 - 10) / b + width / 2)
	ylast = height - (Math.cos(wy * 0 + dy) * (height / 2 - 10) / a + height / 2)

	draw = []

	(0..6.4).step(itlen) do |i|
		x = (Math.cos(wx * i + dx) * (width / 2 - 10) / b + width / 2)
		y = height - (Math.cos(wy * i + dy) * (height / 2 - 10) / a + height / 2);
		draw.push([x, y.round(2)])
		xlast = x
		ylast = y
	end

	flush_graph("xy", draw)

	commands = %Q(

	set terminal gif medium size 1200 800
	set output 'foobar.gif'

	stats 'xy.txt' nooutput

	# set xrange [-10:10]
	# set yrange [-100:100]
	
	plot 'xy.txt' with lines

	)

	gnuplot(commands)

end