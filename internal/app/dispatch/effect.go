package dispatch

import (
	"time"

	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

func DelayedAction(action hypp.Action[*state.State], delay time.Duration) hypp.Effect {
	return Delayed(action, delay)
}

func Delayed(dispatchable hypp.Dispatchable, delay time.Duration) hypp.Effect {
	return hypp.Effect{
		Effecter: func(dispatch hypp.Dispatch, payload hypp.Payload) {
			go func() {
				defer RecoverPanic(dispatch)

				time.Sleep(delay)
				dispatch(dispatchable, payload)
			}()
		},
	}
}

func EffectsAction(effects ...hypp.Effect) hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		return hypp.StateAndEffects[*state.State]{
			State:   s,
			Effects: effects,
		}
	}
}
