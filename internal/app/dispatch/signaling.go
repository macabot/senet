package dispatch

// This file is based on https://stackoverflow.com/a/54985729

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/window"
	"github.com/macabot/senet/internal/app/state"
	"github.com/macabot/senet/internal/pkg/metered"
	"github.com/macabot/senet/internal/pkg/scaledrone"
	"github.com/macabot/senet/internal/pkg/webrtc"
)

type ICECandidateMessage struct {
	// Kind must always equal "candidate".
	Kind      string
	Candidate webrtc.ICECandidate
}

type OfferMessage struct {
	// Kind must always equal "offer".
	Kind  string
	Offer webrtc.SessionDescription
}

type AnswerMessage struct {
	// Kind must always equal "answer".
	Kind   string
	Answer webrtc.SessionDescription
}

type MessageDiscriminator struct {
	Kind string
}

func parseReceivedMessage(raw json.RawMessage) any {
	discriminator := MessageDiscriminator{}
	mustJSONUnmarshal(raw, &discriminator)
	switch discriminator.Kind {
	case "candidate":
		candidate := webrtc.ICECandidate{}
		mustJSONUnmarshal(raw, &candidate)
		return candidate
	case "offer":
		offer := webrtc.SessionDescription{}
		mustJSONUnmarshal(raw, &offer)
		return offer
	case "answer":
		answer := webrtc.SessionDescription{}
		mustJSONUnmarshal(raw, &answer)
		return answer
	default:
		return nil
	}
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
	if newState.WebSocketConnected {
		newState.ConnectingToWebSocket = false
	}
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
		sd.SetOnReceiveMessage(func(message json.RawMessage) {
			dispatch(OnWebSocketMessage, message)
		})
		sd.SetOnObserveMembers(func(members []string) {
			if len(members) != 1 {
				dispatch(AddSignalingError, state.NewSignalingError(
					"Unexpected number of users in the room",
					fmt.Sprintf("After creating the room there are %d users in the room insetead of 1: %s.", len(sd.Members()), strings.Join(members, ", ")),
					errors.New("unexpected number of users in room"),
				))
			}
		})
		sd.SetOnMemberJoin(func(memberID string) {
			if len(sd.Members()) != 2 {
				dispatch(AddSignalingError, state.NewSignalingError(
					"Unexpected number of users in the room",
					fmt.Sprintf("%s has joined and there are now %d users in the room.", memberID, len(sd.Members())),
					errors.New("unexpected number of users in room"),
				))
				return
			}

			dispatch(SetOpponentWebSocketConnected, true)
			dispatch(CreatePeerConnection, nil)
		})
		sd.SetOnMemberLeave(func(memberID string) {
			if len(sd.Members()) < 2 {
				dispatch(SetOpponentWebSocketConnected, false)
			}
		})

		dispatch(SetConnectingToWebSocket, true)
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
			dispatch(AddSignalingError, signalingError)
		}
	}()
}

func SetConnectingToWebSocket(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	newState := s.Clone()
	newState.ConnectingToWebSocket = payload.(bool)
	return newState
}

func OnWebSocketMessage(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	message := payload.(json.RawMessage)
	if s.PeerConnection.IsUndefined() {
		return hypp.StateAndEffects[*state.State]{
			State: s,
			Effects: []hypp.Effect{
				{
					Effecter: CreatePeerConnectionEffecter,
					Payload:  s.Scaledrone,
				},
				{
					Effecter: HandleIncomingMessageAfterPcReadyEffecter,
					Payload:  message,
				},
			},
		}
	} else {
		return hypp.ActionAndPayload[*state.State]{
			Action:  HandleIncomingMessageAfterPcReady,
			Payload: message,
		}
	}
}

func HandleIncomingMessageAfterPcReadyEffecter(dispatch hypp.Dispatch, payload hypp.Payload) {
	message := payload.(json.RawMessage)
	dispatch(HandleIncomingMessageAfterPcReady, message)
}

