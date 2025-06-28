// Scaledrone documentation can be found here https://www.scaledrone.com/docs/api-v3-protocol
package scaledrone

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/macabot/hypp/js"
	"github.com/macabot/hypp/window"
	"github.com/macabot/senet/internal/pkg/websocket"
)

// Set these variables when building using the -ldflags.
var (
	SCALEDRONE_CHANNEL_ID string
)

const (
	ScaledroneWebSocketURL = "wss://api.scaledrone.com/v3/websocket"
	CallbackHandshake      = 1
	CallbackSubscribe      = 2
)

var (
	ErrScaledroneChannelIDNotSet = errors.New("SCALEDRONE_CHANNEL_ID is not set")

	ErrCallback        = errors.New("Scaledrone callback error")
	ErrHandshakeFailed = fmt.Errorf("%w: Scaledrone handshake failed", ErrCallback)
	ErrSubscribeFailed = fmt.Errorf("%w: Scaledrone subscribe failed", ErrCallback)
	ErrUnknownCallback = fmt.Errorf("%w: unknown Scaledrone callback", ErrCallback)

	ErrUnknownMessageType = errors.New("unknown Scaledrone message type")

	ErrConnection = errors.New("Scaledrone connection error")
)

type Scaledrone struct {
	roomName    string
	clientID    string
	ws          websocket.WebSocket
	isConnected bool
	members     []string

	onIsConnected    func()
	onReceiveMessage func(message any)
	onError          func(err error)
	onObserveMembers func(members []string)
	onMemberJoin     func(memberID string)
	onMemberLeave    func(memberID string)
}

func NewScaledrone() *Scaledrone {
	return &Scaledrone{}
}

func (s Scaledrone) IsConnected() bool {
	return s.isConnected
}

func (s Scaledrone) Members() []string {
	return s.members
}

func (s *Scaledrone) SetOnIsConnected(onIsConnected func()) {
	s.onIsConnected = onIsConnected
}

func (s *Scaledrone) SetOnReceiveMessage(onReceiveMessage func(message any)) {
	s.onReceiveMessage = onReceiveMessage
}

func (s *Scaledrone) SetOnError(onError func(err error)) {
	s.onError = onError
}

func (s *Scaledrone) SetOnObserveMembers(onObserveMembers func(members []string)) {
	s.onObserveMembers = onObserveMembers
}

func (s *Scaledrone) SetOnMemberJoin(onMemberJoin func(memberID string)) {
	s.onMemberJoin = onMemberJoin
}

func (s *Scaledrone) SetOnMemberLeave(onMemberLeave func(memberID string)) {
	s.onMemberLeave = onMemberLeave
}

func (s *Scaledrone) Reset() {
	if !s.ws.IsUndefined() {
		s.ws.Close()
	}
	s.ws = websocket.WebSocket{}

	s.roomName = ""
	s.clientID = ""
	s.isConnected = false
	s.onIsConnected = nil
	s.onReceiveMessage = nil
	s.onError = nil
	s.onObserveMembers = nil
	s.onMemberJoin = nil
	s.onMemberLeave = nil
}

// Connect connects the WebSocket to the Scaledrone server.
//
// Make sure to first set the event listeners before connecting to Scaledrone.
func (s *Scaledrone) Connect(roomName string) error {
	if SCALEDRONE_CHANNEL_ID == "" {
		return ErrScaledroneChannelIDNotSet
	}

	s.roomName = roomName
	// Add the "observable-" prefix to ensure we get the observable events:
	// - observable_members
	// - observable_member_join
	// - observable_member_leave
	observableRoomName := "observable-" + roomName
	ws := websocket.NewWebSocket(ScaledroneWebSocketURL)
	ws.OnOpen(func() {
		window.Console().Debug("WebSocket opened")
		handshake := Handshake{
			Kind:     "handshake",
			Channel:  SCALEDRONE_CHANNEL_ID,
			Callback: CallbackHandshake,
		}
		ws.Send(mustJSONMarshal(handshake))
		window.Console().Debug("Sent handshake")
	})

	ws.OnMessage(func(e js.Value) {
		rawData := e.Get("data").String()
		window.Console().Debug("Received message", rawData)

		eventData := parseEventData([]byte(rawData))

		switch data := eventData.(type) {
		case ErrorCallback:
			var err error
			if data.Callback == CallbackHandshake {
				err = fmt.Errorf("%w: %s", ErrHandshakeFailed, rawData)
			} else if data.Callback == CallbackSubscribe {
				err = fmt.Errorf("%w: %s", ErrSubscribeFailed, rawData)
			} else {
				err = fmt.Errorf("%w: %s", ErrUnknownCallback, rawData)
			}
			window.Console().Error(err.Error())
			if s.onError != nil {
				s.onError(err)
			}
		case HandshakeCallback:
			s.clientID = data.ClientID
			subscribe := Subscribe{
				Kind:     "subscribe",
				Room:     observableRoomName,
				Callback: CallbackSubscribe,
			}
			ws.Send(mustJSONMarshal(subscribe))
			window.Console().Debug("Subscribed to room", observableRoomName)
		case SubscribeCallback:
			s.isConnected = true
			window.Console().Debug("Ready to send messages")
			if s.onIsConnected != nil {
				s.onIsConnected()
			}
		case ObservableMembers:
			window.Console().Debug("Members currently in room", string(mustJSONMarshal(data.Data)))
			members := make([]string, len(data.Data))
			for i, member := range data.Data {
				members[i] = member.ID
			}
			s.members = members
			if s.onObserveMembers != nil {
				s.onObserveMembers(members)
			}
		case ObservableMemberJoin:
			window.Console().Debug("Member joined room", data.Data.ID)
			s.members = append(s.members, data.Data.ID)
			if s.onMemberJoin != nil {
				s.onMemberJoin(data.Data.ID)
			}
		case ObservableMemberLeave:
			window.Console().Debug("Member left room", data.Data.ID)
			s.removeMember(data.Data.ID)
			if s.onMemberLeave != nil {
				s.onMemberLeave(data.Data.ID)
			}
		case PublishReceived:
			window.Console().Debug("Message from %s: %s", data.ClientID, string(mustJSONMarshal(data.Message)))
			if s.onReceiveMessage != nil && data.ClientID != s.clientID {
				s.onReceiveMessage(data.Message)
			}
		default:
			window.Console().Error("Unknown message type", rawData)
			if s.onError != nil {
				s.onError(fmt.Errorf("%w: %s", ErrUnknownMessageType, rawData))
			}
		}
	})

	ws.OnError(func(err error) {
		err = fmt.Errorf("%w: %w", ErrConnection, err)
		window.Console().Error(err.Error())
		if s.onError != nil {
			s.onError(err)
		}
	})

	ws.OnClose(func(e js.Value) {
		window.Console().Debug("WebSocket closed", e)
	})

	s.ws = ws
	return nil
}

