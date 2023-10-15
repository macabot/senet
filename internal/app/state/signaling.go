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
	Step    SignalingStep
	Loading bool // Loading Offer or Answer.
	Offer   string
	Answer  string
}

func (s *Signaling) Clone() *Signaling {
	if s == nil {
		return nil
	}
	return &Signaling{
		Step:    s.Step,
		Loading: s.Loading,
		Offer:   s.Offer,
		Answer:  s.Answer,
	}
}

var PeerConnection webrtc.PeerConnection
var DataChannel webrtc.DataChannel
