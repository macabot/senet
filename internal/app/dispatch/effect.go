package dispatch

import (
	"time"

	"github.com/macabot/hypp"
)

func DelayedAction[S hypp.State](action hypp.Action[S], delay time.Duration) hypp.Effect {
	return Delayed(action, delay)
}

func Delayed(dispatchable hypp.Dispatchable, delay time.Duration) hypp.Effect {
	return hypp.Effect{
		Effecter: func(dispatch hypp.Dispatch, payload hypp.Payload) {
			go func() {
				time.Sleep(delay)
				dispatch(dispatchable, payload)
			}()
		},
	}
}
