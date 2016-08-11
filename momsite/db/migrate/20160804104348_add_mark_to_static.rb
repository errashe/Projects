class AddMarkToStatic < ActiveRecord::Migration[5.0]
  def change
  	add_column :pages, :mark, :string
  end
end
