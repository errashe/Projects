class AdminController < ApplicationController
	include ApplicationHelper
	before_action :check_admin

	def index
	end

	def list
		@pages = Page.all
	end

	def static_create
	end

	def static_update
		@up = Page.find_by_mark(params[:mark])
		@up.title = params[:title]
		@up.text = params[:text]
		@up.show_title = params[:show_title] == "on" ? true : false;
		if @up.save
			redirect_to static_path(@up.mark)
		end
	end

	def static_save
		@page = Page.new
		@page.mark = params[:mark]
		if @page.save
			redirect_to page_list_path
		end
	end

	def static_delete
		@page = Page.find_by_mark(params[:mark])
		@page.delete
		redirect_to page_list_path
	end

	def file_list
		@files = Sfile.all.order(id: :DESC)
	end

	def file_create
	end

	def file_save
		s = Sfile.new
		uploaded_io = params[:sfile]
		mark = Digest::MD5.hexdigest("#{uploaded_io.original_filename}#{Time.now.to_i}")

		File.open(Rails.root.join('public', 'uploads', mark), 'wb') do |file|
			# uploaded_io.original_filename
			file.write(uploaded_io.read)
		end

		s.mark = mark
		s.name = uploaded_io.original_filename
		s.title = params[:title]
		s.text = params[:text]
		if s.save
			redirect_to file_list_path
		end
	end

	def file_delete
		@s = Sfile.find_by_mark(params[:mark])
		File.delete(Rails.root.join('public', 'uploads', @s.mark))
		if @s.destroy
			redirect_to file_list_path
		end
	end

	private
	def check_admin
		redirect_to root_path if !admin?
	end
end
