require 'net/pop'

Net::POP3.enable_ssl(OpenSSL::SSL::VERIFY_NONE)
Net::POP3.start('pop.gmail.com', 995, 'defensuer@gmail.com', 'Open!3451') do |pop|
	if pop.mails.empty?
		puts 'No mail.'
	else
		i = 1
		pop.delete_all do |m|
			File.open("inbox/#{i}.txt", 'w') do |f|
				f.write m.pop
			end
			i += 1
		end
	end
end