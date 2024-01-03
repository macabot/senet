package dispatch

import (
	"fmt"
	"runtime/debug"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/js"
	"github.com/macabot/hypp/window"
	"github.com/macabot/senet/internal/app/state"
	"github.com/macabot/senet/internal/pkg/localstorage"
)

func SetPanicTraceAction(panicTrace *string) hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		newState.PanicTrace = panicTrace
		return newState
	}
}

func RecoverPanic(dispatch hypp.Dispatch) {
	if r := recover(); r != nil {
		s := fmt.Sprintf("%v\n%s", r, string(debug.Stack()))
		window.Console().Error(s)
		dispatch(SetPanicTraceAction(&s), nil)
	}
}

func ResetAction() hypp.Action[*state.State] {
	return func(_ *state.State, payload hypp.Payload) hypp.Dispatchable {
		return hypp.StateAndEffects[*state.State]{
			State: &state.State{
				Page: state.StartPage,
			},
			Effects: []hypp.Effect{
				ResetEffect(),
			},
		}
	}
}

func ResetEffect() hypp.Effect {
	return hypp.Effect{
		Effecter: func(_ hypp.Dispatch, _ hypp.Payload) {
			localstorage.RemoveItem("state")
			js.Global().Get("location").Call("reload")
		},
	}
}