func (s *Scaledrone) removeMember(memberID string) {
	for i, member := range s.members {
		if member == memberID {
			s.members = append(s.members[:i], s.members[i+1:]...)
			return
		}
	}
}

func parseEventData(b []byte) any {
	discriminator := EventDataDiscriminator{}
	mustJSONUnmarshal(b, &discriminator)

	if discriminator.Callback != 0 && discriminator.Error != "" {
		errorCallback := ErrorCallback{}
		mustJSONUnmarshal(b, &errorCallback)
		return errorCallback
	}

	if discriminator.Callback == CallbackHandshake {
		handshakeCallback := HandshakeCallback{}
		mustJSONUnmarshal(b, &handshakeCallback)
		return handshakeCallback
	}

	if discriminator.Callback == CallbackSubscribe {
		subscribeCallback := SubscribeCallback{}
		mustJSONUnmarshal(b, &subscribeCallback)
		return subscribeCallback
	}

	if discriminator.Kind == "publish" {
		publishReceived := PublishReceived{}
		mustJSONUnmarshal(b, &publishReceived)
		return publishReceived
	}

	if discriminator.Kind == "observable_members" {
		observableMembers := ObservableMembers{}
		mustJSONUnmarshal(b, &observableMembers)
		return observableMembers
	}

	if discriminator.Kind == "observable_member_join" {
		observableMemberJoin := ObservableMemberJoin{}
		mustJSONUnmarshal(b, &observableMemberJoin)
		return observableMemberJoin
	}

	if discriminator.Kind == "observable_member_leave" {
		observableMemberLeave := ObservableMemberLeave{}
		mustJSONUnmarshal(b, &observableMemberLeave)
		return observableMemberLeave
	}

	return nil
}

// SendMessage sends a message to the connected room.
//
// It must be possible to JSON encode the message. If not, the function will panic.
func (s Scaledrone) SendMessage(message any) {
	publishSent := PublishSend{
		Kind:    "publish",
		Room:    s.roomName,
		Message: message,
	}
	s.ws.Send(mustJSONMarshal(publishSent))
}

type Handshake struct {
	// Kind must always equal "handshake".
	Kind     string `json:"type"`
	Channel  string `json:"channel"`
	Callback int    `json:"callback"`
}

type Subscribe struct {
	// Kind must always equal "subscribe".
	Kind     string `json:"type"`
	Room     string `json:"room"`
	Callback int    `json:"callback"`
}

type PublishSend struct {
	// Kind must always equal "publish".
	Kind string `json:"type"`
	Room string `json:"room"`
	// Message can be one of the following types:
	// - JSON
	// - string
	// - number
	Message any `json:"message"`
}

type EventDataDiscriminator struct {
	Kind     string `json:"type"`
	Callback int    `json:"callback"`
	Error    string `json:"error"`
}

// ErrorCallback can be returned on
// - a failed handshake
// - a failed subscribe
type ErrorCallback struct {
	Callback int    `json:"callback"`
	Error    string `json:"error"`
}

type HandshakeCallback struct {
	Callback int    `json:"callback"`
	ClientID string `json:"client_id"`
}

type SubscribeCallback struct {
	Callback int `json:"callback"`
}

type PublishReceived struct {
	// Kind must always equal "publish".
	Kind string `json:"type"`
	Room string `json:"room"`
	// Message can be one of the following types:
	// - JSON
	// - string
	// - number
	Message  any    `json:"message"`
	ClientID string `json:"client_id"`
}

type Member struct {
	ID string `json:"id"`
}

type ObservableMembers struct {
	// Kind must always equal "observable_members".
	Kind string   `json:"type"`
	Room string   `json:"room"`
	Data []Member `json:"data"`
}

type ObservableMemberJoin struct {
	// Kind must always equal "observable_member_join".
	Kind string `json:"type"`
	Room string `json:"room"`
	Data Member `json:"data"`
}

type ObservableMemberLeave struct {
	// Kind must always equal "observable_member_leave".
	Kind string `json:"type"`
	Room string `json:"room"`
	Data Member `json:"data"`
}

func mustJSONMarshal(v any) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}

func mustJSONUnmarshal(b []byte, v any) {
	if err := json.Unmarshal(b, v); err != nil {
		panic(err)
	}
}
