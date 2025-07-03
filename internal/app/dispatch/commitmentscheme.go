package dispatch

import (
	"encoding/json"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/js"
	"github.com/macabot/hypp/window"
	"github.com/macabot/senet/internal/app/state"
	"github.com/macabot/senet/internal/pkg/webrtc"
)

func registerCommitmentScheme() {
	onThrowSticks = append(onThrowSticks, func(_, newState *state.State) []hypp.Effect {
		return []hypp.Effect{
			{
				Effecter: SendHasThrownEffecter,
				Payload:  newState.DataChannel,
			},
		}
	})

	onMoveToSquare = append(
		onMoveToSquare,
		func(s, newState *state.State, from, to state.Position) []hypp.Effect {
			effects := []hypp.Effect{
				{
					Effecter: SendMoveEffecter,
					Payload: SendMovePayload{
						DataChannel: newState.DataChannel,
						Move: Move{
							From: from,
							To:   to,
						},
					},
				},
			}
			isCaller := !newState.Game.HasTurn()
			effects = append(effects, sendIsReady(newState, isCaller)...)
			return effects
		},
	)

	onNoMove = append(
		onNoMove,
		func(s, newState *state.State) []hypp.Effect {
			effects := []hypp.Effect{
				{
					Effecter: SendNoMoveEffecter,
					Payload:  newState.DataChannel,
				},
			}
			isCaller := !newState.Game.HasTurn()
			effects = append(effects, sendIsReady(newState, isCaller)...)
			return effects
		},
	)
}

type CommitmentMessageKind string

const (
	SendIsReadyKind                    CommitmentMessageKind = "SendIsReady"
	SendFlipperSecretKind              CommitmentMessageKind = "SendFlipperSecret"
	SendCommitmentKind                 CommitmentMessageKind = "SendCommitment"
	SendFlipperResultsKind             CommitmentMessageKind = "SendFlipperResults"
	SendCallerSecretAndPredictionsKind CommitmentMessageKind = "SendCallerSecretAndPredictions"
	SendHasThrownKind                  CommitmentMessageKind = "SendHasThrown"
	SendMoveKind                       CommitmentMessageKind = "SendMove"
	SendNoMoveKind                     CommitmentMessageKind = "SendNoMove"
)

type CallerSecretAndPredictions struct {
	Secret      string
	Predictions [4]bool
}

type CommitmentSchemeMessage[T any] struct {
	Kind CommitmentMessageKind
	Data T
}

