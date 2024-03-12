package dispatch

// This file is based on https://stackoverflow.com/a/54985729

import (
	"encoding/json"

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
	if !state.DataChannel.IsUndefined() {
		state.DataChannel.Close()
	}
	if !state.PeerConnection.IsUndefined() {
		state.PeerConnection.Close()
	}
	state.PeerConnection = webrtc.PeerConnection{}
	state.DataChannel = webrtc.DataChannel{}
	s.Signaling = nil
}

func SetSignalingStatesAction(iceConnectionState, connectionState, readyState string) hypp.Action[*state.State] {
	return func(s *state.State, _ hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		if newState.Signaling == nil {
			newState.Signaling = &state.Signaling{}
		}

		if newState.Signaling.ICEConnectionState != iceConnectionState {
			window.Console().Debug("ICEConnectionState: %s --> %s", newState.Signaling.ICEConnectionState, iceConnectionState)
		}
		if newState.Signaling.ConnectionState != connectionState {
			window.Console().Debug("ConnectionState: %s --> %s", newState.Signaling.ConnectionState, connectionState)
		}
		if newState.Signaling.ReadyState != readyState {
			window.Console().Debug("ReadyState: %s --> %s", newState.Signaling.ReadyState, readyState)
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
		window.Console().Debug("<<< Receive DataChannel message", data)
		var message CommitmentSchemeMessage[json.RawMessage]
		mustJSONUnmarshal([]byte(data.String()), &message)
		switch message.Kind {
		case SendIsReadyKind:
			dispatch(ReceiveIsReadyAction(), nil)
		case SendFlipperSecretKind:
			var flipperSecret string
			mustJSONUnmarshal(message.Data, &flipperSecret)
			dispatch(ReceiveFlipperSecretAction(flipperSecret), nil)
		case SendCommitmentKind:
			var commitment string
			mustJSONUnmarshal(message.Data, &commitment)
			dispatch(ReceiveCommitmentAction(commitment), nil)
		case SendFlipperResultsKind:
			var flipperResults [4]bool
			mustJSONUnmarshal(message.Data, &flipperResults)
			dispatch(ReceiveFlipperResultsAction(flipperResults), nil)
		case SendCallerSecretAndPredictionsKind:
			var callerSecretAndPredictions CallerSecretAndPredictions
			mustJSONUnmarshal(message.Data, &callerSecretAndPredictions)
			dispatch(ReceiveCallerSecretAndPredictionsAction(callerSecretAndPredictions), nil)
		case SendHasThrownKind:
			dispatch(ReceiveHasThrownAction(), nil)
		case SendMoveKind:
			var move Move
			mustJSONUnmarshal(message.Data, &move)
			dispatch(ReceiveMoveAction(move), nil)
		case SendNoMoveKind:
			dispatch(ReceiveNoMoveAction(), nil)
		default:
			window.Console().Warn("Data message has unknown kind %v", message.Kind)
		}
	})
	return func() {}
}

func sendDataChannelMessage(data string) {
	window.Console().Debug(">>> Send DataChannel message", data)
	state.DataChannel.Send(data)
}

func setSignalingError(err error) hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		if newState.Signaling == nil {
			newState.Signaling = &state.Signaling{}
		}
		newState.Signaling.Error = &state.JSONSerializableErr{Err: err}
		newState.Signaling.Loading = false
		return newState
	}
}

func CreatePeerConnectionOfferEffect() hypp.Effect {
	return hypp.Effect{
		Effecter: func(dispatch hypp.Dispatch, payload hypp.Payload) {
			go func() {
				defer RecoverEffectPanic(dispatch)

				offer, err := state.PeerConnection.AwaitCreateOffer()
				if err != nil {
					window.RequestAnimationFrame(func() {
						dispatch(setSignalingError(err), nil)
					})
					return
				}
				if err := state.PeerConnection.AwaitSetLocalDescription(offer); err != nil {
					window.RequestAnimationFrame(func() {
						dispatch(setSignalingError(err), nil)
					})
					return
				}
				state.PeerConnection.SetOnICECandidate(func(pci webrtc.PeerConnectionICEEvent) {
					if pci.Candidate().Truthy() {
						return
					}
					offer := state.PeerConnection.LocalDescription().SDP()
					window.RequestAnimationFrame(func() {
						dispatch(setOfferAction(offer), nil)
					})
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
				defer RecoverEffectPanic(dispatch)

				if state.PeerConnection.SignalingState() != "stable" {
					window.Console().Log(`PeerConnection.SignalingState != "stable"`)
					return
				}
				offer := webrtc.NewSessionDescription("offer", offer)
				if err := state.PeerConnection.AwaitSetRemoteDescription(offer); err != nil {
					window.RequestAnimationFrame(func() {
						dispatch(setSignalingError(err), nil)
					})
					return
				}
				answer, err := state.PeerConnection.AwaitCreateAnswer()
				if err != nil {
					window.RequestAnimationFrame(func() {
						dispatch(setSignalingError(err), nil)
					})
					return
				}
				if err := state.PeerConnection.AwaitSetLocalDescription(answer); err != nil {
					window.RequestAnimationFrame(func() {
						dispatch(setSignalingError(err), nil)
					})
					return
				}
				state.PeerConnection.SetOnICECandidate(func(pci webrtc.PeerConnectionICEEvent) {
					if pci.Candidate().Truthy() {
						return
					}
					answer := state.PeerConnection.LocalDescription().SDP()
					window.RequestAnimationFrame(func() {
						dispatch(setAnswerAction(answer), nil)
					})
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
		newState.Signaling.Step = state.SignalingStepJoinGameAnswer
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
		newState.Signaling.Error = nil
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
		newState.Signaling.Loading = true
		newState.Signaling.Error = nil
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
		newState.Signaling.Error = nil
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
				defer RecoverEffectPanic(dispatch)

				signalingState := state.PeerConnection.SignalingState()
				if signalingState != "have-local-offer" {
					window.Console().Log("ConnectNewGameEffect expected signaling state 'have-local-offer', got '%s'.", signalingState)
					return
				}
				answer := webrtc.NewSessionDescription("answer", answer)
				if err := state.PeerConnection.AwaitSetRemoteDescription(answer); err != nil {
					window.RequestAnimationFrame(func() {
						dispatch(setSignalingError(err), nil)
					})
					return
				}
				window.RequestAnimationFrame(func() {
					dispatch(setSignalingLoadingAction(false), nil)
				})
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
		newState.Signaling.Error = nil
		return newState
	}
}
