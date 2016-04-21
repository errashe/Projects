require 'active_record'
ActiveRecord::Base.establish_connection({adapter: "sqlite3", database: "users.db"})
class User < ActiveRecord::Base; end

def snd(command)
	@u = User.first
	@u.command = command
	@u.save
end

def n1
	@u = User.first
	@u.mark = "256"
	@u.library = "hookx64.dll"
	@u.x = "1"
	@u.path = "C:\\library1.dll"
	@u.process = "notepad.exe"
	@u.save
end

def n2
	@u = User.first
	@u.mark = "255"
	@u.library = "hookx86.dll"
	@u.x = "0"
	@u.path = "C:\\library2.dll"
	@u.process = "notepad++.exe"
	@u.save
end

def list_dll
	Dir.glob("dll/*").each_with_index.map { |e, i| printf("%d, %s\n", i, e) }
end

def custom_dll(num, process)
	@u = User.first
	@files = Dir.glob("dll/*")
	@filename = File.basename File.open(@files[num])
	@name = File.basename @filename, ".dll"
	@razr = @name[5..10]
	@mark = @razr[2..5]
	@razr = @razr[0..1]

	@u.mark = @mark
	@u.library = @files[num]
	@u.x = @razr == "64" ? "1" : "0"
	@u.path = "C:\\lib#{num}.dll"
	@u.process = process
	@u.save
end