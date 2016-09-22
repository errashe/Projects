require "sinatra"
require "sinatra/reloader" if development?

require "data_mapper"
require "require_all"

require_all "./src/*.rb"

run App