class MainController < ApplicationController
  def index
  end

  def updater_gate
  	Match.parse_dotabuff if request.user_agent == "updater"
  end
end
