package dispatch

import (
	"time"

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

type CommitmentMessageKind int

const (
	SendIsReadyKind CommitmentMessageKind = iota
	SendFlipperSecretKind
	SendCommitmentKind
	SendFlipperResultsKind
	SendCallerSecretAndPredictionsKind
	SendHasThrownKind
	SendMoveKind
	SendNoMoveKind
)

func (k CommitmentMessageKind) String() string {
	switch k {
	case SendIsReadyKind:
		return "SendIsReady"
	case SendFlipperSecretKind:
		return "SendFlipperSecret"
	case SendCommitmentKind:
		return "SendCommitment"
	case SendFlipperResultsKind:
		return "SendFlipperResults"
	case SendCallerSecretAndPredictionsKind:
		return "SendCallerSecretAndPredictions"
	case SendHasThrownKind:
		return "SendHasThrown"
	case SendMoveKind:
		return "SendMove"
	case SendNoMoveKind:
		return "SendNoMove"
	default:
		panic("invalid CommitmentMessageKind")
	}
}

func commitmentMessageKindFromString(s string) CommitmentMessageKind {
	switch s {
	case "SendIsReady":
		return SendIsReadyKind
	case "SendFlipperSecret":
		return SendFlipperSecretKind
	case "SendCommitment":
		return SendCommitmentKind
	case "SendFlipperResults":
		return SendFlipperResultsKind
	case "SendCallerSecretAndPredictions":
		return SendCallerSecretAndPredictionsKind
	case "SendHasThrown":
		return SendHasThrownKind
	case "SendMove":
		return SendMoveKind
	case "SendNoMove":
		return SendNoMoveKind
	default:
		panic("invalid CommitmentMessageKind string")
	}
}

type CallerSecretAndPredictions struct {
	Secret      string
	Predictions [4]bool
}

type Move struct {
	From state.Position
	To   state.Position
}

func (m Move) ToValue() js.Value {
	return js.ValueOf(map[string]any{
		"From": int(m.From),
		"To":   int(m.To),
	})
}

func parseMove(value js.Value) Move {
	return Move{
		From: state.Position(value.Get("From").Int()),
		To:   state.Position(value.Get("To").Int()),
	}
}

func flipsToValue(flips [4]bool) js.Value {
	return js.ValueOf([]any{flips[0], flips[1], flips[2], flips[3]})
}

func parseFlips(value js.Value) [4]bool {
	return [4]bool{
		value.Index(0).Bool(),
		value.Index(1).Bool(),
		value.Index(2).Bool(),
		value.Index(3).Bool(),
	}
}

func (c CallerSecretAndPredictions) ToValue() js.Value {
	return js.ValueOf(map[string]any{
		"Secret":      c.Secret,
		"Predictions": flipsToValue(c.Predictions),
	})
}

func parseCallerSecretAndPredictions(value js.Value) CallerSecretAndPredictions {
	return CallerSecretAndPredictions{
		Secret:      value.Get("Secret").String(),
		Predictions: parseFlips(value.Get("Predictions")),
	}
}

type CommitmentSchemeMessage[T any] struct {
	Kind CommitmentMessageKind
	Data T
}

func (m CommitmentSchemeMessage[T]) ToValue() js.Value {
	var data js.Value
	switch d := any(m.Data).(type) {
	case string:
		data = js.ValueOf(d)
	case [4]bool:
		data = flipsToValue(d)
	case CallerSecretAndPredictions:
		data = d.ToValue()
	case struct{}:
		data = js.Undefined()
	case Move:
		data = d.ToValue()
	default:
		panic("cannot convert CommitmentSchemeMessage.Data to js.Value")
	}
	return js.ValueOf(map[string]any{
		"Kind": m.Kind.String(),
		"Data": data,
	})
}

func ParseCommitmentSchemeMessage(s string) CommitmentSchemeMessage[js.Value] {
	parsed := jsonParse(s)
	return CommitmentSchemeMessage[js.Value]{
		Kind: commitmentMessageKindFromString(parsed.Get("Kind").String()),
		Data: parsed.Get("Data"),
	}
}

func jsonStringify(v js.Value) string {
	return js.Global().Get("JSON").Call("stringify", v).String()
}

func jsonParse(s string) js.Value {
	return js.Global().Get("JSON").Call("parse", s)
}

func valueToFlips(value js.Value) [4]bool {
	var flips [4]bool
	flips[0] = value.Index(0).Bool()
	flips[1] = value.Index(1).Bool()
	flips[2] = value.Index(2).Bool()
	flips[3] = value.Index(3).Bool()
	return flips
}

func valueToCallerSecretAndPredictions(value js.Value) CallerSecretAndPredictions {
	return CallerSecretAndPredictions{
		Secret:      value.Get("Secret").String(),
		Predictions: valueToFlips(value.Get("Predictions")),
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
			go func() {
				defer RecoverPanic(dispatch)

				message := CommitmentSchemeMessage[struct{}]{
					Kind: SendIsReadyKind,
				}
				sendDataChannelMessage(jsonStringify(message.ToValue()))
			}()
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
			go func() {
				defer RecoverPanic(dispatch)

				message := CommitmentSchemeMessage[string]{
					Kind: SendFlipperSecretKind,
					Data: flipperSecret,
				}
				sendDataChannelMessage(jsonStringify(message.ToValue()))
			}()
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
			go func() {
				defer RecoverPanic(dispatch)

				message := CommitmentSchemeMessage[string]{
					Kind: SendCommitmentKind,
					Data: commitment,
				}
				sendDataChannelMessage(jsonStringify(message.ToValue()))
			}()
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
			go func() {
				defer RecoverPanic(dispatch)

				message := CommitmentSchemeMessage[[4]bool]{
					Kind: SendFlipperResultsKind,
					Data: flipperResults,
				}
				sendDataChannelMessage(jsonStringify(message.ToValue()))
			}()
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
			go func() {
				defer RecoverPanic(dispatch)

				message := CommitmentSchemeMessage[CallerSecretAndPredictions]{
					Kind: SendCallerSecretAndPredictionsKind,
					Data: CallerSecretAndPredictions{
						Secret:      callerSecret,
						Predictions: callerPredictions,
					},
				}
				sendDataChannelMessage(jsonStringify(message.ToValue()))
			}()
		},
	}
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
			go func() {
				defer RecoverPanic(dispatch)

				message := CommitmentSchemeMessage[struct{}]{
					Kind: SendHasThrownKind,
				}
				sendDataChannelMessage(jsonStringify(message.ToValue()))
			}()
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
			go func() {
				defer RecoverPanic(dispatch)

				message := CommitmentSchemeMessage[Move]{
					Kind: SendMoveKind,
					Data: Move{
						From: from,
						To:   to,
					},
				}
				sendDataChannelMessage(jsonStringify(message.ToValue()))
			}()
		},
	}
}

func ReceiveMoveAction(move Move) hypp.Action[*state.State] {
	return func(s *state.State, _ hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		nextMove, err := newState.Game.Move(newState.Game.Turn, move.From, move.To)
		if err != nil {
			panic(err)
		}
		var effects []hypp.Effect
		if nextMove != nil {
			effects = append(effects, DelayedAction(
				MoveToSquareAction(nextMove.To),
				time.Second,
			))
		}

		isCaller := !newState.Game.HasTurn()
		effects = append(effects, sendIsReady(newState, isCaller)...)

		return hypp.StateAndEffects[*state.State]{
			State:   newState,
			Effects: effects,
		}
	}
}

func SendNoMoveEffect() hypp.Effect {
	return hypp.Effect{
		Effecter: func(dispatch hypp.Dispatch, _ hypp.Payload) {
			go func() {
				defer RecoverPanic(dispatch)

				message := CommitmentSchemeMessage[struct{}]{
					Kind: SendNoMoveKind,
				}
				sendDataChannelMessage(jsonStringify(message.ToValue()))
			}()
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
