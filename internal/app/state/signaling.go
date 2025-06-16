package state

import (
	"encoding/json"
	"fmt"

	"github.com/macabot/senet/internal/pkg/scaledrone"
	"github.com/macabot/senet/internal/pkg/webrtc"
)

// ENUM(
// default
// web-socket-error
// is-connected-to-web-socket
// opponent-is-connected-to-web-socket
// web-rtc-error
// )
//
//go:generate go tool go-enum --marshal
type SignalingStatus int

type SignalingStep int

const (
	SignalingStepDefault SignalingStep = iota
	SignalingStepIsConnectedToWebSocket
	SignalingStepOpponentIsConnectedToWebSocket
	SignalingStepNewGameOffer
	SignalingStepNewGameAnswer
	SignalingStepJoinGameOffer
	SignalingStepJoinGameAnswer
)

func (s SignalingStep) String() string {
	signalingSteps := [...]string{
		"Default",
		"IsConnectedToWebSocket",
		"OpponentIsConnectedToWebSocket",
		"NewGameOffer",
		"NewGameAnswer",
		"JoinGameOffer",
		"JoinGameAnswer",
	}
	return signalingSteps[s]
}

func ToSignalingStep(s string) (SignalingStep, error) {
	var step SignalingStep
	switch s {
	case "Default":
		step = SignalingStepDefault
	case "IsConnectedToWebSocket":
		step = SignalingStepIsConnectedToWebSocket
	case "OpponentIsConnectedToWebSocket":
		step = SignalingStepOpponentIsConnectedToWebSocket
	case "NewGameOfer":
		step = SignalingStepNewGameOffer
	case "NewGameAnswer":
		step = SignalingStepNewGameAnswer
	case "JoinGameOffer":
		step = SignalingStepJoinGameOffer
	case "JoinGameAnswer":
		step = SignalingStepJoinGameAnswer
	default:
		return step, fmt.Errorf("invalid SignalingStep '%s'", s)
	}
	return step, nil
}

func (s SignalingStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *SignalingStep) UnmarshalJSON(data []byte) error {
	var step string
	if err := json.Unmarshal(data, &step); err != nil {
		return err
	}
	var err error
	*s, err = ToSignalingStep(step)
	return err
}

type SignalingError struct {
	Summary string
	Details string
}

type Signaling struct {
	Step SignalingStep
	// Initialized is true when the PeerConnection and DataChannel are set.
	Initialized        bool
	ICEConnectionState string
	ConnectionState    string
	ReadyState         string
	// Loading Offer or Answer.
	Loading bool
	Offer   string
	Answer  string

	Error *SignalingError

	RoomName string
}

func (s *Signaling) Clone() *Signaling {
	if s == nil {
		return nil
	}
	return &Signaling{
		Step:               s.Step,
		Initialized:        s.Initialized,
		ICEConnectionState: s.ICEConnectionState,
		ConnectionState:    s.ConnectionState,
		ReadyState:         s.ReadyState,
		Loading:            s.Loading,
		Offer:              s.Offer,
		Answer:             s.Answer,
		Error:              s.Error,
		RoomName:           s.RoomName,
	}
}

var (
	PeerConnection = webrtc.PeerConnection{}
	DataChannel    = webrtc.DataChannel{}
	Scaledrone     = &scaledrone.Scaledrone{}
)
