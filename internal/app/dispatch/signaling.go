package dispatch

// This file is based on https://stackoverflow.com/a/54985729

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/js"
	"github.com/macabot/hypp/window"
	"github.com/macabot/senet/internal/app/state"
	"github.com/macabot/senet/internal/pkg/webrtc"
)

func initSignaling(s *state.State) {
	state.PeerConnection = webrtc.NewPeerConnection(webrtc.DefaultPeerConnectionConfig)
	state.DataChannel = state.PeerConnection.CreateDataChannel("chat", webrtc.DefaultDataChannelOptions)

	if s.Signaling == nil {
		s.Signaling = &state.Signaling{}
	}
	s.Signaling.Initialized = true
	s.Signaling.ConnectionState = state.PeerConnection.ConnectionState()
	s.Signaling.ICEConnectionState = state.PeerConnection.ICEConnectionState()
	s.Signaling.ReadyState = state.DataChannel.ReadyState()
}

func resetSignaling(s *state.State) {
	state.PeerConnection = webrtc.PeerConnection{}
	state.DataChannel = webrtc.DataChannel{}
	if s.Signaling != nil {
		s.Signaling.Initialized = false
	}
}

func SetSignalingStatesAction(iceConnectionState, connectionState, readyState string) hypp.Action[*state.State] {
	return func(s *state.State, _ hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		if newState.Signaling == nil {
			newState.Signaling = &state.Signaling{}
		}
		newState.Signaling.ICEConnectionState = iceConnectionState
		newState.Signaling.ConnectionState = connectionState
		newState.Signaling.ReadyState = readyState

		return newState
	}
}

func onSignalingStateChange(dispatch hypp.Dispatch) {
	iceConnectionState := state.PeerConnection.ICEConnectionState()
	connectionState := state.PeerConnection.ConnectionState()
	readyState := state.DataChannel.ReadyState()
	dispatch(SetSignalingStatesAction(iceConnectionState, connectionState, readyState), nil)
}

func OnICEConnectionStateChangeSubscriber(dispatch hypp.Dispatch, _ hypp.Payload) hypp.Unsubscribe {
	state.PeerConnection.SetOnICEConnectionStateChange(func() {
		onSignalingStateChange(dispatch)
	})
	return func() {}
}

func OnConnectionStateChangeSubscriber(dispatch hypp.Dispatch, _ hypp.Payload) hypp.Unsubscribe {
	state.PeerConnection.SetOnConnectionStateChange(func() {
		onSignalingStateChange(dispatch)
	})
	return func() {}
}

func OnDataChannelOpenSubscriber(dispatch hypp.Dispatch, _ hypp.Payload) hypp.Unsubscribe {
	state.DataChannel.SetOnOpen(func() {
		onSignalingStateChange(dispatch)
	})
	return func() {}
}

func OnDataChannelMessageSubscriber(dispatch hypp.Dispatch, _ hypp.Payload) hypp.Unsubscribe {
	state.DataChannel.SetOnMessage(func(e js.Value) {
		data := e.Get("data")
		window.Console().Log("DataChannel message event", data)
		message := ParseCommitmentSchemeMessage(data.String())
		switch message.Kind {
		case SendFlipperSecretKind:
			flipperSecret := message.Data.String()
			dispatch(ReceiveFlipperSecretAction(flipperSecret), nil)
		case SendCommitmentKind:
			commitment := message.Data.String()
			dispatch(ReceiveCommitmentAction(commitment), nil)
		case SendFlipperResultsKind:
			flipperResults := parseFlips(message.Data)
			dispatch(ReceiveFlipperResultsAction(flipperResults), nil)
		case SendCallerSecretAndPredictionsKind:
			callerSecretAndPredictions := parseCallerSecretAndPredictions(message.Data)
			dispatch(ReceiveCallerSecretAndPredictionsAction(callerSecretAndPredictions), nil)
		case SendHasThrownKind:
			dispatch(ReceiveHasThrownAction(), nil)
		case SendMoveKind:
			move := parseMove(message.Data)
			dispatch(ReceiveMoveAction(move), nil)
		case SendNoMoveKind:
			dispatch(ReceiveNoMoveAction(), nil)
		default:
			window.Console().Warn("Data message has unknown kind %d", int(message.Kind))
		}
	})
	return func() {}
}

