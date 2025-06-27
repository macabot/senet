package dispatch

// This file is based on https://stackoverflow.com/a/54985729

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/js"
	"github.com/macabot/hypp/window"
	"github.com/macabot/senet/internal/app/state"
	"github.com/macabot/senet/internal/pkg/metered"
	"github.com/macabot/senet/internal/pkg/scaledrone"
	"github.com/macabot/senet/internal/pkg/webrtc"
)

type SendICECandidateMessage struct {
	// Kind must always equal "candidate".
	Kind      string
	Candidate webrtc.ICECandidate
}

func SetRoomName(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	roomName := payload.(string)
	newState := s.Clone()
	newState.RoomName = roomName
	return newState
}

func SetRoomNameByEvent(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	event := payload.(window.Event)
	event.PreventDefault()
	return hypp.ActionAndPayload[*state.State]{
		Action:  SetRoomName,
		Payload: strings.ToUpper(event.Target().Value()),
	}
}

func SetWebSocketConnected(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	newState := s.Clone()
	newState.WebSocketConnected = payload.(bool)
	return newState
}

func SetOpponentWebSocketConnected(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	newState := s.Clone()
	newState.OpponentWebsocketConnected = payload.(bool)
	return newState
}

func CreateRoomEffecter(dispatch hypp.Dispatch, payload hypp.Payload) {
	sd := payload.(*scaledrone.Scaledrone)
	go func() {
		defer RecoverEffectPanic(dispatch)

		roomName := state.RandomRoomName()

		sd.SetOnError(createScaledroneErrorHandler(dispatch))
		sd.SetOnIsConnected(func() {
			dispatch(SetRoomName, roomName)
			dispatch(SetWebSocketConnected, true)
		})
		sd.SetOnMemberJoin(func(memberID string) {
			if len(sd.Members()) != 2 {
				dispatch(SetSignalingError, state.NewSignalingError(
					"Unexpected number of users in the room",
					fmt.Sprintf("%s has joined and there are now %d users in the room.", memberID, len(sd.Members())),
					errors.New("unexpected number of users in room"),
				))
				return
			}

			dispatch(SetOpponentWebSocketConnected, true)
			dispatch(EffectsAction(hypp.Effect{Effecter: CreatePeerConnectionEffecter}), nil)
		})
		sd.SetOnMemberLeave(func(memberID string) {
			if len(sd.Members()) < 2 {
				dispatch(SetOpponentWebSocketConnected, false)
			}
		})

		if err := sd.Connect(roomName); err != nil {
			summary := "Failed to create room"
			description := "Unknown error"
			if errors.Is(err, scaledrone.ErrScaledroneChannelIDNotSet) {
				description = "ScaledroneChannel ID is not set."
			}
			signalingError := state.NewSignalingError(
				summary,
				description,
				err,
			)
			dispatch(SetSignalingError, signalingError)
		}
	}()
}

func JoinRoomEffecter(dispatch hypp.Dispatch, payload hypp.Payload) {
	roomName := payload.(string)
	go func() {
		defer RecoverEffectPanic(dispatch)

		// dispatch(setSignalingStep, state.SignalingStepConnectingToWebSocket)

		state.Scaledrone.SetOnError(createScaledroneErrorHandler(dispatch))
		state.Scaledrone.SetOnIsConnected(func() {
			dispatch(setSignalingStepIsConnectedToWebSocket, nil)
		})
		state.Scaledrone.SetOnObserveMembers(func(members []string) {
			if len(members) >= 2 {
				dispatch(setSignalingStepOpponentIsConnectedToWebsocket, nil)
			}
		})
		state.Scaledrone.SetOnMemberJoin(func(memberID string) {
			if len(state.Scaledrone.Members()) >= 2 {
				dispatch(setSignalingStepOpponentIsConnectedToWebsocket, nil)
			}
		})

		if err := state.Scaledrone.Connect(roomName); errors.Is(err, scaledrone.ErrScaledroneChannelIDNotSet) {
			dispatch(SetSignalingError, state.SignalingError{
				Summary:     "Failed to connect to join room",
				Description: "ScaledroneChannel ID is not set.",
				Err: &state.JSONSerializableErr{
					Err: err,
				},
			})
		} else if err != nil {
			dispatch(SetSignalingError, state.SignalingError{
				Summary:     "Failed to connect to join room",
				Description: "Unknown error",
				Err: &state.JSONSerializableErr{
					Err: err,
				},
			})
		}
	}()
}