func HandleIncomingMessageAfterPcReady(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	if s.PeerConnection.IsUndefined() {
		panic("HandleIncomingMessageAfterPcReady called without a valid peer connection.")
	}

	message := payload.(json.RawMessage)
	parsedMessage := parseReceivedMessage(message)
	switch m := parsedMessage.(type) {
	case OfferMessage:
		return hypp.StateAndEffects[*state.State]{
			State: s,
			Effects: []hypp.Effect{
				{
					Effecter: CreateAnswerEffecter,
					Payload: CreateAnswerEffecterPayload{
						PeerConnection: s.PeerConnection,
						Scaledrone:     s.Scaledrone,
						Message:        m,
					},
				},
			},
		}
	case AnswerMessage:
		return hypp.StateAndEffects[*state.State]{
			State: s,
			Effects: []hypp.Effect{
				{
					Effecter: HandleAnswerMessageEffecter,
					Payload: HandleAnswerMessageEffecerPayload{
						PeerConnection: s.PeerConnection,
						Message:        m,
					},
				},
			},
		}
	case ICECandidateMessage:
		return hypp.StateAndEffects[*state.State]{
			State: s,
			Effects: []hypp.Effect{
				{
					Effecter: AddICECandidateEffecter,
					Payload: AddICECandidateEffecterPayload{
						PeerConnection: s.PeerConnection,
						Candidate:      m.Candidate,
					},
				},
			},
		}
	default:
		return s
	}
}

type CreateAnswerEffecterPayload struct {
	PeerConnection webrtc.PeerConnection
	Scaledrone     *scaledrone.Scaledrone
	Message        OfferMessage
}

func CreateAnswerEffecter(dispatch hypp.Dispatch, payload hypp.Payload) {
	p := payload.(CreateAnswerEffecterPayload)
	go func() {
		defer RecoverEffectPanic(dispatch)

		if err := p.PeerConnection.AwaitSetRemoteDescription(p.Message.Offer); err != nil {
			dispatch(AddSignalingError, state.NewSignalingError(
				"Failed to connect to opponent",
				"Failed to set peer connection remote description",
				err,
			))
			return
		}
		answer, err := p.PeerConnection.AwaitCreateAnswer()
		if err != nil {
			dispatch(AddSignalingError, state.NewSignalingError(
				"Failed to connect to opponent",
				"Failed to create peer connection answer",
				err,
			))
			return
		}
		if err := p.PeerConnection.AwaitSetLocalDescription(answer); err != nil {
			dispatch(AddSignalingError, state.NewSignalingError(
				"Failed to connect to opponent",
				"Failed to set peer connection local description",
				err,
			))
			return
		}
		p.Scaledrone.SendMessage(AnswerMessage{
			Kind:   "answer",
			Answer: answer,
		})
	}()
}

type HandleAnswerMessageEffecerPayload struct {
	PeerConnection webrtc.PeerConnection
	Message        AnswerMessage
}

func HandleAnswerMessageEffecter(dispatch hypp.Dispatch, payload hypp.Payload) {
	p := payload.(HandleAnswerMessageEffecerPayload)
	go func() {
		defer RecoverEffectPanic(dispatch)

		if err := p.PeerConnection.AwaitSetRemoteDescription(p.Message.Answer); err != nil {
			dispatch(AddSignalingError, state.NewSignalingError(
				"Failed to connect to opponent",
				"Failed to set peer connection remote description",
				err,
			))
			return
		}
	}()
}

type AddICECandidateEffecterPayload struct {
	PeerConnection webrtc.PeerConnection
	Candidate      webrtc.ICECandidate
}

func AddICECandidateEffecter(dispatch hypp.Dispatch, payload hypp.Payload) {
	p := payload.(AddICECandidateEffecterPayload)
	go func() {
		defer RecoverEffectPanic(dispatch)

		if err := p.PeerConnection.AwaitAddICECandidate(p.Candidate); err != nil {
			dispatch(AddSignalingError, state.NewSignalingError(
				"Failed to connect to opponent",
				"Failed to add ICE candidate",
				err,
			))
		}
	}()
}

type HandleIncomingMessagePayload struct {
	Message        json.RawMessage
	PeerConnection webrtc.PeerConnection
}

type ScaledroneAndRoomName struct {
	Scaledrone *scaledrone.Scaledrone
	RoomName   string
}

