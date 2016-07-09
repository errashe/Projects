# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# Note that this schema.rb definition is the authoritative source for your
# database schema. If you need to create the application database on another
# system, you should be using db:schema:load, not running all the migrations
# from scratch. The latter is a flawed and unsustainable approach (the more migrations
# you'll amass, the slower it'll run and the greater likelihood for issues).
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 20160709160506) do

  create_table "gamers", force: :cascade do |t|
    t.string   "nick"
    t.integer  "uid",        limit: 8
    t.datetime "created_at",           null: false
    t.datetime "updated_at",           null: false
  end

  create_table "matches", force: :cascade do |t|
    t.string   "hero"
    t.boolean  "stats"
    t.integer  "uid",        limit: 8
    t.datetime "created_at",           null: false
    t.datetime "updated_at",           null: false
    t.string   "uhash"
    t.datetime "match_time"
    t.integer  "puid",       limit: 8
    t.index ["uhash"], name: "index_matches_on_uhash", unique: true
  end

end
