enable :sessions

configure do
	set :erb, :layout => true
	set :views, File.dirname(__FILE__) + '/../views'
	set :public_folder, File.dirname(__FILE__) + '/../public'

	set :session_secret, '*&(^B234'

	set :bind, "192.168.1.31" if development?
	set :bind, "helper-stud.ru" if production?
	set :port, "80" if production?

	set :db, Mongo::Client.new('mongodb://localhost:27017/momsite')
end

Mongo::Logger.logger.level = Logger::WARN