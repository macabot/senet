package state

import "github.com/macabot/senet/internal/pkg/webrtc"

type Signaling struct {
	Offer  string
	Answer string
}

func (s *Signaling) Clone() *Signaling {
	if s == nil {
		return nil
	}
	return &Signaling{
		Offer:  s.Offer,
		Answer: s.Answer,
	}
}

var PeerConnection = webrtc.NewPeerConnection(webrtc.DefaultPeerConnectionConfig)
