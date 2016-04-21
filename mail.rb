require 'pony'

Pony.mail({
	:to => 'but3nko@gmail.com',
	:html_body => '
	<h1>Hello, world!</h1>
	I gonna try to use <b>some html magic</b> here!
	<img src="http://lorempixel.com/400/200/" />	
	',
	:via => :smtp,
	:via_options => {
		:address              => 'smtp.gmail.com',
		:port                 => '587',
		:enable_starttls_auto => true,
		:user_name            => 'but3nko@gmail.com',
		:password             => 'Fk3rc4ylh<en3yr0',
    :authentication       => :plain, # :plain, :login, :cram_md5, no auth by default
    :domain               => "localhost.localdomain" # the HELO domain provided by the client to the server
  }
  })