func JoinRoomEffecter(dispatch hypp.Dispatch, payload hypp.Payload) {
	scaledroneAndRoomName := payload.(ScaledroneAndRoomName)
	sd := scaledroneAndRoomName.Scaledrone
	roomName := scaledroneAndRoomName.RoomName
	go func() {
		defer RecoverEffectPanic(dispatch)

		sd.SetOnError(createScaledroneErrorHandler(dispatch))
		sd.SetOnIsConnected(func() {
			dispatch(SetWebSocketConnected, true)
		})
		sd.SetOnObserveMembers(func(members []string) {
			if len(members) > 2 {
				dispatch(AddSignalingError, state.NewSignalingError(
					"Unexpected number of users in the room",
					fmt.Sprintf("After joining the room there are %d users in the room instead of 2.: %s", len(sd.Members()), strings.Join(members, ", ")),
					errors.New("unexpected number of users in room"),
				))
				return
			}

			dispatch(SetOpponentWebSocketConnected, len(members) == 2)
			dispatch(CreatePeerConnection, nil)
		})
		sd.SetOnMemberJoin(func(memberID string) {
			if len(sd.Members()) != 2 {
				dispatch(AddSignalingError, state.NewSignalingError(
					"Unexpected number of users in the room",
					fmt.Sprintf("%s has joined and there are now %d users in the room.", memberID, len(sd.Members())),
					errors.New("unexpected number of users in room"),
				))
			}
		})
		sd.SetOnMemberLeave(func(memberID string) {
			if len(sd.Members()) < 2 {
				dispatch(SetOpponentWebSocketConnected, false)
			}
		})

		dispatch(SetConnectingToWebSocket, true)
		if err := sd.Connect(roomName); err != nil {
			summary := "Failed to join room"
			description := "Unknown error"
			if errors.Is(err, scaledrone.ErrScaledroneChannelIDNotSet) {
				description = "ScaledroneChannel ID is not set."
			}
			signalingError := state.NewSignalingError(
				summary,
				description,
				err,
			)
			dispatch(AddSignalingError, signalingError)
		}
	}()
}

func JoinGame(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	event := payload.(window.Event)
	event.PreventDefault()

	return hypp.StateAndEffects[*state.State]{
		State: s,
		Effects: []hypp.Effect{{
			Effecter: JoinRoomEffecter,
			Payload: ScaledroneAndRoomName{
				Scaledrone: s.Scaledrone,
				RoomName:   s.RoomName,
			},
		}},
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
		dispatch(AddSignalingError, signalingError)
	}
}

func CreatePeerConnection(s *state.State, _ hypp.Payload) hypp.Dispatchable {
	return hypp.StateAndEffects[*state.State]{
		State: s,
		Effects: []hypp.Effect{
			{
				Effecter: CreatePeerConnectionEffecter,
				Payload:  s.Scaledrone,
			},
		},
	}
}

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
			if !sd.IsConnected() {
				dispatch(AddSignalingError, state.NewSignalingError(
					"Failed to connect to opponent",
					"Cannot send ICE candidate when WebSocket isn't connected.",
					errors.New("websocket not connected"),
				))
				return
			}

			window.Console().Debug("Sending ICE candidate", event.Candidate())
			sd.SendMessage(ICECandidateMessage{
				Kind:      "candidate",
				Candidate: event.Candidate(),
			})
		}
	})

	// Handle incoming DataChannel for answering peer.
	pc.SetOnDataChannel(func(event webrtc.DataChannelEvent) {
		dc := event.Channel()
		window.Console().Debug("Received DataChannel", dc.Label())
		dispatch(setupDataChannelListeners, dc)
		dispatch(SetDataChannel, dc)
	})

	pc.SetOnICEConnectionStateChange(func() {
		window.Console().Debug("ICEConnectionState:", pc.ICEConnectionState())
	})

	pc.SetOnSignalingStateChange(func() {
		window.Console().Debug("SignalingState:", pc.SignalingState())
	})

	dispatch(SetPeerConnection, pc)
	// TODO Offer must only be create by user that created the game
	dispatch(CreateOfferAndDataChannel, nil)
}

func CreateOfferAndDataChannel(s *state.State, _ hypp.Payload) hypp.Dispatchable {
	return hypp.StateAndEffects[*state.State]{
		State: s,
		Effects: []hypp.Effect{
			{
				Effecter: CreateOfferAndDataChannelEffecter,
				Payload: PeerConnectionAndScaledrone{
					PeerConection: s.PeerConnection,
					Scaledrone:    s.Scaledrone,
				},
			},
		},
	}
}

type PeerConnectionAndScaledrone struct {
	PeerConection webrtc.PeerConnection
	Scaledrone    *scaledrone.Scaledrone
}

