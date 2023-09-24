package webrtc

import (
	"errors"
	"syscall/js"
)

// await awaits a Promise.
// Based on https://stackoverflow.com/a/68427221
func await(awaitable js.Value) (js.Value, js.Value) {
	then := make(chan js.Value)
	defer close(then)
	thenFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		then <- args[0]
		return nil
	})
	defer thenFunc.Release()

	catch := make(chan js.Value)
	defer close(catch)
	catchFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		catch <- args[0]
		return nil
	})
	defer catchFunc.Release()

	awaitable.Call("then", thenFunc).Call("catch", catchFunc)

	select {
	case result := <-then:
		return result, js.Null()
	case err := <-catch:
		return js.Null(), err
	}
}

type ICEServer struct {
	URLs string
}

func (s ICEServer) Value() map[string]any {
	return map[string]any{
		"urls": s.URLs,
	}
}

type PeerConnectionConfig struct {
	ICEServers []ICEServer
}

func (c PeerConnectionConfig) Value() map[string]any {
	servers := make([]any, len(c.ICEServers))
	for i, s := range c.ICEServers {
		servers[i] = s.Value()
	}
	return map[string]any{
		"iceServers": servers,
	}
}

var DefaultPeerConnectionConfig = PeerConnectionConfig{
	ICEServers: []ICEServer{{
		URLs: "stun:stun.l.google.com:19302",
	}},
}

type PeerConnection js.Value

func NewPeerConnection(config PeerConnectionConfig) PeerConnection {
	return PeerConnection(js.Global().Get("RTCPeerConnection").New(config.Value()))
}

type DataChannelOptions struct {
	Negotiated bool
	ID         int
}

func (o DataChannelOptions) Value() map[string]any {
	return map[string]any{
		"negotiated": o.Negotiated,
		"id":         o.ID,
	}
}

var DefaultDataChannelOptions = DataChannelOptions{
	Negotiated: true,
	ID:         0,
}

func (c PeerConnection) CreateDataChannel(label string, options DataChannelOptions) DataChannel {
	return DataChannel(js.Value(c).Call("createDataChannel", label, options.Value()))
}

func (c PeerConnection) SetOnICEConnectionStateChange(onICEConnectionStateChange func()) {
	js.Value(c).Set("oniceconnectionstatechange", js.FuncOf(func(this js.Value, args []js.Value) any {
		onICEConnectionStateChange()
		return nil
	}))
}

func (c PeerConnection) SetOnConnectionStateChange(onConnectionStateChange func()) {
	js.Value(c).Set("onconnectionstatechange", js.FuncOf(func(this js.Value, args []js.Value) any {
		onConnectionStateChange()
		return nil
	}))
}

func (c PeerConnection) ICEConnectionState() string {
	return js.Value(c).Get("iceConnectionState").String()
}

func (c PeerConnection) SetLocalDescription(description SessionDescription) {
	promise := js.Value(c).Call("setLocalDescription", description.Value())
	if _, err := await(promise); !err.IsNull() {
		panic(errors.New(err.String()))
	}
}

func (c PeerConnection) LocalDescription() SessionDescription {
	return SessionDescription(js.Value(c).Get("localDescription"))
}

func (c PeerConnection) SetRemoteDescription(description SessionDescription) {
	promise := js.Value(c).Call("setRemoteDescription", description.Value())
	if _, err := await(promise); !err.IsNull() {
		panic(errors.New(err.String()))
	}
}

func (c PeerConnection) CreateOffer() SessionDescription {
	promise := js.Value(c).Call("createOffer")
	v, err := await(promise)
	if !err.IsNull() {
		panic(errors.New(err.String()))
	}
	return SessionDescription(v)
}

func (c PeerConnection) CreateAnswer() SessionDescription {
	promise := js.Value(c).Call("createAnswer")
	v, err := await(promise)
	if !err.IsNull() {
		panic(errors.New(err.String()))
	}
	return SessionDescription(v)
}

type ICECandidate js.Value

type PeerConnectionICEEvent js.Value

func (e PeerConnectionICEEvent) Candidate() ICECandidate {
	return ICECandidate(js.Value(e).Get("candidate"))
}

func (c PeerConnection) SetOnICECandidate(onICECandidate func(PeerConnectionICEEvent)) {
	js.Value(c).Set("onicecandidate", js.FuncOf(func(this js.Value, args []js.Value) any {
		onICECandidate(PeerConnectionICEEvent(args[0]))
		return nil
	}))
}

func (c PeerConnection) SignalingState() string {
	return js.Value(c).Get("signalingState").String()
}

type SessionDescription js.Value

func NewSessionDescription(kind string, sdp string) SessionDescription {
	return SessionDescription(js.ValueOf(map[string]any{
		"type": kind,
		"sdp":  sdp,
	}))
}

func (d SessionDescription) Value() js.Value {
	return js.Value(d)
}

func (d SessionDescription) SDP() string {
	return js.Value(d).Get("sdp").String()
}

type DataChannel js.Value

func (c DataChannel) SetOnOpen(onOpen func()) {
	js.Value(c).Set("onopen", js.FuncOf(func(this js.Value, args []js.Value) any {
		onOpen()
		return nil
	}))
}

func (c DataChannel) SetOnMessage(onMessage func()) {
	js.Value(c).Set("onmessage", js.FuncOf(func(this js.Value, args []js.Value) any {
		onMessage()
		return nil
	}))
}

func (c DataChannel) Send(data string) {
	js.Value(c).Call("send", data)
}
