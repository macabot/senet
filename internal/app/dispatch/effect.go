package dispatch

import (
	"time"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/window"
	"github.com/macabot/senet/internal/app/state"
)

func Delayed(dispatchable hypp.Dispatchable, delay time.Duration) hypp.Effect {
	return hypp.Effect{
		Effecter: func(dispatch hypp.Dispatch, payload hypp.Payload) {
			go func() {
				defer RecoverEffectPanic(dispatch)

				time.Sleep(delay)
				window.RequestAnimationFrame(func() {
					dispatch(dispatchable, payload)
				})
			}()
		},
	}
}

func EffectsAction(effects ...hypp.Effect) hypp.Action[*state.State] {
	return func(s *state.State, _ hypp.Payload) hypp.Dispatchable {
		return hypp.StateAndEffects[*state.State]{
			State:   s,
			Effects: effects,
		}
	}
}
