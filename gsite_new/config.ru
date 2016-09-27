require 'rubygems'
require 'bundler'

Bundler.require(:default)

require_all './src/controllers'

map('/') { run Main }