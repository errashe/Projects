#!/usr/bin/env ruby

require 'sinatra'
require 'sinatra-websocket'

require 'json'

def rand_nick
	(0...10).map { ('a'..'z').to_a[rand(26)] }.join
end

set :sessions, :enable
set :session_secret, '*&(^B234'

set :bind, '192.168.1.140'
set :sockets, []

get '/' do
	if !request.websocket?
		erb :index
	else
		request.websocket do |ws|
			ws.onopen do
				session[:nick] = rand_nick
				settings.sockets << ws
			end
			ws.onmessage do |msg|
				# EM.next_tick { settings.sockets.each{|s| s.send(msg) } }
				settings.sockets.each{|s| s.send([session[:nick], msg].to_json) }
			end
			ws.onclose do
				# settings.sockets.each_with_index{|s, i| settings.sockets.delete(i) if settings.sockets[i][1] == ws}
				settings.sockets.delete(ws)
			end
		end
	end
end

__END__

@@ index
<html>
<body>
	<h1>Chat</h1>
	<form id="form">
		<input type="text" id="input" placeholder="Print here!"/>
	</form>
	<div id="msgs"></div>
</body>

<script type="text/javascript">
	window.onload = function(){
		(function(){
			var show = function(el){
				return function(msg){ el.innerHTML = msg + '<br />' + el.innerHTML; }
			}(document.getElementById('msgs'));

			var ws       = new WebSocket('ws://' + window.location.host + window.location.pathname);
			// ws.onopen    = function()  { show('websocket opened'); };
			// ws.onclose   = function()  { show('websocket closed'); }
			ws.onmessage = function(m) { let msg = JSON.parse(m.data); show(`<b>${msg[0]}</b> - ${msg[1]}`); };

			var sender = function(f){
				var input     = document.getElementById('input');
				f.onsubmit    = function(){
					ws.send(input.value);
					f.reset();
					return false;
				}
			}(document.getElementById('form'));
		})();
	}
</script>
</html>