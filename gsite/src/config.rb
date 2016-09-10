include RethinkDB::Shortcuts

enable :sessions

configure do
	set :erb, :layout => true
	set :views, File.dirname(__FILE__) + '/../views'
	set :public_folder, File.dirname(__FILE__) + '/../public'

	set :session_secret, '*&(^B234'

	set :db, r.connect(:host=>"localhost", :port=>28015)
	settings.db.use("momsite")
end