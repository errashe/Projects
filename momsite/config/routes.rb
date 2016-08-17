Rails.application.routes.draw do
  root "main#static"

  get "files_list", as: "files_list", to: "files#index"
  get "files_list/:offset", as: "files_list_paged", to: "files#index"
  get "files/:mark", as: "file_link", to: "files#get"
  get "files/download/:mark", as: "file_download", to: "files#download"

  get "admin", as: "admin", to: "admin#index"

  get "admin/list", as: "page_list", to: "admin#list"
  get "admin/static_create", as: "static_create", to: "admin#static_create"
  post "admin/static_create", as: "static_save", to: "admin#static_save"
  post "admin/static_update", as: "static_update", to: "admin#static_update"
  get "admin/static_delete/:mark", as: "static_delete", to: "admin#static_delete"

  get "admin/file_list", as: "file_list", to: "admin#file_list"
  get "admin/file_create", as: "file_create", to: "admin#file_create"
  post "admin/file_create", as: "file_save", to: "admin#file_save"
  get "admin/file_delete/:mark", as: "file_delete", to: "admin#file_delete"

  get "profile", as: "profile", to: "main#profile"

  get "login", as: "login", to: "main#login"
  post "login", as: "auth", to: "main#auth"
  get "logout", as: "logout", to: "main#logout"

  get ":page", as: "static", to: "main#static"

  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
end
