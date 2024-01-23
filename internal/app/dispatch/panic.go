package dispatch

import (
	"fmt"
	"runtime/debug"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/js"
	"github.com/macabot/hypp/window"
	"github.com/macabot/senet/internal/app/state"
)

func SetPanicStackTraceAction(panicStackTrace *string) hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		newState.PanicStackTrace = panicStackTrace
		return newState
	}
}

func RecoverEffectPanic(dispatch hypp.Dispatch) {
	if r := recover(); r != nil {
		panicStackTrace := fmt.Sprintf("%v\n%s", r, string(debug.Stack()))
		window.Console().Error(panicStackTrace)
		dispatch(SetPanicStackTraceAction(&panicStackTrace), nil)
	}
}

func RecoverWrapStateAndEffects(stateAndEffects hypp.StateAndEffects[*state.State]) hypp.StateAndEffects[*state.State] {
	wrappedEffects := make([]hypp.Effect, len(stateAndEffects.Effects))
	for i, e := range stateAndEffects.Effects {
		eCopy := e // Can probably be removed in Go 1.22.
		wrappedEffects[i] = hypp.Effect{
			Effecter: func(dispatch hypp.Dispatch, payload hypp.Payload) {
				defer RecoverEffectPanic(dispatch)
				eCopy.Effecter(dispatch, payload)
			},
			Payload: eCopy.Payload,
		}
	}
	return hypp.StateAndEffects[*state.State]{
		State:   stateAndEffects.State,
		Effects: wrappedEffects,
	}
}

func RecoverWrapAction(action hypp.Action[*state.State]) hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) (dispatchable hypp.Dispatchable) {
		defer func() {
			if r := recover(); r != nil {
				panicStackTrace := fmt.Sprintf("%v\n%s", r, string(debug.Stack()))
				window.Console().Error(panicStackTrace)
				dispatchable = SetPanicStackTraceAction(&panicStackTrace)
			}
		}()
		return action(s, payload)
	}
}

func ReloadPageAction() hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		return hypp.StateAndEffects[*state.State]{
			State: s,
			Effects: []hypp.Effect{
				ReloadPageEffect(),
			},
		}
	}
}

func ReloadPageEffect() hypp.Effect {
	return hypp.Effect{
		Effecter: func(_ hypp.Dispatch, _ hypp.Payload) {
			js.Global().Get("location").Call("reload")
		},
	}
}
