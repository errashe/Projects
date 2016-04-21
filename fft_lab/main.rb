require 'dentaku'
require './dft.rb'

def flush_graph(graph, points)
	File.open("#{graph}.txt", "a") do |f|
		points.each do |point|
			f.puts "#{point.join(' ')}\n"
		end
		f.puts "\n\n"
	end
end

def gnuplot(commands)
	IO.popen("gnuplot", "w") { |io| io.puts commands }
end

@filenames = ["xy", "dft", "de-dft"]
@filenames.each{ |e| File.delete("#{e}.txt") if File.exists? "#{e}.txt" }
calculator = Dentaku::Calculator.new

range = (-10..10).step(0.1)
x = range.map { |e| e }
y = x.map { |e| calculator.evaluate(ARGV[0], x: e) }
z = DFT.dft(y)
c = DFT.dedft(z)

flush_graph(@filenames[0], range.each_with_index.map{ |e, i| [x[i], y[i].round(2)] })
flush_graph(@filenames[1], range.each_with_index.map{ |e, i| [x[i], Math.log10(z[i].abs.round(2))] })
flush_graph(@filenames[2], range.each_with_index.map{ |e, i| [x[i], c[i].round(2)] })


commands = %Q(

set terminal png medium size 1200 800
set output 'foobar.png'
# set xrange [-10:10]
# set yrange [-100:100]

plot \
'#{@filenames[0]}.txt' with lines, \
'#{@filenames[1]}.txt' with impulses, \
'#{@filenames[2]}.txt' with lines

)

gnuplot(commands)
`open foobar.png`