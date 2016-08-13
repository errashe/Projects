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

	private
	def check_admin
		redirect_to root_path if !admin?
	end
end
