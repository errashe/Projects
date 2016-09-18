helpers do
	def db
		settings.db
	end

	def check_user(email, password)
		db[:users].find({ email: email, password: password }).first
	end

	def current_user
		db[:users].find({ _id: session[:_id] || 0}).first
	end

	def user?
		!!current_user
	end

	def admin?
		current_user && current_user[:role] == "admin"
	end

	def authorize_admin!
		pass if admin?
		flash[:error] = "Нужно авторизоваться"
		redirect to("/")
	end

	def partial(name)
		erb name, :layout => false
	end

	def hash_date
		md5 = Digest::MD5.new
		md5.update(Time.now.to_f.to_s)
		md5.hexdigest
	end

	def get_check_page(name)
		@page = db[:pages].find({name: name}).first
		if @page
			erb :"pages/page"
		else
			flash[:error] = "Страница не найдена"
			redirect to("/")
		end
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