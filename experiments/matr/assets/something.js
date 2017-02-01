$ = (selector, el) => { if (!el) {el = document;} return el.querySelector(selector); }

app = new Vue({
	delimiters: ['${', '}'], el: '#tags',
	data: {
		messages: []
	},
	mounted: () => { $("#tags").style.display = "block"; },
	methods: {
		addItem: (event) => { ws.send(Math.random()); },
		playMusic: (e) => { app.$refs.player.play(); },
		stopMusic: (e) => { app.$refs.player.pause(); app.$refs.player.currentTime = 0; }
	}
});

const ws = new WebSocket("ws://" + window.location.host + "/ws");
ws.onmessage = (msg) => { app.messages.unshift(msg.data); };