func CreateOfferAndDataChannelEffecter(dispatch hypp.Dispatch, payload hypp.Payload) {
	peerConnectionAndWebSocket := payload.(PeerConnectionAndScaledrone)
	pc := peerConnectionAndWebSocket.PeerConection
	sd := peerConnectionAndWebSocket.Scaledrone

	if !sd.IsConnected() {
		dispatch(AddSignalingError, state.NewSignalingError(
			"Failed to connect to the opponent",
			"Cannot create offer when WebSocket isn't connected.",
			errors.New("websocket not connected"),
		))
		return
	}

	dc := pc.CreateDataChannel("chat")
	dispatch(setupDataChannelListeners, dc)
	dispatch(SetDataChannel, dc)

	offer, err := pc.AwaitCreateOffer()
	if err != nil {
		dispatch(AddSignalingError, state.NewSignalingError(
			"Failed to connect to the opponent",
			"Failed to create peer connection offer",
			err,
		))
		return
	}
	if err := pc.AwaitSetLocalDescription(offer); err != nil {
		dispatch(AddSignalingError, state.NewSignalingError(
			"Failed to connect to the opponent",
			"Failed to set peer connection local description",
			err,
		))
		return
	}
	window.Console().Debug("Sending offer:", offer.Value)
	sd.SendMessage(OfferMessage{
		Kind:  "offer",
		Offer: offer,
	})
}

func SetPeerConnection(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	pc := payload.(webrtc.PeerConnection)
	newState := s.Clone()
	newState.PeerConnection = pc
	return newState
}

func setupDataChannelListeners(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	dc := payload.(webrtc.DataChannel)
	return hypp.StateAndEffects[*state.State]{
		State: s,
		Effects: []hypp.Effect{
			{
				Effecter: setupDataChannelListenersEffecter,
				Payload:  dc,
			},
		},
	}
}

func setupDataChannelListenersEffecter(dispatch hypp.Dispatch, payload hypp.Payload) {
	dc := payload.(webrtc.DataChannel)

	dc.SetOnOpen(func() {
		dispatch(DataChannelOpen, nil)
	})
	dc.SetOnMessage(func(event webrtc.MessageEvent) {
		dispatch(DataChannelMessage, event)
	})
	dc.SetOnClose(func() {
		dispatch(DataChannelClose, dc)
	})
	dc.SetOnError(func(err error) {
		dispatch(AddSignalingError, state.NewSignalingError("DataChannel error", err.Error(), err))
	})
}

func DataChannelOpen(s *state.State, _ hypp.Payload) hypp.Dispatchable {
	window.Console().Debug("DataChannel opened")
	newState := s.Clone()
	newState.WebRTCConnected = true
	return newState
}

func DataChannelMessage(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	return s
}

func DataChannelClose(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	dc := payload.(webrtc.DataChannel)
	window.Console().Debug("DataChannel closed", dc.Label())
	return hypp.StateAndEffects[*state.State]{
		State: s,
		Effects: []hypp.Effect{
			{Effecter: func(dispatch hypp.Dispatch, _ hypp.Payload) {
				dispatch(HangUp, nil)
			}},
		},
	}
}

func HangUpEffecter(dispatch hypp.Dispatch, payload hypp.Payload) {
	dispatch(HangUp, nil)
}

func HangUp(s *state.State, _ hypp.Payload) hypp.Dispatchable {
	newState := s.Clone()

	if !newState.DataChannel.IsUndefined() {
		newState.DataChannel.Close()
	}
	if !newState.PeerConnection.IsUndefined() {
		newState.PeerConnection.Close()
	}
	if newState.Scaledrone.IsConnected() {
		newState.Scaledrone.Reset()
	}

	newState.WebRTCConnected = false
	newState.OpponentWebsocketConnected = false
	newState.WebRTCConnected = false
	newState.SignalingErrors = nil
	newState.RoomName = ""
	newState.PeerConnection = webrtc.PeerConnection{}
	newState.DataChannel = webrtc.DataChannel{}
	newState.Scaledrone = scaledrone.NewScaledrone()
	return newState
}

func SetDataChannel(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	dc := payload.(webrtc.DataChannel)
	newState := s.Clone()
	newState.DataChannel = dc
	return newState
}

func OnDataChannelMessageSubscriber(dispatch hypp.Dispatch, payload hypp.Payload) hypp.Unsubscribe {
	dc := payload.(webrtc.DataChannel)
	dc.SetOnMessage(func(e webrtc.MessageEvent) {
		data := e.Data()
		window.Console().Debug("<<< Receive DataChannel message", data)
		var message CommitmentSchemeMessage[json.RawMessage]
		mustJSONUnmarshal([]byte(data), &message)
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

func AddSignalingError(s *state.State, payload hypp.Payload) hypp.Dispatchable {
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

	fmt.Println(signalingError.Description)

	newState := s.Clone()
	newState.SignalingErrors = append(newState.SignalingErrors, *signalingError)
	return newState
}
