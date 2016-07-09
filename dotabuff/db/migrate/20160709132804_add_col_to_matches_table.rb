class AddColToMatchesTable < ActiveRecord::Migration[5.0]
	def change
		add_column :matches, :hash, :string
	end
end
