package state

import (
	"github.com/macabot/hypp/js"
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

var PeerConnection webrtc.PeerConnection = webrtc.PeerConnection{Value: js.Undefined()}
var DataChannel webrtc.DataChannel = webrtc.DataChannel{Value: js.Undefined()}
