class App < Sinatra::Base
	configure do
		register Sinatra::Flash
		register Sinatra::Reloader if development?

		set :views, Proc.new { File.join(root, "..", "views") }
		set :public_folder, Proc.new { File.join(root, "..", "public") }

		set :session_secret, "SUPER SECRET STRING"
		set :sessions, :enable
	end

	helpers do
		def partial(name); erb name, :layout => false; end
		def user; User.get(session[:user]); end
	end

	post "/login" do
		if user = User.first(params)
			session[:user] = user[:id]
		else
			flash[:error] = "Wrong auth data"
		end
		redirect to request.referrer
	end

	get "/logout" do
		session[:user] = nil
		redirect to request.referrer
	end

	get "/" do
		@sections = Section.all
		erb :sections
	end

	get "/subsection/:id" do
		@subsection = Subsection.get(params[:id])
		erb :themes
	end

	get "/subsection/:id/create_theme" do
		erb :create_theme
	end

	put "/subsection/:id/create_theme" do
		@subsection = Subsection.get(params[:id])
		@subsection.themes.create(:user_id => user[:id], :name => params[:name], :text => params[:name])
		redirect to "/subsection/#{params[:id]}"
	end

	get "/theme/:id" do
		@theme = Theme.get(params[:id])
		erb :messages
	end

	put "/theme/:id" do
		@theme = Theme.get(params[:id])
		@theme.messages.create(:user_id => user[:id], :text => params[:text])
		redirect to request.referrer
	end

	get "/tree" do
		@sections = Section.all
		erb :tree
	end
end