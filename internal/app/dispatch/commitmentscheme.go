package dispatch

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/js"
	"github.com/macabot/senet/internal/app/state"
)

type CommitmentMessageKind int

const (
	SendFlipperSecretKind CommitmentMessageKind = iota
	SendCommitmentKind
	SendFlipperResultsKind
	SendCallerSecretAndPredictionsKind
)

type CallerSecretAndPredictions struct {
	Secret      string
	Predictions [4]bool
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
	default:
		panic("cannot convert CommitmentSchemeMessage.Data to js.Value")
	}
	return js.ValueOf(map[string]any{
		"Kind": int(m.Kind),
		"Data": data,
	})
}

func ParseCommitmentSchemeMessage(s string) CommitmentSchemeMessage[js.Value] {
	parsed := jsonParse(s)
	return CommitmentSchemeMessage[js.Value]{
		Kind: CommitmentMessageKind(parsed.Get("Kind").Int()),
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

func SendFlipperSecretAction() hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		newState.CommitmentScheme = state.CommitmentScheme{
			FlipperSecret: state.GenerateSecret(),
		}
		return hypp.StateAndEffects[*state.State]{
			State: newState,
			Effects: []hypp.Effect{
				SendFlipperSecretEffect(newState.CommitmentScheme.FlipperSecret),
			},
		}
	}
}

func SendFlipperSecretEffect(flipperSecret string) hypp.Effect {
	return hypp.Effect{
		Effecter: func(_ hypp.Dispatch, _ hypp.Payload) {
			go func() {
				message := CommitmentSchemeMessage[string]{
					Kind: SendFlipperSecretKind,
					Data: flipperSecret,
				}
				state.DataChannel.Send(jsonStringify(message.ToValue()))
			}()
		},
	}
}

func ReceiveFlipperSecretAction(flipperSecret string) hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		newState.CommitmentScheme.FlipperSecret = flipperSecret
		return sendCommitment(newState)
	}
}

func sendCommitment(newState *state.State) hypp.StateAndEffects[*state.State] {
	newState.CommitmentScheme = state.CommitmentScheme{
		IsCaller:             true,
		CallerSecret:         state.GenerateSecret(),
		FlipperSecret:        newState.CommitmentScheme.FlipperSecret,
		HasCallerPredictions: true,
		CallerPredictions:    state.GenerateFlips(),
	}
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
		Effecter: func(_ hypp.Dispatch, _ hypp.Payload) {
			go func() {
				message := CommitmentSchemeMessage[string]{
					Kind: SendCommitmentKind,
					Data: commitment,
				}
				state.DataChannel.Send(jsonStringify(message.ToValue()))
			}()
		},
	}
}

func ReceiveCommitmentAction(commitment string) hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		newState.CommitmentScheme.Commitment = commitment
		return sendFlipperResults(newState)
	}
}

func sendFlipperResults(newState *state.State) hypp.StateAndEffects[*state.State] {
	newState.CommitmentScheme = state.CommitmentScheme{
		FlipperSecret:     newState.CommitmentScheme.FlipperSecret,
		HasFlipperResults: true,
		FlipperResults:    state.GenerateFlips(),
		Commitment:        newState.CommitmentScheme.Commitment,
	}
	return hypp.StateAndEffects[*state.State]{
		State: newState,
		Effects: []hypp.Effect{
			SendFlipperResultsEffect(newState.CommitmentScheme.FlipperResults),
		},
	}
}

func SendFlipperResultsEffect(flipperResults [4]bool) hypp.Effect {
	return hypp.Effect{
		Effecter: func(_ hypp.Dispatch, _ hypp.Payload) {
			go func() {
				message := CommitmentSchemeMessage[[4]bool]{
					Kind: SendFlipperResultsKind,
					Data: flipperResults,
				}
				state.DataChannel.Send(jsonStringify(message.ToValue()))
			}()
		},
	}
}

func ReceiveFlipperResultsAction(flipperResults [4]bool) hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		newState.CommitmentScheme.FlipperResults = flipperResults
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
		Effecter: func(_ hypp.Dispatch, _ hypp.Payload) {
			go func() {
				message := CommitmentSchemeMessage[CallerSecretAndPredictions]{
					Kind: SendCallerSecretAndPredictionsKind,
					Data: CallerSecretAndPredictions{
						Secret:      callerSecret,
						Predictions: callerPredictions,
					},
				}
				state.DataChannel.Send(jsonStringify(message.ToValue()))
			}()
		},
	}
}
