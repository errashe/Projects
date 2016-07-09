class AddUniqInMatches < ActiveRecord::Migration[5.0]
	def change
		add_index :matches, :hash, :unique => true
	end
end
