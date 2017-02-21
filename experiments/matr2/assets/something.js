$ = (selector, el) => { if (!el) {el = document;} return el.querySelector(selector); }

Vue.component("some", {
	delimiters: ['${', '}'],
	template: "#some",
	data: {
		test: "qweqwe"
	}
})

new Vue({ delimiters: ['${', '}'], el: '#app' });

const ws = new WebSocket("ws://localhost:8001/ws");
ws.onmessage = (msg) => { console.log(msg); };