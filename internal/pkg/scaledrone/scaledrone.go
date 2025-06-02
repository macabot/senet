// Scaledrone documentation can be found here https://www.scaledrone.com/docs/api-v3-protocol
package scaledrone

import (
	"encoding/json"
	"errors"

	"github.com/macabot/hypp/js"
	"github.com/macabot/hypp/window"
	"github.com/macabot/senet/internal/pkg/websocket"
)

// Set these variables when building using the -ldflags.
var (
	SCALEDRONE_CHANNEL_ID string
)

const (
	scaledroneWebSocketURL = "wss://api.scaledrone.com/v3/websocket"
	callbackHandshake      = 1
	callbackSubscribe      = 2
)

type Scaledrone struct {
	channelID   string
	roomName    string
	clientID    string
	ws          websocket.WebSocket
	isConnected bool
}

func NewScaledrone(channelID string) *Scaledrone {
	return &Scaledrone{
		channelID: channelID,
	}
}

func (s *Scaledrone) Connect(roomName string) {
	s.roomName = roomName
	// Add the "observable-" prefix to ensure we get the observable events:
	// - observable_members
	// - observable_member_join
	// - observable_member_leave
	observableRoomName := "observable-" + roomName
	ws := websocket.NewWebSocket(scaledroneWebSocketURL)
	ws.OnOpen(func() {
		window.Console().Debug("WebSocket opened")
		handshake := Handshake{
			Kind:     "handshake",
			Channel:  s.channelID,
			Callback: callbackHandshake,
		}
		ws.Send(mustJSONMarshal(handshake))
		window.Console().Debug("Sent handshake")
	})

	ws.OnMessage(func(e js.Value) {
		data := e.Get("data").String()
		window.Console().Debug("Received message", data)

		eventData := parseEventData([]byte(data))

		switch data := eventData.(type) {
		case ErrorCallback:
			window.Console().Error(data)
		case HandshakeCallback:
			s.clientID = data.ClientID
			subscribe := Subscribe{
				Kind:     "subscribe",
				Room:     observableRoomName,
				Callback: callbackSubscribe,
			}
			ws.Send(mustJSONMarshal(subscribe))
			window.Console().Debug("Subscribed to room", observableRoomName)
		case SubscribeCallback:
			s.isConnected = true
			window.Console().Debug("Ready to send messages")
		case ObservableMembers:
			window.Console().Debug("Members currently in room", string(mustJSONMarshal(data.Data)))
		case ObservableMemberJoin:
			window.Console().Debug("Member joined room", data.Data.ID)
		case ObservableMemberLeave:
			window.Console().Debug("Member left room", data.Data.ID)
		case PublishReceived:
			window.Console().Debug("Message from %s: %s", data.ClientID, string(mustJSONMarshal(data.Message)))
		default:
			window.Console().Error("Unknown message type", data)
		}
	})

	ws.OnError(func(e js.Value) {
		window.Console().Error("WebSocket error", e)
	})

	ws.OnClose(func(e js.Value) {
		window.Console().Debug("WebSocket closed", e)
	})

	s.ws = ws
}

func parseEventData(b []byte) any {
	discriminator := EventDataDiscriminator{}
	mustJSONUnmarshal(b, &discriminator)

	if discriminator.Callback != 0 && discriminator.Error != "" {
		errorCallback := ErrorCallback{}
		mustJSONUnmarshal(b, &errorCallback)
		return errorCallback
	}

	if discriminator.Callback == callbackHandshake {
		handshakeCallback := HandshakeCallback{}
		mustJSONUnmarshal(b, &handshakeCallback)
		return handshakeCallback
	}

	if discriminator.Callback == callbackSubscribe {
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

func (s Scaledrone) SendMessage(message any) error {
	if !s.isConnected {
		return errors.New("Scaledrone is not connected")
	}

	publishSent := PublishSent{
		Kind:    "publish",
		Room:    s.roomName,
		Message: message,
	}
	s.ws.Send(mustJSONMarshal(publishSent))
	return nil
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

type PublishSent struct {
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

type ObservableMembers struct {
	// Kind must always equal "observable_members".
	Kind string `json:"type"`
	Room string `json:"room"`
	Data []struct {
		ID string `json:"id"`
	} `json:"data"`
}

type ObservableMemberJoin struct {
	// Kind must always equal "observable_member_join".
	Kind string `json:"type"`
	Room string `json:"room"`
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

type ObservableMemberLeave struct {
	// Kind must always equal "observable_member_leave".
	Kind string `json:"type"`
	Room string `json:"room"`
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
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
