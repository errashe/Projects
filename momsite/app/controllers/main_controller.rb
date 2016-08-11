class MainController < ApplicationController
	before_action :permit_fields, only: [:static, :update]

	def index
		@page = Page.find_by_mark("index")
		render :static
	end

	def static
		@page = Page.find_by_mark(@params[:page])
	end

	def update
		@up = Page.find_by_mark(@params[:mark])
		@up.text = @params[:text]
		if @up.save
			redirect_to static_path(@up.mark)
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

	private
	def permit_fields
		@params = params.permit(:page, :mark, :text)
	end
end
