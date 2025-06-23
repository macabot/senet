package state

import (
	"encoding/json"
	"fmt"

	"github.com/macabot/senet/internal/pkg/scaledrone"
	"github.com/macabot/senet/internal/pkg/webrtc"
)

type SignalingStep int

const (
	SignalingStepDefault SignalingStep = iota
	SignalingStepConnectingToWebSocket
	SignalingStepIsConnectedToWebSocket
	SignalingStepOpponentIsConnectedToWebSocket
	SignalingStepNewGameOffer
	SignalingStepNewGameAnswer
	SignalingStepJoinGameOffer
	SignalingStepJoinGameAnswer
	SignalingStepHasWebRTCConnection
)

func (s SignalingStep) String() string {
	signalingSteps := [...]string{
		"SignalingStepDefault",
		"SignalingStepConnectingToWebSocket",
		"SignalingStepIsConnectedToWebSocket",
		"SignalingStepOpponentIsConnectedToWebSocket",
		"SignalingStepNewGameOffer",
		"SignalingStepNewGameAnswer",
		"SignalingStepJoinGameOffer",
		"SignalingStepJoinGameAnswer",
		"SignalingStepHasWebRTCConnection",
	}
	return signalingSteps[s]
}

func ToSignalingStep(s string) (SignalingStep, error) {
	var step SignalingStep
	switch s {
	case "SignalingStepDefault":
		step = SignalingStepDefault
	case "SignalingStepConnectingToWebSocket":
		step = SignalingStepConnectingToWebSocket
	case "SignalingStepIsConnectedToWebSocket":
		step = SignalingStepIsConnectedToWebSocket
	case "SignalingStepOpponentIsConnectedToWebSocket":
		step = SignalingStepOpponentIsConnectedToWebSocket
	case "SignalingStepNewGameOffer":
		step = SignalingStepNewGameOffer
	case "SignalingStepNewGameAnswer":
		step = SignalingStepNewGameAnswer
	case "SignalingStepJoinGameOffer":
		step = SignalingStepJoinGameOffer
	case "SignalingStepJoinGameAnswer":
		step = SignalingStepJoinGameAnswer
	case "SignalingStepHasWebRTCConnection":
		step = SignalingStepHasWebRTCConnection
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
	// Summary must contain a user friendly message.
	Summary string
	// Description must contain detailed information that is useful for debugging by a developer.
	Description string
	Err         JSONSerializableError
}

func NewSignalingError(summary string, description string, err error) *SignalingError {
	return &SignalingError{
		Summary:     summary,
		Description: description,
		Err:         &JSONSerializableErr{Err: err},
	}
}

// Error implements the error interface for SignalingError.
func (se *SignalingError) Error() string {
	return se.Err.Error()
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
