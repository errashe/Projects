module ApplicationHelper
	def current_user
		User.find(session[:user]) if !session[:user].nil?
	end
end
