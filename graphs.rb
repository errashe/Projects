f = Float::INFINITY

input = [
	[0, 10, 18, f, 14, 9],
	[10, 0, 12, f, f, f],
	[18, 12, 0, 6, 14, f],
	[f, f, 6, 0, 10, 15],
	[14, f, 14, 10, 0, 7],
	[9, f, f, 15, 7, 0]
]

output = Array.new(input.length) { Array.new(input.length) }

(input.length-1).times do |q|

	input.each_with_index do |arr, i|
		arr.each_with_index do |arr2, j|
			output[i][j] = [input[i][q]+input[q][j], input[i][j]].min
		end
	end

	input = output
	p q+1; output.map{ |e| p e }
end