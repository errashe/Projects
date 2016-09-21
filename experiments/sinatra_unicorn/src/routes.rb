class App < Sinatra::Base

	configure do
		register Sinatra::Reloader
		set :session, :enable
	end

	helpers do
		def test
			Time.now.to_s
		end
	end

	get "/" do
		test
	end

end