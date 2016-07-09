class CreateGamers < ActiveRecord::Migration[5.0]
  def change
    create_table :gamers do |t|
      t.string :nick
      t.integer :uid

      t.timestamps
    end
  end
end
