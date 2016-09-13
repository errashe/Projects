#!/usr/bin/env ruby

require 'sinatra'
require 'sinatra-websocket'

set :sockets, []

get '/' do
	if !request.websocket?
		erb :index
	else
		request.websocket do |ws|
			ws.onopen do
				settings.sockets << ws
			end
			ws.onmessage do |msg|
				# EM.next_tick { settings.sockets.each{|s| s.send(msg) } }
				settings.sockets.each{|s| s.send(msg) }
			end
			ws.onclose do
				settings.sockets.delete(ws)
			end
		end
	end
end

__END__
@@ index
<html>
<body>
	<h1>Simple Echo & Chat Server</h1>
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
			ws.onmessage = function(m) { show(m.data); };

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