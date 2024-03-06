package state

import (
	"encoding/json"
	"fmt"

	"github.com/macabot/senet/internal/pkg/webrtc"
)

type SignalingStep int

const (
	SignalingStepDefault SignalingStep = iota
	SignalingStepNewGameOffer
	SignalingStepNewGameAnswer
	SignalingStepJoinGameOffer
	SignalingStepJoinGameAnswer
)

func (s SignalingStep) String() string {
	signalingSteps := [...]string{"Default", "NewGameOffer", "NewGameAnswer", "JoinGameOffer", "JoinGameAnswer"}
	return signalingSteps[s]
}

func ToSignalingStep(s string) (SignalingStep, error) {
	var step SignalingStep
	switch s {
	case "Default":
		step = SignalingStepDefault
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
	var err error
	*s, err = ToSignalingStep(string(data))
	return err
}

type Signaling struct {
	Step SignalingStep
	// When true, the PeerConnection and DataChannel are set.
	Initialized        bool
	ICEConnectionState string
	ConnectionState    string
	ReadyState         string
	// Loading Offer or Answer.
	Loading bool
	Offer   string
	Answer  string
	Error   error
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
	}
}

var PeerConnection webrtc.PeerConnection = webrtc.PeerConnection{}
var DataChannel webrtc.DataChannel = webrtc.DataChannel{}
