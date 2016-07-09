class Gamer < ApplicationRecord
	has_many :matches, foreign_key: "puid", primary_key: "uid"
end
