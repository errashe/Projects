require 'pp'

@map = File.read("map.txt").split.map{|q| q.split('')}

def find(s)
	@map.each_with_index { |row, i| 
		row.each_with_index { |cell, j| 
			return [i, j] if cell == s
		}
	}
end

def astar(arr, start, goal)
	
end


p astar(@map, find("S"), find("F"))
pp @map