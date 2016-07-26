class MainController < ApplicationController
	before_action :params_permit

	def index
	end

	def static

	end

	private
	def params_permit
		params.permit(:page)
	end
end
