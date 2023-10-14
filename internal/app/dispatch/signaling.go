package dispatch

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/window"
	"github.com/macabot/senet/internal/app/state"
	"github.com/macabot/senet/internal/pkg/webrtc"
)

func initSignaling() {
	state.PeerConnection = webrtc.NewPeerConnection(webrtc.DefaultPeerConnectionConfig)
	state.PeerConnection.SetOnConnectionStateChange(func() {
		window.Console().Log("PeerConnection.ICEConnectionState", state.PeerConnection.ICEConnectionState())
	})
}

func resetSignaling() {
	state.PeerConnection = webrtc.PeerConnection{}
}

func CreatePeerConnectionOfferEffect() hypp.Effect {
	return hypp.Effect{
		Effecter: func(dispatch hypp.Dispatch, payload hypp.Payload) {
			go func() {
				state.PeerConnection.SetLocalDescription(state.PeerConnection.CreateOffer())
				state.PeerConnection.SetOnICECandidate(func(pci webrtc.PeerConnectionICEEvent) {
					if pci.Candidate().Truthy() {
						return
					}
					offer := state.PeerConnection.LocalDescription().SDP()
					dispatch(setOfferAction(offer), nil)
				})
			}()
		},
	}
}

func setOfferAction(offer string) hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		if newState.Signaling == nil {
			newState.Signaling = &state.Signaling{}
		}
		newState.Signaling.Offer = offer
		newState.Signaling.Loading = false
		return newState
	}
}
