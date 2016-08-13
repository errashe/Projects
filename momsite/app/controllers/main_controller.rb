class MainController < ApplicationController

	def static
		p = params[:page] || "index"
		@page = Page.find_by_mark(p)
		if @page.nil?
			render :file => "#{Rails.root}/public/404", :layout => false, :status => 404
		end
	end

	def login
	end

	def auth
		# render plain: params.permit(:email, :password).inspect
		@user = User.find_by_email_and_password(params[:email], params[:password])
		if !@user.nil?
			session[:user] = @user.id
		end
		redirect_to root_path
	end

	def logout
		session.delete(:user)
		redirect_to root_path
	end

	def profile
	end
end
