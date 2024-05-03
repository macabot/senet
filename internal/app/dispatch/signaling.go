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

type SignalingStates struct {
	ICEConnectionState string
	ConnectionState    string
	ReadyState         string
}

func SetSignalingStates(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	signalingStates := payload.(SignalingStates)
	newState := s.Clone()
	if newState.Signaling == nil {
		newState.Signaling = &state.Signaling{}
	}

	if newState.Signaling.ICEConnectionState != signalingStates.ICEConnectionState {
		window.Console().Debug("ICEConnectionState: %s --> %s", newState.Signaling.ICEConnectionState, signalingStates.ICEConnectionState)
	}
	if newState.Signaling.ConnectionState != signalingStates.ConnectionState {
		window.Console().Debug("ConnectionState: %s --> %s", newState.Signaling.ConnectionState, signalingStates.ConnectionState)
	}
	if newState.Signaling.ReadyState != signalingStates.ReadyState {
		window.Console().Debug("ReadyState: %s --> %s", newState.Signaling.ReadyState, signalingStates.ReadyState)
	}

	newState.Signaling.ICEConnectionState = signalingStates.ICEConnectionState
	newState.Signaling.ConnectionState = signalingStates.ConnectionState
	newState.Signaling.ReadyState = signalingStates.ReadyState

	return newState
}

func onSignalingStateChange(dispatch hypp.Dispatch) {
	iceConnectionState := state.PeerConnection.ICEConnectionState()
	connectionState := state.PeerConnection.ConnectionState()
	readyState := state.DataChannel.ReadyState()
	dispatch(SetSignalingStates, SignalingStates{
		ICEConnectionState: iceConnectionState,
		ConnectionState:    connectionState,
		ReadyState:         readyState,
	})
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
			dispatch(ReceiveIsReady, nil)
		case SendFlipperSecretKind:
			var flipperSecret string
			mustJSONUnmarshal(message.Data, &flipperSecret)
			dispatch(ReceiveFlipperSecret, flipperSecret)
		case SendCommitmentKind:
			var commitment string
			mustJSONUnmarshal(message.Data, &commitment)
			dispatch(ReceiveCommitment, commitment)
		case SendFlipperResultsKind:
			var flipperResults [4]bool
			mustJSONUnmarshal(message.Data, &flipperResults)
			dispatch(ReceiveFlipperResults, flipperResults)
		case SendCallerSecretAndPredictionsKind:
			var callerSecretAndPredictions CallerSecretAndPredictions
			mustJSONUnmarshal(message.Data, &callerSecretAndPredictions)
			dispatch(ReceiveCallerSecretAndPredictions, callerSecretAndPredictions)
		case SendHasThrownKind:
			dispatch(ReceiveHasThrown, nil)
		case SendMoveKind:
			var move Move
			mustJSONUnmarshal(message.Data, &move)
			dispatch(ReceiveMove, move)
		case SendNoMoveKind:
			dispatch(ReceiveNoMove, nil)
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

func setSignalingError(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	err := payload.(error)
	newState := s.Clone()
	if newState.Signaling == nil {
		newState.Signaling = &state.Signaling{}
	}
	newState.Signaling.Error = &state.JSONSerializableErr{Err: err}
	newState.Signaling.Loading = false
	return newState
}

func CreatePeerConnectionOfferEffect() hypp.Effect {
	return hypp.Effect{
		Effecter: func(dispatch hypp.Dispatch, payload hypp.Payload) {
			go func() {
				defer RecoverEffectPanic(dispatch)

				offer, err := state.PeerConnection.AwaitCreateOffer()
				if err != nil {
					window.RequestAnimationFrame(func() {
						dispatch(setSignalingError, err)
					})
					return
				}
				if err := state.PeerConnection.AwaitSetLocalDescription(offer); err != nil {
					window.RequestAnimationFrame(func() {
						dispatch(setSignalingError, err)
					})
					return
				}
				state.PeerConnection.SetOnICECandidate(func(pci webrtc.PeerConnectionICEEvent) {
					if pci.Candidate().Truthy() {
						return
					}
					offer := state.PeerConnection.LocalDescription().SDP()
					window.RequestAnimationFrame(func() {
						dispatch(setOffer, offer)
					})
				})
			}()
		},
	}
}

func setOffer(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	offer := payload.(string)
	newState := s.Clone()
	if newState.Signaling == nil {
		newState.Signaling = &state.Signaling{}
	}
	newState.Signaling.Offer = offer
	newState.Signaling.Loading = false
	return newState
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
						dispatch(setSignalingError, err)
					})
					return
				}
				answer, err := state.PeerConnection.AwaitCreateAnswer()
				if err != nil {
					window.RequestAnimationFrame(func() {
						dispatch(setSignalingError, err)
					})
					return
				}
				if err := state.PeerConnection.AwaitSetLocalDescription(answer); err != nil {
					window.RequestAnimationFrame(func() {
						dispatch(setSignalingError, err)
					})
					return
				}
				state.PeerConnection.SetOnICECandidate(func(pci webrtc.PeerConnectionICEEvent) {
					if pci.Candidate().Truthy() {
						return
					}
					answer := state.PeerConnection.LocalDescription().SDP()
					window.RequestAnimationFrame(func() {
						dispatch(setAnswer, answer)
					})
				})
			}()
		},
	}
}

func setAnswer(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	answer := payload.(string)
	newState := s.Clone()
	if newState.Signaling == nil {
		newState.Signaling = &state.Signaling{}
	}
	newState.Signaling.Step = state.SignalingStepJoinGameAnswer
	newState.Signaling.Answer = answer
	newState.Signaling.Loading = false
	return newState
}

func SetSignalingStepNewGameOffer(s *state.State, _ hypp.Payload) hypp.Dispatchable {
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

func SetSignalingStepNewGameAnswer(s *state.State, _ hypp.Payload) hypp.Dispatchable {
	newState := s.Clone()
	if newState.Signaling == nil {
		newState.Signaling = &state.Signaling{}
	}
	if s.Signaling.Loading {
		return s
	}
	newState.Signaling.Step = state.SignalingStepNewGameAnswer
	return newState
}

func SetSignalingStepJoinGameOffer(s *state.State, _ hypp.Payload) hypp.Dispatchable {
	newState := s.Clone()
	if newState.Signaling == nil {
		newState.Signaling = &state.Signaling{}
	}
	newState.Signaling.Step = state.SignalingStepJoinGameOffer
	return newState
}

func SetSignalingStepJoinGameAnswer(s *state.State, _ hypp.Payload) hypp.Dispatchable {
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

func SetSignalingOffer(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	event := payload.(window.Event)
	newState := s.Clone()
	if newState.Signaling == nil {
		newState.Signaling = &state.Signaling{}
	}
	newState.Signaling.Offer = event.Target().Value()
	return newState
}

func SetSignalingAnswer(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	event := payload.(window.Event)
	newState := s.Clone()
	if newState.Signaling == nil {
		newState.Signaling = &state.Signaling{}
	}
	newState.Signaling.Answer = event.Target().Value()
	return newState
}

func ConnectNewGame(s *state.State, _ hypp.Payload) hypp.Dispatchable {
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
						dispatch(setSignalingError, err)
					})
					return
				}
				window.RequestAnimationFrame(func() {
					dispatch(setSignalingLoading, false)
				})
			}()
		},
	}
}

func setSignalingLoading(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	loading := payload.(bool)
	newState := s.Clone()
	if newState.Signaling == nil {
		newState.Signaling = &state.Signaling{}
	}
	newState.Signaling.Loading = loading
	newState.Signaling.Error = nil
	return newState
}
