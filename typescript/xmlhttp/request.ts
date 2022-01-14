import { XMLHttpRequest } from 'xmlhttprequest';

function RequestGET() {
	const req = new XMLHttpRequest();
	const url = "https://cloudflare-quic.com/b/ip";

	req.open('GET', url);
	req.setRequestHeader("Accept", "application/json")
	req.setRequestHeader("Content-Type", "application/json")
	req.onload = function () {
		console.log("Start to test GET");
		console.log(req.status);
		console.log(req.responseType);
		console.log(req.response);
		console.log(req.responseText);
		console.log();
	}
	req.send(null);
	// NOTE: req response not work here, if it's async quire
	// console.log(req.status);
	// console.log(req.responseType);
	// console.log(req.response);
	// console.log(req.responseText);
}

function RequestPOST() {
	const req = new XMLHttpRequest();
	const url = "https://cloudflare-quic.com/b/post";

	req.open('POST', url);
	req.setRequestHeader("Accept", "application/json")
	req.setRequestHeader("Content-Type", "application/json")
	req.onload = function () {
		console.log("Start to test POST");
		console.log(req.status);
		console.log(req.responseType);
		console.log(req.response);
		console.log(req.responseText);
		console.log();
	}
	req.send("This is POST test string");
}

function RequestCallback(callback: (string) => void) {
	const req = new XMLHttpRequest();
	const url = "https://cloudflare-quic.com/b/ip";

	req.open('GET', url, false);
	req.setRequestHeader("Accept", "application/json")
	req.setRequestHeader("Content-Type", "application/json")
	req.onload = function () {
		console.log("Test callback text string");
		console.log(req.status);
		console.log(req.responseType);
		console.log(req.response);
		console.log(req.responseText);
		callback(String(req.responseText));
		console.log();
	}
	req.send(null); 
}

RequestGET();
RequestPOST();

let value: string = "";
const callbackFunc = (text: string): void => {
	value = text;
}
RequestCallback(callbackFunc);
console.log("Test callback text string");
console.log(value);