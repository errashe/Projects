class MainController < ApplicationController
  def index
  	@matches = Match.order("match_time DESC").limit(10)
  end

  def updater_gate
  	if request.user_agent == "updater"
  		render plain: Match.parse_dotabuff
  	else
  		redirect_to root_path
  	end
  end
end