func JoinGame(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	event := payload.(window.Event)
	event.PreventDefault()

	roomName := ""
	if s.Signaling != nil {
		roomName = s.Signaling.RoomName
	}

	return hypp.StateAndEffects[*state.State]{
		State:   s,
		Effects: []hypp.Effect{{Effecter: JoinRoomEffecter, Payload: roomName}},
	}
}

func createScaledroneErrorHandler(dispatch hypp.Dispatch) func(err error) {
	return func(err error) {
		if errors.Is(err, scaledrone.ErrUnknownMessageType) {
			// Carry on and hope for the best.
			window.Console().Warn("Ignoring unknown message type error:", err.Error())
			return
		}

		summary := "unknown error"
		if errors.Is(err, scaledrone.ErrCallback) {
			summary = "Failed to create room"
		} else if errors.Is(err, scaledrone.ErrConnection) {
			summary = "Connection lost"
		}
		signalingError := state.NewSignalingError(
			summary,
			"",
			err,
		)
		dispatch(SetSignalingError, signalingError)
	}
}

// func setSignalingStep(s *state.State, payload hypp.Payload) hypp.Dispatchable {
// 	newState := s.Clone()
// 	if newState.Signaling == nil {
// 		newState.Signaling = &state.Signaling{}
// 	}
// 	newState.Signaling.Step = payload.(state.SignalingStep)
// 	return newState
// }

// func setSignalingStepIsConnectedToWebSocket(s *state.State, payload hypp.Payload) hypp.Dispatchable {
// 	newState := s.Clone()
// 	if newState.Signaling == nil {
// 		newState.Signaling = &state.Signaling{}
// 	}
// 	newState.Signaling.Step = state.SignalingStepIsConnectedToWebSocket
// 	if roomName, ok := payload.(string); ok {
// 		newState.Signaling.RoomName = roomName
// 	}
// 	return newState
// }

// func setSignalingStepOpponentIsConnectedToWebsocket(s *state.State, _ hypp.Payload) hypp.Dispatchable {
// 	newState := s.Clone()
// 	if newState.Signaling == nil {
// 		newState.Signaling = &state.Signaling{}
// 	}
// 	newState.Signaling.Step = state.SignalingStepOpponentIsConnectedToWebSocket
// 	// TODO create offer/answer

// 	return hypp.StateAndEffects[*state.State]{
// 		State: newState,
// 		Effects: []hypp.Effect{
// 			{Effecter: createPeerConnectionEffecter},
// 		},
// 	}
// }

func CreatePeerConnectionEffecter(dispatch hypp.Dispatch, payload hypp.Payload) {
	sd := payload.(*scaledrone.Scaledrone)

	peerConnectionConfig := webrtc.DefaultPeerConnectionConfig
	if len(metered.FetchedICEServers) > 0 && metered.FetchErr == nil {
		peerConnectionConfig.ICEServers = metered.FetchedICEServers
	} else {
		window.Console().Error("Could not fetch metered ICE servers. Using default peer connection config.", metered.FetchErr.Error())
	}

	pc := webrtc.NewPeerConnection(peerConnectionConfig)

	pc.SetOnICECandidate(func(event webrtc.PeerConnectionICEEvent) {
		if event.Candidate().Truthy() {
			window.Console().Debug("Sending ICE candidate", event.Candidate())
			sd.SendMessage(SendICECandidateMessage{
				Kind:      "candidate",
				Candidate: event.Candidate(),
			})
		}
	})

	// Handle incoming DataChannel for answering peer.
	pc.SetOnDataChannel(func(dce webrtc.DataChannelEvent) {
		state.DataChannel = dce.Channel()
		window.Console().Debug("Received DataChannel", state.DataChannel.Label())
		setupDataChannelListeners(dispatch)
	})

	pc.SetOnICEConnectionStateChange(func() {
		window.Console().Debug("ICEConnectionState:", pc.ICEConnectionState())
	})

	pc.SetOnSignalingStateChange(func() {
		window.Console().Debug("SignalingState:", pc.SignalingState())
	})

	dispatch(SetPeerConnection, pc)
}

func SetPeerConnection(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	pc := payload.(webrtc.PeerConnection)
	newState := s.Clone()
	newState.PeerConnection = pc
	return newState
}

func setupDataChannelListeners(dispatch hypp.Dispatch) {
	state.DataChannel.SetOnOpen(func() {
		window.Console().Debug("DataChannel opened")
		dispatch(setSignalingStep, state.SignalingStepHasWebRTCConnection)
	})

	state.DataChannel.SetOnMessage(func(e js.Value) {
		data := e.Get("data")
		window.Console().Debug("<<< Receive DataChannel message", data)
	})

	state.DataChannel.SetOnClose(func() {
		window.Console().Error("DataChannel closed")
	})

	state.DataChannel.SetOnError(func(err error) {
		window.Console().Error("DataChannel error", err.Error())
	})
}

