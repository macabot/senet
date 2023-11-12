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
)

type CommitmentSchemeMessage[T any] struct {
	Kind CommitmentMessageKind
	Data T
}

func jsonStrigify(v any) string {
	return js.Global().Get("JSON").Call("stringify", v).String()
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
				state.DataChannel.Send(jsonStrigify(message))
			}()
		},
	}
}

func SendCommitmentAction() hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		newState.CommitmentScheme = state.CommitmentScheme{
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
				state.DataChannel.Send(jsonStrigify(message))
			}()
		},
	}
}
