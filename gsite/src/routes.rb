get "/" do
	erb :main
end

post "/login" do
	user = check_user(params["email"], params["password"]).first

	if user
		session[:id] = user["id"]
	end
	redirect to("/")
end

get "/registration" do
	erb :registration
end

get "/logout" do
	session[:id] = nil
	redirect to("/")
end