package dispatch

// This file is based on https://stackoverflow.com/a/54985729

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/js"
	"github.com/macabot/hypp/window"
	"github.com/macabot/senet/internal/app/state"
	"github.com/macabot/senet/internal/pkg/webrtc"
)

func initSignaling() {
	state.PeerConnection = webrtc.NewPeerConnection(webrtc.DefaultPeerConnectionConfig)
	state.PeerConnection.SetOnICEConnectionStateChange(func() {
		window.Console().Log("PeerConnection.ICEConnectionState", state.PeerConnection.ICEConnectionState())
	})
	state.PeerConnection.SetOnConnectionStateChange(func() {
		window.Console().Log("PeerConnection.ConnectionState", state.PeerConnection.ConnectionState())
	})

	state.DataChannel = state.PeerConnection.CreateDataChannel("chat", webrtc.DefaultDataChannelOptions)
	state.DataChannel.SetOnOpen(func() {
		window.Console().Log("DataChannel open event")
	})
	state.DataChannel.SetOnMessage(func(e js.Value) {
		window.Console().Log("DataChannel message event", e.Get("data"))
	})
}

func resetSignaling() {
	state.PeerConnection = webrtc.PeerConnection{}
	state.DataChannel = webrtc.DataChannel{}
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

func SetSignalingStepNewGameOfferAction() hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		if newState.Signaling == nil {
			newState.Signaling = &state.Signaling{}
		}
		newState.Signaling.Step = state.SignalingStepNewGameOffer
		newState.Signaling.Loading = true
		return hypp.StateAndEffects[*state.State]{
			State:   newState,
			Effects: []hypp.Effect{CreatePeerConnectionOfferEffect()},
		}
	}
}

func SetSignalingStepNewGameAnswerAction() hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		if newState.Signaling == nil {
			newState.Signaling = &state.Signaling{}
		}
		newState.Signaling.Step = state.SignalingStepNewGameAnswer
		return newState
	}
}
