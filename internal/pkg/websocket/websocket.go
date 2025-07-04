package websocket

import (
	"errors"
	"fmt"

	"github.com/macabot/hypp/js"
	"github.com/macabot/hypp/window"
)

var (
	ErrWebSocket = errors.New("WebSocket error")
)

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

// TODO Pass []byte to callback?
func (ws WebSocket) OnMessage(onMessage func(e js.Value)) {
	ws.Value.Set("onmessage", js.FuncOf(func(this js.Value, args []js.Value) any {
		window.Console().Debug("<<< Receive WebSocket message:", args[0].Get("data").String())
		onMessage(args[0])
		return nil
	}))
}

func (ws WebSocket) OnError(onError func(err error)) {
	ws.Value.Set("onerror", js.FuncOf(func(this js.Value, args []js.Value) any {
		onError(fmt.Errorf("%w: %s", ErrWebSocket, args[0].Get("message").String()))
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
	window.Console().Debug(">>> Send WebSocket message:", string(data))
	uint8Array := js.Global().Get("Uint8Array").New(len(data))
	js.CopyBytesToJS(uint8Array, data)
	ws.Value.Call("send", uint8Array)
}

func (ws WebSocket) Close() {
	window.Console().Debug("Closing WebSocket")
	ws.Value.Call("close")
}