func mustJSONMarshal(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func mustJSONUnmarshal(b []byte, v any) {
	if err := json.Unmarshal(b, v); err != nil {
		panic(err)
	}
}

func sendIsReady(newState *state.State, isCaller bool) []hypp.Effect {
	newState.CommitmentScheme = state.CommitmentScheme{
		IsReady:         true,
		OpponentIsReady: newState.CommitmentScheme.OpponentIsReady,
		IsCaller:        isCaller,
	}
	effects := []hypp.Effect{
		{Effecter: SendIsReadyEffecter, Payload: newState.DataChannel},
	}
	if !isCaller && newState.CommitmentScheme.OpponentIsReady {
		effects = append(effects, sendFlipperSecret(newState)...)
	}
	if newState.CommitmentScheme.OpponentIsReady {
		newState.CommitmentScheme.IsReady = false
		newState.CommitmentScheme.OpponentIsReady = false
	}
	return effects
}

func SendIsReadyEffecter(dispatch hypp.Dispatch, payload hypp.Payload) {
	dc := payload.(webrtc.DataChannel)
	message := CommitmentSchemeMessage[struct{}]{
		Kind: SendIsReadyKind,
	}
	dc.Send(mustJSONMarshal(message))
}

func ReceiveIsReady(s *state.State, _ hypp.Payload) hypp.Dispatchable {
	newState := s.Clone()
	newState.CommitmentScheme.OpponentIsReady = true
	var effects []hypp.Effect
	if !newState.CommitmentScheme.IsCaller && newState.CommitmentScheme.IsReady {
		effects = append(effects, sendFlipperSecret(newState)...)
	}
	if newState.CommitmentScheme.IsReady {
		newState.CommitmentScheme.IsReady = false
		newState.CommitmentScheme.OpponentIsReady = false
	}
	return hypp.StateAndEffects[*state.State]{
		State:   newState,
		Effects: effects,
	}
}

func sendFlipperSecret(newState *state.State) []hypp.Effect {
	newState.CommitmentScheme.IsCaller = false
	newState.CommitmentScheme.FlipperSecret = state.GenerateSecret()
	return []hypp.Effect{
		{
			Effecter: SendFlipperSecretEffecter,
			Payload: SendFlipperSecretPayload{
				DataChannel:   newState.DataChannel,
				FlipperSecret: newState.CommitmentScheme.FlipperSecret,
			},
		},
	}
}

type SendFlipperSecretPayload struct {
	DataChannel   webrtc.DataChannel
	FlipperSecret string
}

func SendFlipperSecretEffecter(dispatch hypp.Dispatch, payload hypp.Payload) {
	p := payload.(SendFlipperSecretPayload)
	message := CommitmentSchemeMessage[string]{
		Kind: SendFlipperSecretKind,
		Data: p.FlipperSecret,
	}
	p.DataChannel.Send(mustJSONMarshal(message))
}

func ReceiveFlipperSecret(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	flipperSecret := payload.(string)
	newState := s.Clone()
	newState.CommitmentScheme.IsCaller = true
	newState.CommitmentScheme.FlipperSecret = flipperSecret
	return sendCommitment(newState)
}

func sendCommitment(newState *state.State) hypp.StateAndEffects[*state.State] {
	newState.CommitmentScheme.CallerSecret = state.GenerateSecret()
	newState.CommitmentScheme.CallerPredictions = state.GenerateFlips()
	newState.CommitmentScheme.HasCallerPredictions = true
	newState.CommitmentScheme.Commitment = state.GenerateCommitmentHash(
		newState.CommitmentScheme.CallerSecret,
		newState.CommitmentScheme.FlipperSecret,
		newState.CommitmentScheme.CallerPredictions,
	)
	return hypp.StateAndEffects[*state.State]{
		State: newState,
		Effects: []hypp.Effect{
			{
				Effecter: SendCommitmentEffecter,
				Payload: SendCommitmentPayload{
					DataChannel: newState.DataChannel,
					Commitment:  newState.CommitmentScheme.Commitment,
				},
			},
		},
	}
}

type SendCommitmentPayload struct {
	DataChannel webrtc.DataChannel
	Commitment  string
}

func SendCommitmentEffecter(dispatch hypp.Dispatch, payload hypp.Payload) {
	p := payload.(SendCommitmentPayload)
	message := CommitmentSchemeMessage[string]{
		Kind: SendCommitmentKind,
		Data: p.Commitment,
	}
	p.DataChannel.Send(mustJSONMarshal(message))
}

func ReceiveCommitment(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	if s.CommitmentScheme.IsCaller {
		window.Console().Warn("Caller received commitment")
		return s
	}

	commitment := payload.(string)
	newState := s.Clone()
	newState.CommitmentScheme.Commitment = commitment
	return sendFlipperResults(newState)
}

func sendFlipperResults(newState *state.State) hypp.StateAndEffects[*state.State] {
	newState.CommitmentScheme.FlipperResults = state.GenerateFlips()
	newState.CommitmentScheme.HasFlipperResults = true
	return hypp.StateAndEffects[*state.State]{
		State: newState,
		Effects: []hypp.Effect{
			{
				Effecter: SendFlipperResultsEffecter,
				Payload: SendFlipperResultsPayload{
					DataChannel:    newState.DataChannel,
					FlipperResults: newState.CommitmentScheme.FlipperResults,
				},
			},
		},
	}
}

type SendFlipperResultsPayload struct {
	DataChannel    webrtc.DataChannel
	FlipperResults [4]bool
}

func SendFlipperResultsEffecter(dispatch hypp.Dispatch, payload hypp.Payload) {
	p := payload.(SendFlipperResultsPayload)
	message := CommitmentSchemeMessage[[4]bool]{
		Kind: SendFlipperResultsKind,
		Data: p.FlipperResults,
	}
	p.DataChannel.Send(mustJSONMarshal(message))
}

func ReceiveFlipperResults(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	if !s.CommitmentScheme.IsCaller {
		window.Console().Warn("Flipper received flipperResults")
		return s
	}

	flipperResults := payload.([4]bool)
	newState := s.Clone()
	newState.CommitmentScheme.FlipperResults = flipperResults
	newState.CommitmentScheme.HasFlipperResults = true
	return sendCallerSecretAndPredictions(newState)
}

func sendCallerSecretAndPredictions(newState *state.State) hypp.StateAndEffects[*state.State] {
	return hypp.StateAndEffects[*state.State]{
		State: newState,
		Effects: []hypp.Effect{
			{
				Effecter: SendCallerSecretAndPredictionsEffecter,
				Payload: SendCallerSecretAndPredictionsPayload{
					DataChannel:       newState.DataChannel,
					CallerSecret:      newState.CommitmentScheme.CallerSecret,
					CallerPredictions: newState.CommitmentScheme.CallerPredictions,
				},
			},
		},
	}
}

type SendCallerSecretAndPredictionsPayload struct {
	DataChannel       webrtc.DataChannel
	CallerSecret      string
	CallerPredictions [4]bool
}

func SendCallerSecretAndPredictionsEffecter(dispatch hypp.Dispatch, payload hypp.Payload) {
	p := payload.(SendCallerSecretAndPredictionsPayload)
	message := CommitmentSchemeMessage[CallerSecretAndPredictions]{
		Kind: SendCallerSecretAndPredictionsKind,
		Data: CallerSecretAndPredictions{
			Secret:      p.CallerSecret,
			Predictions: p.CallerPredictions,
		},
	}
	p.DataChannel.Send(mustJSONMarshal(message))
}

func flipsToValue(flips [4]bool) js.Value {
	return js.ValueOf([]any{flips[0], flips[1], flips[2], flips[3]})
}

func ReceiveCallerSecretAndPredictions(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	callerSecretAndPredictions := payload.(CallerSecretAndPredictions)
	newState := s.Clone()
	newState.CommitmentScheme.CallerSecret = callerSecretAndPredictions.Secret
	newState.CommitmentScheme.CallerPredictions = callerSecretAndPredictions.Predictions
	newState.CommitmentScheme.HasCallerPredictions = true
	isExpectedCommitment := state.IsExpectedCommitment(
		newState.CommitmentScheme.CallerSecret,
		newState.CommitmentScheme.FlipperSecret,
		newState.CommitmentScheme.CallerPredictions,
		newState.CommitmentScheme.Commitment,
	)
	if !isExpectedCommitment {
		window.Console().Warn(
			"Unexpected commitment",
			newState.CommitmentScheme.CallerSecret,
			newState.CommitmentScheme.FlipperSecret,
			flipsToValue(newState.CommitmentScheme.CallerPredictions),
			newState.CommitmentScheme.Commitment,
		)
	}

	return newState
}

func SendHasThrownEffecter(dispatch hypp.Dispatch, payload hypp.Payload) {
	dc := payload.(webrtc.DataChannel)
	message := CommitmentSchemeMessage[struct{}]{
		Kind: SendHasThrownKind,
	}
	dc.Send(mustJSONMarshal(message))
}

func ReceiveHasThrown(s *state.State, _ hypp.Payload) hypp.Dispatchable {
	newState := s.Clone()
	newState.Game.ThrowSticks(newState)
	return newState
}

type SendMovePayload struct {
	DataChannel webrtc.DataChannel
	Move        Move
}

func SendMoveEffecter(dispatch hypp.Dispatch, payload hypp.Payload) {
	p := payload.(SendMovePayload)
	message := CommitmentSchemeMessage[Move]{
		Kind: SendMoveKind,
		Data: p.Move,
	}
	p.DataChannel.Send(mustJSONMarshal(message))
}

func ReceiveMove(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	move := payload.(Move)
	newState := s.Clone()
	// Ignore the nextMove, because we will receive a separate message for it.
	_, err := newState.Game.Move(newState.Game.Turn, move.From, move.To)
	if err != nil {
		panic(err)
	}

	isCaller := !newState.Game.HasTurn()
	effects := sendIsReady(newState, isCaller)

	return hypp.StateAndEffects[*state.State]{
		State:   newState,
		Effects: effects,
	}
}

func SendNoMoveEffecter(dispatch hypp.Dispatch, payload hypp.Payload) {
	dc := payload.(webrtc.DataChannel)
	message := CommitmentSchemeMessage[struct{}]{
		Kind: SendNoMoveKind,
	}
	dc.Send(mustJSONMarshal(message))
}

func ReceiveNoMove(s *state.State, _ hypp.Payload) hypp.Dispatchable {
	newState := s.Clone()
	if err := newState.Game.NoMove(newState.Game.Turn); err != nil {
		panic(err)
	}

	isCaller := !newState.Game.HasTurn()
	effects := sendIsReady(newState, isCaller)

	return hypp.StateAndEffects[*state.State]{
		State:   newState,
		Effects: effects,
	}
}
