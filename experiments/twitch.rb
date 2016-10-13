require 'socket'

abort("Not enough arguments") if ARGV.count == 0

name = ARGV[0]
s = TCPSocket.open "irc.chat.twitch.tv", 6667
s.puts "PASS oauth:7mlkb36s7gtzyycdsfo45ydsr4ltqf"
s.puts "NICK e4stw00d"
s.puts "JOIN ##{name}"

puts "Move on"

begin
	while line = s.gets
		temp = line.split("PRIVMSG ##{name} :")
		if temp.count > 1
			nick = temp[0].split('!')[0][1..-1]
			text = temp[1]

			puts "#{nick} - #{text}"
		end
	end
rescue SystemExit, Interrupt
	s.close
end