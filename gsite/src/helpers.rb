helpers do
	def db
		settings.db
	end

	def check_user(login, password)
		r.table("users").filter{ |user| user["email"].eq(login) && user["password"].eq(password) }.run(db)
	end

	def current_user
		r.table("users").get(session[:id] || "").run(db)
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