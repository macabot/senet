package webrtc

import (
	"encoding/json"
	"errors"

	"github.com/macabot/hypp/js"
	"github.com/macabot/senet/internal/pkg/promise"
)

type ICEServer struct {
	URLs       string
	Username   string
	Credential string
}

func (s ICEServer) Value() map[string]any {
	return map[string]any{
		"urls":       s.URLs,
		"username":   s.Username,
		"credential": s.Credential,
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
	js.Value
}

func NewPeerConnection(config PeerConnectionConfig) PeerConnection {
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
	return DataChannel{c.Value.Call("createDataChannel", label, options.Value())}
}

func (c PeerConnection) SetOnICEConnectionStateChange(onICEConnectionStateChange func()) {
	c.Value.Set("oniceconnectionstatechange", js.FuncOf(func(this js.Value, args []js.Value) any {
		onICEConnectionStateChange()
		return nil
	}))
}

type DataChannelEvent struct {
	js.Value
}

func (e DataChannelEvent) Channel() DataChannel {
	return DataChannel{e.Value.Get("channel")}
}

func (c PeerConnection) SetOnDataChannel(onDataChannel func(DataChannelEvent)) {
	c.Value.Set("ondatachannel", js.FuncOf(func(this js.Value, args []js.Value) any {
		onDataChannel(DataChannelEvent{args[0]})
		return nil
	}))
}

func (c PeerConnection) SetOnConnectionStateChange(onConnectionStateChange func()) {
	c.Value.Set("onconnectionstatechange", js.FuncOf(func(this js.Value, args []js.Value) any {
		onConnectionStateChange()
		return nil
	}))
}

func (c PeerConnection) SetOnSignalingStateChange(onSignalingStateChange func()) {
	c.Value.Set("onsignalingstatechange", js.FuncOf(func(this js.Value, args []js.Value) any {
		onSignalingStateChange()
		return nil
	}))
}

func (c PeerConnection) ICEConnectionState() string {
	return c.Value.Get("iceConnectionState").String()
}

func (c PeerConnection) ConnectionState() string {
	return c.Value.Get("connectionState").String()
}

func (c PeerConnection) AwaitSetLocalDescription(description SessionDescription) error {
	p := c.Value.Call("setLocalDescription", description.Value)
	_, err := promise.Await(p)
	return err
}

func (c PeerConnection) LocalDescription() SessionDescription {
	return SessionDescription{c.Value.Get("localDescription")}
}

func (c PeerConnection) AwaitSetRemoteDescription(description SessionDescription) error {
	p := c.Value.Call("setRemoteDescription", description.Value)
	_, err := promise.Await(p)
	return err
}

func (c PeerConnection) AwaitCreateOffer() (SessionDescription, error) {
	p := c.Value.Call("createOffer")
	v, err := promise.Await(p)
	return SessionDescription{v}, err
}

func (c PeerConnection) AwaitCreateAnswer() (SessionDescription, error) {
	p := c.Value.Call("createAnswer")
	v, err := promise.Await(p)
	return SessionDescription{v}, err
}

type ICECandidate struct {
	js.Value
}

// ToJSON converts the ICECandidate to a map that can be JSON serialized.
//
// See https://developer.mozilla.org/en-US/docs/Web/API/RTCIceCandidate/toJSON
func (c ICECandidate) ToJSON() map[string]any {
	m := map[string]any{}
	if candidate := c.Value.Get("candidate"); !candidate.IsUndefined() {
		m["candidate"] = c.Value.Get("candidate").String()
	}
	if sdpMid := c.Value.Get("sdpMid"); !sdpMid.IsUndefined() {
		if sdpMid.IsNull() {
			m["sdpMid"] = nil
		} else {
			m["sdpMid"] = c.Value.Get("sdpMid").String()
		}
	}
	if sdpMLineIndex := c.Value.Get("sdpMLineIndex"); !sdpMLineIndex.IsUndefined() {
		if sdpMLineIndex.IsNull() {
			m["sdpMLineIndex"] = nil
		} else {
			m["sdpMLineIndex"] = c.Value.Get("sdpMLineIndex").Int()
		}
	}
	if userFragment := c.Value.Get("userFragment"); !userFragment.IsUndefined() {
		m["userFragment"] = c.Value.Get("userFragment").String()
	}
	return m
}

func (c ICECandidate) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.ToJSON())
}

type PeerConnectionICEEvent struct {
	js.Value
}

func (e PeerConnectionICEEvent) Candidate() ICECandidate {
	return ICECandidate{e.Value.Get("candidate")}
}

func (c PeerConnection) SetOnICECandidate(onICECandidate func(PeerConnectionICEEvent)) {
	c.Value.Set("onicecandidate", js.FuncOf(func(this js.Value, args []js.Value) any {
		onICECandidate(PeerConnectionICEEvent{args[0]})
		return nil
	}))
}

func (c PeerConnection) SignalingState() string {
	return c.Value.Get("signalingState").String()
}

// Close terminates the RTCPeerConnection's ICE agent, ending any ongoing ICE processing and any active streams.
// See https://developer.mozilla.org/en-US/docs/Web/API/RTCPeerConnection/close
func (c PeerConnection) Close() {
	c.Value.Call("close")
}

type SessionDescription struct {
	js.Value
}

// NewSessionDescription creates a SessionDescription.
// Kind must be "offer" or "answer".
func NewSessionDescription(kind string, sdp string) SessionDescription {
	return SessionDescription{js.ValueOf(map[string]any{
		"type": kind,
		"sdp":  sdp,
	})}
}

func (d SessionDescription) SDP() string {
	return d.Value.Get("sdp").String()
}

type DataChannel struct {
	js.Value
}

func (c DataChannel) SetOnOpen(onOpen func()) {
	c.Value.Set("onopen", js.FuncOf(func(this js.Value, args []js.Value) any {
		onOpen()
		return nil
	}))
}

func (c DataChannel) SetOnClose(onClose func()) {
	c.Value.Set("onclose", js.FuncOf(func(this js.Value, args []js.Value) any {
		onClose()
		return nil
	}))
}

func (c DataChannel) SetOnError(onError func(err error)) {
	c.Value.Set("onerror", js.FuncOf(func(this js.Value, args []js.Value) any {
		onError(errors.New(args[0].String()))
		return nil
	}))
}

func (c DataChannel) SetOnMessage(onMessage func(e js.Value)) {
	c.Value.Set("onmessage", js.FuncOf(func(this js.Value, args []js.Value) any {
		onMessage(args[0])
		return nil
	}))
}

func (c DataChannel) Send(data string) {
	c.Value.Call("send", data)
}

func (c DataChannel) ReadyState() string {
	return c.Value.Get("readyState").String()
}

func (c DataChannel) Label() string {
	return c.Value.Get("label").String()
}

// Close closes the RTCDataChannel. Closure of the data channel is not instantaneous.
// See https://developer.mozilla.org/en-US/docs/Web/API/RTCDataChannel/close
func (c DataChannel) Close() {
	c.Value.Call("close")
}
