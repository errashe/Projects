Rails.application.routes.draw do
  root "main#static"

  get "admin", as: "admin", to: "admin#index"
  get "admin/list", as: "page_list", to: "admin#list"
  get "admin/static_create", as: "static_create", to: "admin#static_create"
  post "admin/static_create", as: "static_save", to: "admin#static_save"
  post "admin/static_update", as: "static_update", to: "admin#static_update"
  get "admin/static_delete/:mark", as: "static_delete", to: "admin#static_delete"

  get "profile", as: "profile", to: "main#profile"

  get "login", as: "login", to: "main#login"
  post "login", as: "auth", to: "main#auth"
  get "logout", as: "logout", to: "main#logout"

  get ":page", as: "static", to: "main#static"

  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
end
