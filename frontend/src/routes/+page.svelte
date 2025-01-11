<script>
	let chatlog = '';

	let websocket = new WebSocket('ws://localhost:3000/ws');
	websocket.onmessage = (e) => {
		console.log('Message from server:', e.data);
		handleIncomingMessage(e.data);
	};
	function handleIncomingMessage(message) {
		chatlog += message + '\n';
	}

	let message = '';
	function sendMsg() {
		websocket.send(message);
		message = '';
	}
</script>

<form onsubmit={sendMsg}>
	<input
		required
		placeholder="Send a Message"
		type="text"
		id="message-input"
		bind:value={message}
	/>
	<button type="submit">Send</button>
</form>

<textarea readonly bind:value={chatlog}></textarea>
