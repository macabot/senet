package dispatch

import (
	"encoding/json"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/js"
	"github.com/macabot/hypp/window"
	"github.com/macabot/senet/internal/app/state"
)

func registerCommitmentScheme() {
	onThrowSticks = append(onThrowSticks, func(_, newState *state.State) []hypp.Effect {
		return []hypp.Effect{
			SendHasThrownEffect(),
		}
	})

	onMoveToSquare = append(
		onMoveToSquare,
		func(s, newState *state.State, from, to state.Position) []hypp.Effect {
			effects := []hypp.Effect{
				SendMoveEffect(from, to),
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
				SendNoMoveEffect(),
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

type Move struct {
	From state.Position
	To   state.Position
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
		SendIsReadyEffect(),
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

func SendIsReadyEffect() hypp.Effect {
	return hypp.Effect{
		Effecter: func(dispatch hypp.Dispatch, _ hypp.Payload) {
			message := CommitmentSchemeMessage[struct{}]{
				Kind: SendIsReadyKind,
			}
			sendDataChannelMessage(mustJSONMarshal(message))
		},
	}
}

func ReceiveIsReadyAction() hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
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
}

func sendFlipperSecret(newState *state.State) []hypp.Effect {
	newState.CommitmentScheme.IsCaller = false
	newState.CommitmentScheme.FlipperSecret = state.GenerateSecret()
	return []hypp.Effect{
		SendFlipperSecretEffect(newState.CommitmentScheme.FlipperSecret),
	}
}

func SendFlipperSecretEffect(flipperSecret string) hypp.Effect {
	return hypp.Effect{
		Effecter: func(dispatch hypp.Dispatch, _ hypp.Payload) {
			message := CommitmentSchemeMessage[string]{
				Kind: SendFlipperSecretKind,
				Data: flipperSecret,
			}
			sendDataChannelMessage(mustJSONMarshal(message))
		},
	}
}

func ReceiveFlipperSecretAction(flipperSecret string) hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		newState.CommitmentScheme.IsCaller = true
		newState.CommitmentScheme.FlipperSecret = flipperSecret
		return sendCommitment(newState)
	}
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
			SendCommitmentEffect(newState.CommitmentScheme.Commitment),
		},
	}
}

func SendCommitmentEffect(commitment string) hypp.Effect {
	return hypp.Effect{
		Effecter: func(dispatch hypp.Dispatch, _ hypp.Payload) {
			message := CommitmentSchemeMessage[string]{
				Kind: SendCommitmentKind,
				Data: commitment,
			}
			sendDataChannelMessage(mustJSONMarshal(message))
		},
	}
}

func ReceiveCommitmentAction(commitment string) hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		if s.CommitmentScheme.IsCaller {
			window.Console().Warn("Caller received commitment")
			return s
		}

		newState := s.Clone()
		newState.CommitmentScheme.Commitment = commitment
		return sendFlipperResults(newState)
	}
}

func sendFlipperResults(newState *state.State) hypp.StateAndEffects[*state.State] {
	newState.CommitmentScheme.FlipperResults = state.GenerateFlips()
	newState.CommitmentScheme.HasFlipperResults = true
	return hypp.StateAndEffects[*state.State]{
		State: newState,
		Effects: []hypp.Effect{
			SendFlipperResultsEffect(newState.CommitmentScheme.FlipperResults),
		},
	}
}

func SendFlipperResultsEffect(flipperResults [4]bool) hypp.Effect {
	return hypp.Effect{
		Effecter: func(dispatch hypp.Dispatch, _ hypp.Payload) {
			message := CommitmentSchemeMessage[[4]bool]{
				Kind: SendFlipperResultsKind,
				Data: flipperResults,
			}
			sendDataChannelMessage(mustJSONMarshal(message))
		},
	}
}

func ReceiveFlipperResultsAction(flipperResults [4]bool) hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		if !s.CommitmentScheme.IsCaller {
			window.Console().Warn("Flipper received flipperResults")
			return s
		}

		newState := s.Clone()
		newState.CommitmentScheme.FlipperResults = flipperResults
		newState.CommitmentScheme.HasFlipperResults = true
		return sendCallerSecretAndPredictions(newState)
	}
}

func sendCallerSecretAndPredictions(newState *state.State) hypp.StateAndEffects[*state.State] {
	return hypp.StateAndEffects[*state.State]{
		State: newState,
		Effects: []hypp.Effect{
			SendCallerSecretAndPredictionsEffect(
				newState.CommitmentScheme.CallerSecret,
				newState.CommitmentScheme.CallerPredictions,
			),
		},
	}
}

func SendCallerSecretAndPredictionsEffect(
	callerSecret string,
	callerPredictions [4]bool,
) hypp.Effect {
	return hypp.Effect{
		Effecter: func(dispatch hypp.Dispatch, _ hypp.Payload) {
			message := CommitmentSchemeMessage[CallerSecretAndPredictions]{
				Kind: SendCallerSecretAndPredictionsKind,
				Data: CallerSecretAndPredictions{
					Secret:      callerSecret,
					Predictions: callerPredictions,
				},
			}
			sendDataChannelMessage(mustJSONMarshal(message))
		},
	}
}

func flipsToValue(flips [4]bool) js.Value {
	return js.ValueOf([]any{flips[0], flips[1], flips[2], flips[3]})
}

func ReceiveCallerSecretAndPredictionsAction(callerSecretAndPredictions CallerSecretAndPredictions) hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
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
}

func SendHasThrownEffect() hypp.Effect {
	return hypp.Effect{
		Effecter: func(dispatch hypp.Dispatch, _ hypp.Payload) {
			message := CommitmentSchemeMessage[struct{}]{
				Kind: SendHasThrownKind,
			}
			sendDataChannelMessage(mustJSONMarshal(message))
		},
	}
}

func ReceiveHasThrownAction() hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		newState.Game.ThrowSticks(newState)
		return newState
	}
}

func SendMoveEffect(from, to state.Position) hypp.Effect {
	return hypp.Effect{
		Effecter: func(dispatch hypp.Dispatch, _ hypp.Payload) {
			message := CommitmentSchemeMessage[Move]{
				Kind: SendMoveKind,
				Data: Move{
					From: from,
					To:   to,
				},
			}
			sendDataChannelMessage(mustJSONMarshal(message))
		},
	}
}

func ReceiveMoveAction(move Move) hypp.Action[*state.State] {
	return func(s *state.State, _ hypp.Payload) hypp.Dispatchable {
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
}

func SendNoMoveEffect() hypp.Effect {
	return hypp.Effect{
		Effecter: func(dispatch hypp.Dispatch, _ hypp.Payload) {
			message := CommitmentSchemeMessage[struct{}]{
				Kind: SendNoMoveKind,
			}
			sendDataChannelMessage(mustJSONMarshal(message))
		},
	}
}

func ReceiveNoMoveAction() hypp.Action[*state.State] {
	return func(s *state.State, _ hypp.Payload) hypp.Dispatchable {
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
}
