package websocket

import "github.com/macabot/hypp/js"

type WebSocket struct {
	js.Value
}

func NewWebSocket(url string) WebSocket {
	return WebSocket{js.Global().Get("WebSocket").New(url)}
}

func (ws WebSocket) OnOpen(onOpen func()) {
	ws.Value.Set("onopen", js.FuncOf(func(this js.Value, args []js.Value) any {
		onOpen()
		return nil
	}))
}

func (ws WebSocket) OnMessage(onMessage func(e js.Value)) {
	ws.Value.Set("onmessage", js.FuncOf(func(this js.Value, args []js.Value) any {
		onMessage(args[0])
		return nil
	}))
}

func (ws WebSocket) OnError(onError func(e js.Value)) {
	ws.Value.Set("onerror", js.FuncOf(func(this js.Value, args []js.Value) any {
		onError(args[0])
		return nil
	}))
}

func (ws WebSocket) OnClose(onClose func(e js.Value)) {
	ws.Value.Set("onclose", js.FuncOf(func(this js.Value, args []js.Value) any {
		onClose(args[0])
		return nil
	}))
}

func (ws WebSocket) Send(data []byte) {
	uint8Array := js.Global().Get("Uint8Array").New(len(data))
	js.CopyBytesToJS(uint8Array, data)
	ws.Value.Call("send", uint8Array)
}

func (ws WebSocket) Close() {
	ws.Value.Call("close")
}
