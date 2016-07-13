require 'socket'

name = "nalcs1"
s = TCPSocket.open "irc.chat.twitch.tv", 6667
s.puts "PASS oauth:7mlkb36s7gtzyycdsfo45ydsr4ltqf"
s.puts "NICK e4stw00d"
s.puts "JOIN ##{name}"

while line = s.gets
	temp = line.split("PRIVMSG ##{name} :")
	if temp.count > 1
		nick = temp[0].split('!')[0][1..-1]
		text = temp[1]

		puts "#{nick} - #{text}"
	end
end

s.close