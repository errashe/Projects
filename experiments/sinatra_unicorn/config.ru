require "sinatra"
require "sinatra/reloader" if development?
require "sinatra/namespace"

require "./src/routes"

run App