helpers do
	def db
		settings.db
	end

	def check_user(email, password)
		db[:users].find({ email: email, password: password })
	end

	def current_user
		db[:users].find({ _id: session[:_id] || 0}).first
	end

	def partial(name)
		erb name, :layout => false
	end

	def styles
		files = Dir["public/assets/css/*"].map do |file|
			"<link rel='stylesheet' href='#{file.gsub(/^public/, "")}'>"
		end

		files.join("\n")
	end

	def scripts
		files = Dir["public/assets/js/*"].map do |file|
			"<script type='text/javascript' src='#{file.gsub(/^public/, "")}'></script>"
		end

		files.join("\n")
	end
end