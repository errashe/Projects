require 'thread'
semaphore = Mutex.new

@shared = 0.0

a = Thread.new do
	loop do
		semaphore.synchronize { @shared = rand(-100.0..100.0) }
		sleep(Float::MIN)
	end
end

b = Thread.new do
	loop do
		p @shared
		sleep(Float::MIN)
	end
end

a.join
b.join


# trap( "INT" ) do
# 	@exit_requested = true
# 	a.kill
# 	b.kill
# 	exit
# end