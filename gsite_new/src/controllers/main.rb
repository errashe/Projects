class Main < App

	get '/' do
		Time.now.to_f.to_s
	end

end