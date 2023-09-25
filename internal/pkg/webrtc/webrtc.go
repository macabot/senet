package webrtc

import (
	"errors"

	"github.com/macabot/hypp"
)

// await awaits a Promise.
// Based on https://stackoverflow.com/a/68427221
func await(js hypp.JavaScript, awaitable hypp.Value) (hypp.Value, hypp.Value) {
	then := make(chan hypp.Value)
	defer close(then)
	thenFunc := js.FuncOf(func(this hypp.Value, args []hypp.Value) interface{} {
		then <- args[0]
		return nil
	})
	defer thenFunc.Release()

	catch := make(chan hypp.Value)
	defer close(catch)
	catchFunc := js.FuncOf(func(this hypp.Value, args []hypp.Value) interface{} {
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

type PeerConnection struct {
	hypp.Value
}

func NewPeerConnection(js hypp.JavaScript, config PeerConnectionConfig) PeerConnection {
	return PeerConnection{js.Global().Get("RTCPeerConnection").New(config.Value())}
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
	return DataChannel{hypp.Value(c).Call("createDataChannel", label, options.Value())}
}

func (c PeerConnection) SetOnICEConnectionStateChange(js hypp.JavaScript, onICEConnectionStateChange func()) {
	hypp.Value(c).Set("oniceconnectionstatechange", js.FuncOf(func(this hypp.Value, args []hypp.Value) any {
		onICEConnectionStateChange()
		return nil
	}))
}

func (c PeerConnection) SetOnConnectionStateChange(js hypp.JavaScript, onConnectionStateChange func()) {
	hypp.Value(c).Set("onconnectionstatechange", js.FuncOf(func(this hypp.Value, args []hypp.Value) any {
		onConnectionStateChange()
		return nil
	}))
}

func (c PeerConnection) ICEConnectionState() string {
	return hypp.Value(c).Get("iceConnectionState").String()
}

func (c PeerConnection) SetLocalDescription(js hypp.JavaScript, description SessionDescription) {
	promise := hypp.Value(c).Call("setLocalDescription", description.Value)
	if _, err := await(js, promise); !err.IsNull() {
		panic(errors.New(err.String()))
	}
}

func (c PeerConnection) LocalDescription() SessionDescription {
	return SessionDescription{hypp.Value(c).Get("localDescription")}
}

func (c PeerConnection) SetRemoteDescription(js hypp.JavaScript, description SessionDescription) {
	promise := hypp.Value(c).Call("setRemoteDescription", description.Value)
	if _, err := await(js, promise); !err.IsNull() {
		panic(errors.New(err.String()))
	}
}

func (c PeerConnection) CreateOffer(js hypp.JavaScript) SessionDescription {
	promise := hypp.Value(c).Call("createOffer")
	v, err := await(js, promise)
	if !err.IsNull() {
		panic(errors.New(err.String()))
	}
	return SessionDescription{v}
}

func (c PeerConnection) CreateAnswer(js hypp.JavaScript) SessionDescription {
	promise := hypp.Value(c).Call("createAnswer")
	v, err := await(js, promise)
	if !err.IsNull() {
		panic(errors.New(err.String()))
	}
	return SessionDescription{v}
}

type ICECandidate hypp.Value

type PeerConnectionICEEvent struct {
	hypp.Value
}

func (e PeerConnectionICEEvent) Candidate() ICECandidate {
	return ICECandidate(hypp.Value(e).Get("candidate"))
}

func (c PeerConnection) SetOnICECandidate(js hypp.JavaScript, onICECandidate func(PeerConnectionICEEvent)) {
	hypp.Value(c).Set("onicecandidate", js.FuncOf(func(this hypp.Value, args []hypp.Value) any {
		onICECandidate(PeerConnectionICEEvent{args[0]})
		return nil
	}))
}

func (c PeerConnection) SignalingState() string {
	return hypp.Value(c).Get("signalingState").String()
}

type SessionDescription struct {
	hypp.Value
}

func NewSessionDescription(js hypp.JavaScript, kind string, sdp string) SessionDescription {
	return SessionDescription{js.ValueOf(map[string]any{
		"type": kind,
		"sdp":  sdp,
	})}
}

func (d SessionDescription) SDP() string {
	return hypp.Value(d).Get("sdp").String()
}

type DataChannel struct {
	hypp.Value
}

func (c DataChannel) SetOnOpen(js hypp.JavaScript, onOpen func()) {
	hypp.Value(c).Set("onopen", js.FuncOf(func(this hypp.Value, args []hypp.Value) any {
		onOpen()
		return nil
	}))
}

func (c DataChannel) SetOnMessage(js hypp.JavaScript, onMessage func()) {
	hypp.Value(c).Set("onmessage", js.FuncOf(func(this hypp.Value, args []hypp.Value) any {
		onMessage()
		return nil
	}))
}

func (c DataChannel) Send(data string) {
	hypp.Value(c).Call("send", data)
}
