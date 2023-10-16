package state

import "github.com/macabot/senet/internal/pkg/webrtc"

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
	// When true, create the PeerConnection and DataChannel.
	// When false, reset the PeerConnection and DataChannel.
	Initialize         bool
	ICEConnectionState string
	ConnectionState    string
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
		Initialize:         s.Initialize,
		ICEConnectionState: s.ICEConnectionState,
		ConnectionState:    s.ConnectionState,
		Loading:            s.Loading,
		Offer:              s.Offer,
		Answer:             s.Answer,
	}
}

var PeerConnection webrtc.PeerConnection
var DataChannel webrtc.DataChannel
