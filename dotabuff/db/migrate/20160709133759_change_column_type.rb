class ChangeColumnType < ActiveRecord::Migration[5.0]
	def change
		change_column :matches, :uid, :integer, :limit => 8
		change_column :gamers, :uid, :integer, :limit => 8
	end
end
