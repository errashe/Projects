class CreateMatches < ActiveRecord::Migration[5.0]
	def change
		create_table :matches do |t|
			t.string :hero
			t.boolean :stats
			t.integer :uid

			t.timestamps
		end
	end
end
