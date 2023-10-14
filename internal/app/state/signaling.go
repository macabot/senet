package state

import "github.com/macabot/senet/internal/pkg/webrtc"

type Signaling struct {
	Loading bool // Loading Offer or Answer.
	Offer   string
	Answer  string
}

func (s *Signaling) Clone() *Signaling {
	if s == nil {
		return nil
	}
	return &Signaling{
		Loading: s.Loading,
		Offer:   s.Offer,
		Answer:  s.Answer,
	}
}

var PeerConnection webrtc.PeerConnection
