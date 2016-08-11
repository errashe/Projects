class CreatePages < ActiveRecord::Migration[5.0]
  def change
    create_table :pages do |t|

    	t.string :title
    	t.boolean :show_title, :default => true
    	t.text :text

      t.timestamps
    end
  end
end
