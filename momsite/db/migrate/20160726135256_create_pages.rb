class CreatePages < ActiveRecord::Migration[5.0]
  def change
    create_table :pages do |t|

    	t.string :label
    	t.string :type
    	t.text :text

      t.timestamps
    end
  end
end
