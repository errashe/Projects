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
				flash[:success] = "Вы успешно зарегистрировались."
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

namespace "/guestbook" do

	before do
		pass if user?
		flash[:error] = "Авторизуйтесь, чтобы получить доступ сюда"
		redirect to("/")
	end

	get do
		erb :guestbook
	end

	put do
		Pony.mail(
		{
			:subject => "Сообщение из гостевой книги",
			:html_body => params[:body],
			:to => 'but3nko@gmail.com',
			:via => :smtp,
			:via_options => {
				:address              => 'smtp.gmail.com',
				:port                 => '587',
				:enable_starttls_auto => true,
				:user_name            => 'but3nko',
				:password             => 'Fk3rc4ylh<en3yr0',
				:authentication       => :plain,
				:domain               => "localhost.localdomain"
			}
		}
		)

		flash[:success] = "Письмо успешно отправлено"
		redirect to("/")

	end

end

namespace "/pages" do

	before do
		pass if user?

		flash[:error] = "Авторизуйтесь, чтобы просмотреть страницы"
		redirect to("/")
	end

	get "/:name" do
		get_check_page(params[:name])
	end

end

namespace "/files" do

	before do
		pass if user?

		flash[:error] = "Авторизуйтесь, чтобы просмотреть страницы"
		redirect to("/")
	end

	get do
		@files = db[:files].find()
		erb :"files/list"
	end

	get "/:id" do
		@file = db[:files].find({:_id => BSON::ObjectId(params[:id])}).first
		erb :"files/one"
	end

	get "/:id/get" do
		@file = db[:files].find({:_id => BSON::ObjectId(params[:id])}).first
		send_file("public/files/%s" % @file[:save_name], :filename => @file[:filename])
	end

end

namespace "/admin" do
	before do
		authorize_admin!
	end

	get do
		erb :"admin/main"
	end

	namespace "/pages" do

		get do
			@pages = db[:pages].find()
			erb :"admin/pages/list"
		end

		get "/new" do
			erb :"admin/pages/new"
		end

		put do
			begin
				ins = db[:pages].insert_one({
					:title => params[:title],
					:name => params[:name],
					:body => params[:body]
					})

				if ins
					flash[:success] = "Успешно создана"
				else
					flash[:error] = "Что-то произошло"
				end
			rescue Mongo::Error => e
				if e.message.include? "E11000"
					flash[:error] = "Такая страница уже существует"
				end
			end

			redirect to("/admin/pages")
		end

		get "/:id/edit" do
			@page = db[:pages].find({:_id => BSON::ObjectId(params[:id])}).first
			erb :"admin/pages/edit"
		end

		patch "/:id" do
			begin
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
			rescue Mongo::Error => e
				if e.message.include? "E11000"
					flash[:error] = "Такая страница уже существует"
				end
			end

			redirect to("/admin/pages")
		end

		get "/:id/delete" do
			ins = db[:pages].delete_one(:_id => BSON::ObjectId(params[:id]))

			if ins
				flash[:success] = "Успешно удалено"
			else
				flash[:error] = "Что-то произошло"
			end

			redirect to("/admin/pages")
		end

	end

	namespace "/files" do

		get do
			@files = db[:files].find()
			erb :"admin/files/list"
		end

		get "/new" do
			erb :"admin/files/new"
		end

		put do
			name = hash_date

			File.open('public/files/%s' % name, 'w') do |f|
				f.write(params[:file][:tempfile].read)
			end

			ins = db[:files].insert_one(
			{
				:name => params[:name],
				:description => params[:description],
				:filename => params[:file][:filename],
				:save_name => name
			}
			)

			if ins
				flash[:success] = "Файл был успешно загружен"
			else
				flash[:error] = "Что-то пошло не так"
			end

			redirect to("/admin/files")
		end

		get "/:id/delete" do
			files = db[:files].find({:_id => BSON::ObjectId(params[:id])})
			file = files.first

			File.delete("public/files/%s" % file[:save_name])
			ins = files.delete_one

			if ins
				flash[:success] = "Файл успешно удален"
			else
				flash[:error] = "Что-то пошло не так"
			end

			redirect to("/admin/files")
		end
	end

	namespace "/users" do

		get do
			@users = db[:users].find()
			erb :"admin/users/list"
		end

	end

end