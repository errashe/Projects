class AddPlayerUidToMatch < ActiveRecord::Migration[5.0]
  def change
  	add_column :matches, :puid, :integer, :limit => 8
  end
end
