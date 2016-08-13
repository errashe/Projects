class ChangeDefaultTitle < ActiveRecord::Migration[5.0]
  def change
  	change_column_default :pages, :show_title, false
  end
end
