module ApplicationHelper
	def current_user
		User.find(session[:user]) if !session[:user].nil?
	end

	def admin?
		current_user.role == "admin" if current_user
	end

	def jsload(name)
	"<script type='text/javascript' src='/#{name}.js'></script>"
	end
end
