class App < Sinatra::Base
	configure do
		register Sinatra::Reloader if development?
	end

	get "/" do
		Time.now.to_s
	end
end