// func initSignaling(s *state.State) {
// 	peerConnectionConfig := webrtc.DefaultPeerConnectionConfig
// 	if len(metered.FetchedICEServers) > 0 && metered.FetchErr == nil {
// 		peerConnectionConfig.ICEServers = metered.FetchedICEServers
// 	} else {
// 		window.Console().Error("Could not fetch metered ICE servers. Using default peer connection config.", metered.FetchErr.Error())
// 	}

// 	state.PeerConnection = webrtc.NewPeerConnection(peerConnectionConfig)
// 	state.DataChannel = state.PeerConnection.CreateDataChannel("chat", webrtc.DefaultDataChannelOptions)

// 	if s.Signaling == nil {
// 		s.Signaling = &state.Signaling{}
// 	}
// 	s.Signaling.Initialized = true
// 	s.Signaling.ConnectionState = state.PeerConnection.ConnectionState()
// 	s.Signaling.ICEConnectionState = state.PeerConnection.ICEConnectionState()
// 	s.Signaling.ReadyState = state.DataChannel.ReadyState()
// }

func resetSignaling(s *state.State) {
	if !state.DataChannel.IsUndefined() {
		state.DataChannel.Close()
	}
	if !state.PeerConnection.IsUndefined() {
		state.PeerConnection.Close()
	}
	if state.Scaledrone.IsConnected() {
		state.Scaledrone.Reset()
	}
	state.PeerConnection = webrtc.PeerConnection{}
	state.DataChannel = webrtc.DataChannel{}
	s.Signaling = nil
}

// type SignalingStates struct {
// 	ICEConnectionState string
// 	ConnectionState    string
// 	ReadyState         string
// }

// func SetSignalingStates(s *state.State, payload hypp.Payload) hypp.Dispatchable {
// 	signalingStates := payload.(SignalingStates)
// 	newState := s.Clone()
// 	if newState.Signaling == nil {
// 		newState.Signaling = &state.Signaling{}
// 	}

// 	if newState.Signaling.ICEConnectionState != signalingStates.ICEConnectionState {
// 		window.Console().Debug("ICEConnectionState: %s --> %s", newState.Signaling.ICEConnectionState, signalingStates.ICEConnectionState)
// 	}
// 	if newState.Signaling.ConnectionState != signalingStates.ConnectionState {
// 		window.Console().Debug("ConnectionState: %s --> %s", newState.Signaling.ConnectionState, signalingStates.ConnectionState)
// 	}
// 	if newState.Signaling.ReadyState != signalingStates.ReadyState {
// 		window.Console().Debug("ReadyState: %s --> %s", newState.Signaling.ReadyState, signalingStates.ReadyState)
// 	}

// 	newState.Signaling.ICEConnectionState = signalingStates.ICEConnectionState
// 	newState.Signaling.ConnectionState = signalingStates.ConnectionState
// 	newState.Signaling.ReadyState = signalingStates.ReadyState

// 	return newState
// }

// func onSignalingStateChange(dispatch hypp.Dispatch) {
// 	iceConnectionState := state.PeerConnection.ICEConnectionState()
// 	connectionState := state.PeerConnection.ConnectionState()
// 	readyState := state.DataChannel.ReadyState()
// 	dispatch(SetSignalingStates, SignalingStates{
// 		ICEConnectionState: iceConnectionState,
// 		ConnectionState:    connectionState,
// 		ReadyState:         readyState,
// 	})
// }

// func OnICEConnectionStateChangeSubscriber(dispatch hypp.Dispatch, _ hypp.Payload) hypp.Unsubscribe {
// 	state.PeerConnection.SetOnICEConnectionStateChange(func() {
// 		onSignalingStateChange(dispatch)
// 	})
// 	return func() {}
// }

// func OnConnectionStateChangeSubscriber(dispatch hypp.Dispatch, _ hypp.Payload) hypp.Unsubscribe {
// 	state.PeerConnection.SetOnConnectionStateChange(func() {
// 		onSignalingStateChange(dispatch)
// 	})
// 	return func() {}
// }

// func OnDataChannelOpenSubscriber(dispatch hypp.Dispatch, _ hypp.Payload) hypp.Unsubscribe {
// 	state.DataChannel.SetOnOpen(func() {
// 		onSignalingStateChange(dispatch)
// 	})
// 	return func() {}
// }

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

