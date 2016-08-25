require 'mechanize'

@a = Mechanize.new
@a.get("https://m.vk.com")

form =  @a.page.form
form.email = "e4stw00d@icloud.com"
form.pass = "the{}Pre4cher"
form.submit

@a.get("/audio")

pp @a.page