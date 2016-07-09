Rails.application.routes.draw do
	root 'main#index'
	get 'ug', to: "main#updater_gate"

  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
end
