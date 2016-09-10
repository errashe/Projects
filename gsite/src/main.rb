require "sinatra"
require "sinatra/reloader" if development?
require "rethinkdb"

require "./src/helpers.rb"
require "./src/config.rb"

require "./src/routes.rb"