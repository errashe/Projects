class AddGamedTime < ActiveRecord::Migration[5.0]
	def change
		add_column :matches, :match_time, :datetime
	end
end
