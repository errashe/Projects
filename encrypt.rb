@freq = Array.new(4) { Array.new(256, 0) }
@file = File.read("encrypted.txt").force_encoding("windows-1251")

current = 0
@file.each_byte do |ch|
	@freq[current][ch] += 1
	current = current == 3 ? 0 : current + 1
end

@key = []
@freq.each do |key|
	@key.push key.index(key.max) ^ ' '.ord
end

puts "#{@key}"

# @out_file = File.open("dectypted.txt", "w")
# current = 0
# @file.each_byte do |ch|
# 	@out_file << (ch ^ @key[current]).chr
# 	current = current == 3 ? 0 : current + 1
# end
# @out_file.close
