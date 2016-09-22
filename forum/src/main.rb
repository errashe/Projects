class App < Sinatra::Base
	configure do
		register Sinatra::Reloader if development?

		set :views, Proc.new { File.join(root, "..", "views") }
		set :public_folder, Proc.new { File.join(root, "..", "public") }
	end

	helpers do
		def partial(name); erb name, :layout => false; end
	end

	get "/" do
		@sections = Section.all
		erb :sections
	end

	get "/section/:id" do
		@subsections = Subsection.all(:section_id => params[:id])
		erb :subsections
	end

	get "/subsection/:id" do
		@themes = Theme.all(:subsection_id => params[:id])
		erb :themes
	end

	get "/theme/:id" do
		@theme = Theme.get(params[:id])
		@messages = Message.all(:theme_id => params[:id])
		erb :messages
	end

	get "/tree" do
		@sections = Section.all
		erb :tree
	end
end