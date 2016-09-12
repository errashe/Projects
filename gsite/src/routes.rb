get "/" do
	erb :main
end

post "/login" do
	user = check_user(params["email"], params["password"]).first

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