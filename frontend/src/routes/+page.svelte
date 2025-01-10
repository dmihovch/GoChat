<script>
	let websocket = new WebSocket('ws://localhost:3000/ws');
	websocket.onmessage = (e) => {
		console.log('Message from server:', e.data);
	};

	function sendMsg() {
		websocket.send('Hello from sveltekit');
	}
</script>

<button on:click={sendMsg}>send msg</button>
