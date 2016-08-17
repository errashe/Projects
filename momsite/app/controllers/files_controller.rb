class FilesController < ApplicationController
	def index
		offset = params[:offset] || 0
		@files = Sfile.offset(offset).limit(10).order(id: :DESC)
	end

	def get
		@file = Sfile.find_by_mark(params[:mark])
	end

	def download
		@file = Sfile.find_by_mark(params[:mark])
		send_file Rails.root.join('public', 'uploads', @file.mark), :filename => @file.name
	end
end
