require "data_mapper"
require "dm-migrations" if development?

DataMapper.setup(:default, "sqlite3://#{Dir.pwd}/dev.db")

class User
	include DataMapper::Resource

	property :id, Serial
	property :login, String
	property :password, String
	property :group_id, Integer

	has n, :sections
	has n, :subsections
	has n, :themes
	has n, :messages
end

class Section
	include DataMapper::Resource

	property :id, Serial
	property :name, String

	belongs_to :user
	has n, :subsections
end

class Subsection
	include DataMapper::Resource

	property :id, Serial
	property :name, String

	belongs_to :user
	belongs_to :section
	has n, :themes
end

class Theme
	include DataMapper::Resource

	property :id, Serial
	property :name, String
	property :text, Text

	belongs_to :user
	belongs_to :subsection
	has n, :messages
end

class Message
	include DataMapper::Resource

	property :id, Serial
	property :text, Text

	belongs_to :user
	belongs_to :theme
end

DataMapper.finalize
DataMapper.auto_upgrade! if development?