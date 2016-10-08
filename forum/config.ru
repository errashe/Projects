require "sinatra"
require "sinatra/reloader" if development?
require "sinatra/flash"

require "sqlite3"
require "require_all"

require_all "./src/*.rb"
# require "./src/models.rb"
# require "./src/main.rb"

use Rack::MethodOverride

run App