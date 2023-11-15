package dispatch

import (
	"fmt"

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
				fmt.Println("DataChannel readyState", state.DataChannel.ReadyState())
				state.DataChannel.Send("Test")
				// message := CommitmentSchemeMessage[string]{
				// 	Kind: SendFlipperSecretKind,
				// 	Data: flipperSecret,
				// }
				// state.DataChannel.Send(jsonStringify(message.ToValue()))
			}()
		},
	}
}

func SendCommitmentAction() hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
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

func SendFlipperResultsAction() hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
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

func SendCallerSecretAndPredictionsAction() hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		return hypp.StateAndEffects[*state.State]{
			State: s,
			Effects: []hypp.Effect{
				SendCallerSecretAndPredictionsEffect(
					s.CommitmentScheme.CallerSecret,
					s.CommitmentScheme.CallerPredictions,
				),
			},
		}
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