func CreatePeerConnectionOfferEffect() hypp.Effect {
	return hypp.Effect{
		Effecter: func(dispatch hypp.Dispatch, payload hypp.Payload) {
			go func() {
				state.PeerConnection.AwaitSetLocalDescription(state.PeerConnection.AwaitCreateOffer())
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

func CreatePeerConnectionAnswerEffect(offer string) hypp.Effect {
	return hypp.Effect{
		Effecter: func(dispatch hypp.Dispatch, payload hypp.Payload) {
			go func() {
				if state.PeerConnection.SignalingState() != "stable" {
					window.Console().Log(`PeerConnection.SignalingState != "stable"`)
					return
				}
				state.PeerConnection.AwaitSetRemoteDescription(webrtc.NewSessionDescription("offer", offer))
				state.PeerConnection.AwaitSetLocalDescription(state.PeerConnection.AwaitCreateAnswer())
				state.PeerConnection.SetOnICECandidate(func(pci webrtc.PeerConnectionICEEvent) {
					if pci.Candidate().Truthy() {
						return
					}
					answer := state.PeerConnection.LocalDescription().SDP()
					dispatch(setAnswerAction(answer), nil)
				})
			}()
		},
	}
}

func setAnswerAction(answer string) hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		if newState.Signaling == nil {
			newState.Signaling = &state.Signaling{}
		}
		newState.Signaling.Answer = answer
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

func SetSignalingStepJoinGameOfferAction() hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		if newState.Signaling == nil {
			newState.Signaling = &state.Signaling{}
		}
		newState.Signaling.Step = state.SignalingStepJoinGameOffer
		return newState
	}
}

func SetSignalingStepJoinGameAnswerAction() hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		if newState.Signaling == nil {
			newState.Signaling = &state.Signaling{}
		}
		newState.Signaling.Step = state.SignalingStepJoinGameAnswer
		newState.Signaling.Loading = true
		return hypp.StateAndEffects[*state.State]{
			State:   newState,
			Effects: []hypp.Effect{CreatePeerConnectionAnswerEffect(newState.Signaling.Offer)},
		}
	}
}

func SetSignalingOfferAction() hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		event := payload.(window.Event)
		newState := s.Clone()
		if newState.Signaling == nil {
			newState.Signaling = &state.Signaling{}
		}
		newState.Signaling.Offer = event.Target().Value()
		return newState
	}
}

func SetSignalingAnswerAction() hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		event := payload.(window.Event)
		newState := s.Clone()
		if newState.Signaling == nil {
			newState.Signaling = &state.Signaling{}
		}
		newState.Signaling.Answer = event.Target().Value()
		return newState
	}
}

func ConnectNewGameAction() hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		if newState.Signaling == nil {
			newState.Signaling = &state.Signaling{}
		}
		newState.Signaling.Loading = true
		return hypp.StateAndEffects[*state.State]{
			State:   newState,
			Effects: []hypp.Effect{ConnectNewGameEffect(newState.Signaling.Answer)},
		}
	}
}

func ConnectNewGameEffect(answer string) hypp.Effect {
	return hypp.Effect{
		Effecter: func(dispatch hypp.Dispatch, payload hypp.Payload) {
			go func() {
				signalingState := state.PeerConnection.SignalingState()
				if signalingState != "have-local-offer" {
					window.Console().Log("ConnectNewGameEffect expected signaling state 'have-local-offer', got '%s'.", signalingState)
					return
				}
				state.PeerConnection.AwaitSetRemoteDescription(webrtc.NewSessionDescription("answer", answer))
				dispatch(setSignalingLoadingAction(false), nil)
			}()
		},
	}
}

func setSignalingLoadingAction(loading bool) hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		if newState.Signaling == nil {
			newState.Signaling = &state.Signaling{}
		}
		newState.Signaling.Loading = loading
		return newState
	}
}