func SetSignalingError(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	signalingError := payload.(*state.SignalingError)

	var stateDescription string
	if b, err := json.MarshalIndent(s, "", "  "); err == nil {
		stateDescription = "[State]\n" + string(b)
	} else {
		stateDescription = "[Could not JSON encode state]\n" + err.Error()
	}
	if signalingError.Description == "" {
		signalingError.Description = stateDescription
	} else {
		signalingError.Description += "\n" + stateDescription
	}

	newState := s.Clone()
	newState.SignalingError = signalingError
	return newState
}

func CreatePeerConnectionOfferEffect() hypp.Effect {
	return hypp.Effect{
		Effecter: func(dispatch hypp.Dispatch, _ hypp.Payload) {
			go func() {
				defer RecoverEffectPanic(dispatch)

				summary := "Failed to create the offer"

				offer, err := state.PeerConnection.AwaitCreateOffer()
				if err != nil {
					window.RequestAnimationFrame(func() {
						dispatch(SetSignalingError, state.NewSignalingError(summary, "Failed to create peer connection offer", err))
					})
					return
				}
				if err := state.PeerConnection.AwaitSetLocalDescription(offer); err != nil {
					window.RequestAnimationFrame(func() {
						dispatch(SetSignalingError, state.NewSignalingError(summary, "Failed to set peer connection local description", err))
					})
					return
				}
				state.PeerConnection.SetOnICECandidate(func(pci webrtc.PeerConnectionICEEvent) {
					if pci.Candidate().Truthy() {
						return
					}
					offer := state.PeerConnection.LocalDescription().SDP()
					window.Console().Debug("Offer:", offer)
					// window.RequestAnimationFrame(func() {
					// 	dispatch(setOffer, offer)
					// })
				})
			}()
		},
	}
}

// func setOffer(s *state.State, payload hypp.Payload) hypp.Dispatchable {
// 	offer := payload.(string)
// 	newState := s.Clone()
// 	if newState.Signaling == nil {
// 		newState.Signaling = &state.Signaling{}
// 	}
// 	newState.Signaling.Offer = offer
// 	newState.Signaling.Loading = false
// 	return newState
// }

// func CreatePeerConnectionAnswerEffect(offer string) hypp.Effect {
// 	return hypp.Effect{
// 		Effecter: func(dispatch hypp.Dispatch, payload hypp.Payload) {
// 			go func() {
// 				defer RecoverEffectPanic(dispatch)

// 				summary := "Failed to create the answer"

// 				if state.PeerConnection.SignalingState() != "stable" {
// 					window.Console().Log(`PeerConnection.SignalingState != "stable"`)
// 					return
// 				}
// 				offer := webrtc.NewSessionDescription("offer", offer)
// 				if err := state.PeerConnection.AwaitSetRemoteDescription(offer); err != nil {
// 					window.RequestAnimationFrame(func() {
// 						dispatch(setSignalingError, state.NewSignalingError(summary, "Failed to set peer connection remote description", err))
// 					})
// 					return
// 				}
// 				answer, err := state.PeerConnection.AwaitCreateAnswer()
// 				if err != nil {
// 					window.RequestAnimationFrame(func() {
// 						dispatch(setSignalingError, state.NewSignalingError(summary, "Failed to create peer connection answer", err))
// 					})
// 					return
// 				}
// 				if err := state.PeerConnection.AwaitSetLocalDescription(answer); err != nil {
// 					window.RequestAnimationFrame(func() {
// 						dispatch(setSignalingError, state.NewSignalingError(summary, "Failed to set peer connection local description", err))
// 					})
// 					return
// 				}
// 				state.PeerConnection.SetOnICECandidate(func(pci webrtc.PeerConnectionICEEvent) {
// 					if pci.Candidate().Truthy() {
// 						return
// 					}
// 					answer := state.PeerConnection.LocalDescription().SDP()
// 					window.RequestAnimationFrame(func() {
// 						dispatch(setAnswer, answer)
// 					})
// 				})
// 			}()
// 		},
// 	}
// }

// func setAnswer(s *state.State, payload hypp.Payload) hypp.Dispatchable {
// 	answer := payload.(string)
// 	newState := s.Clone()
// 	if newState.Signaling == nil {
// 		newState.Signaling = &state.Signaling{}
// 	}
// 	newState.Signaling.Step = state.SignalingStepJoinGameAnswer
// 	newState.Signaling.Answer = answer
// 	newState.Signaling.Loading = false
// 	return newState
// }

