Rails.application.routes.draw do
  root "main#index"

  get "login", as: "login", to: "main#login"
  post "login", as: "auth", to: "main#auth"
  get "logout", as: "logout", to: "main#logout"

  post "update", as: "update", to: "main#update"
  get ":page", as: "static", to: "main#static"

  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
end
