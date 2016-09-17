get "/" do
	get_check_page("main")
end

get "/profile" do
	current_user.inspect
end

post "/login" do
	user = check_user(params["email"], params["password"])

	if user
		session[:_id] = user["_id"]
		flash[:success] = "Вы были успешно авторизованы."
	else
		flash[:error] = "Что-то не так"
	end
	redirect to("/")
end

get "/registration" do
	erb :registration
end

post "/registration" do
	if params["password"] == params["password-conf"]
		begin
			ins = db[:users].insert_one(
				{
					email: params["email"],
					password: params["password"],
					fio: params["fio"]
				}
			)

			if ins
				flash[:success] = "Вы успешно зарегистрировались, подтвердите вашу почту."
			else
				flash[:error] = "Что-то не так с базой, придется подождать"
			end
		rescue Mongo::Error => e
			if e.message.include? "E11000"
				flash[:error] = "Такая почта уже существует"
			end
		end
		redirect to("/")
	else
		flash[:error] = "Что-то пошло не так..."
	end
	redirect to("/registration")
end

get "/logout" do
	session[:_id] = nil
	redirect to("/")
end

namespace "/pages" do

	get "/:name" do
		get_check_page(params[:name])
	end

end

namespace "/admin" do
	before do
		authorize_admin!
	end

	get do
		erb :"admin/main"
	end

	get "/pages" do
		@pages = db[:pages].find()
		erb :"admin/pages/list"
	end

	get "/pages/new" do
		erb :"admin/pages/new"
	end

	get "/pages/:id/edit" do
		# "Edit this - %s (Show form here)" % params[:id]
		@page = db[:pages].find({:_id => BSON::ObjectId(params[:id])}).first
		erb :"admin/pages/edit"
	end

	put "/pages/:id/edit" do
		ins = db[:pages].update_one({:_id => BSON::ObjectId(params[:id])}, {'$set' => {
			:title => params[:title],
			:name => params[:name],
			:body => params[:body]
		}})

		if ins
			flash[:success] = "Успешно сохранено"
		else
			flash[:error] = "Что-то произошло"
		end

		redirect to("/admin/pages")
	end

	get "/pages/:id/delete" do
		# "Delete this - %s" % params[:id]
		ins = db[:pages].delete_one(:_id => BSON::ObjectId(params[:id]))

		if ins
			flash[:success] = "Успешно удалено"
		else
			flash[:error] = "Что-то произошло"
		end

		redirect to("/admin/pages")
	end
end