// func SetSignalingStepNewGameOffer(s *state.State, _ hypp.Payload) hypp.Dispatchable {
// 	newState := s.Clone()
// 	if newState.Signaling == nil {
// 		newState.Signaling = &state.Signaling{}
// 	}
// 	newState.Signaling.Step = state.SignalingStepNewGameOffer
// 	newState.Signaling.Loading = true
// 	newState.Signaling.Error = nil
// 	newState.Signaling.RoomName = state.RandomRoomName()
// 	return hypp.StateAndEffects[*state.State]{
// 		State:   newState,
// 		Effects: []hypp.Effect{CreatePeerConnectionOfferEffect()},
// 	}
// }

// func SetSignalingStepNewGameAnswer(s *state.State, _ hypp.Payload) hypp.Dispatchable {
// 	newState := s.Clone()
// 	if newState.Signaling == nil {
// 		newState.Signaling = &state.Signaling{}
// 	}
// 	if s.Signaling.Loading {
// 		return s
// 	}
// 	newState.Signaling.Step = state.SignalingStepNewGameAnswer
// 	return newState
// }

// func SetSignalingStepJoinGameOffer(s *state.State, _ hypp.Payload) hypp.Dispatchable {
// 	newState := s.Clone()
// 	if newState.Signaling == nil {
// 		newState.Signaling = &state.Signaling{}
// 	}
// 	newState.Signaling.Step = state.SignalingStepJoinGameOffer
// 	return newState
// }

// func SetSignalingStepJoinGameAnswer(s *state.State, _ hypp.Payload) hypp.Dispatchable {
// 	newState := s.Clone()
// 	if newState.Signaling == nil {
// 		newState.Signaling = &state.Signaling{}
// 	}
// 	newState.Signaling.Loading = true
// 	newState.Signaling.Error = nil
// 	newState.Signaling.Step = state.SignalingStepJoinGameAnswer
// 	return hypp.StateAndEffects[*state.State]{
// 		State:   newState,
// 		Effects: []hypp.Effect{CreatePeerConnectionAnswerEffect(newState.Signaling.Offer)},
// 	}
// }

// func SetSignalingOffer(s *state.State, payload hypp.Payload) hypp.Dispatchable {
// 	event := payload.(window.Event)
// 	newState := s.Clone()
// 	if newState.Signaling == nil {
// 		newState.Signaling = &state.Signaling{}
// 	}
// 	newState.Signaling.Offer = event.Target().Value()
// 	return newState
// }

// func SetSignalingAnswer(s *state.State, payload hypp.Payload) hypp.Dispatchable {
// 	event := payload.(window.Event)
// 	newState := s.Clone()
// 	if newState.Signaling == nil {
// 		newState.Signaling = &state.Signaling{}
// 	}
// 	newState.Signaling.Answer = event.Target().Value()
// 	return newState
// }

// func ConnectNewGame(s *state.State, _ hypp.Payload) hypp.Dispatchable {
// 	newState := s.Clone()
// 	if newState.Signaling == nil {
// 		newState.Signaling = &state.Signaling{}
// 	}
// 	newState.Signaling.Loading = true
// 	newState.Signaling.Error = nil
// 	return hypp.StateAndEffects[*state.State]{
// 		State:   newState,
// 		Effects: []hypp.Effect{ConnectNewGameEffect(newState.Signaling.Answer)},
// 	}
// }

// func ConnectNewGameEffect(answer string) hypp.Effect {
// 	return hypp.Effect{
// 		Effecter: func(dispatch hypp.Dispatch, payload hypp.Payload) {
// 			go func() {
// 				defer RecoverEffectPanic(dispatch)

// 				signalingState := state.PeerConnection.SignalingState()
// 				if signalingState != "have-local-offer" {
// 					window.Console().Log("ConnectNewGameEffect expected signaling state 'have-local-offer', got '%s'.", signalingState)
// 					return
// 				}
// 				answer := webrtc.NewSessionDescription("answer", answer)
// 				if err := state.PeerConnection.AwaitSetRemoteDescription(answer); err != nil {
// 					window.RequestAnimationFrame(func() {
// 						dispatch(setSignalingError, state.NewSignalingError("Failed to create a new game", "Failed to set peer connection remote description", err))
// 					})
// 					return
// 				}
// 				window.RequestAnimationFrame(func() {
// 					dispatch(setSignalingLoading, false)
// 				})
// 			}()
// 		},
// 	}
// }

// func setSignalingLoading(s *state.State, payload hypp.Payload) hypp.Dispatchable {
// 	loading := payload.(bool)
// 	newState := s.Clone()
// 	if newState.Signaling == nil {
// 		newState.Signaling = &state.Signaling{}
// 	}
// 	newState.Signaling.Loading = loading
// 	newState.Signaling.Error = nil
// 	return newState
// }
