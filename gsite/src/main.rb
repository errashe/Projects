require "sinatra"
require "sinatra/reloader" if development?
require "sinatra/flash"
require "sinatra/namespace"

require "mongo"

require "./src/helpers.rb"
require "./src/config.rb"

require "./src/routes.rb"