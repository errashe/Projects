class CreateSfiles < ActiveRecord::Migration[5.0]
  def change
    create_table :sfiles do |t|

    	t.string :mark
    	t.string :name
    	t.string :title
    	t.string :text

      t.timestamps
    end
  end
end
