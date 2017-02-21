$ = (selector, el) => { if (!el) {el = document;} return el.querySelector(selector); }

app = new Vue({
	delimiters: ['${', '}'], el: '#app',
	created: () => {

	},
	data: {
		message: "",
		messages: []
	},
	methods: {
		send: () => {
			ws.send(app.$data.message);
			app.$data.message = "";
		}
	}
});

ws = new WebSocket("ws://localhost:8080/ws");
ws.onmessage = (msg) => { app.$data.messages.unshift(msg.data); }