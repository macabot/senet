package app

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/window"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
)

// dispatchWrapper recovers panics throws in effects and actions.
// An effect that starts a new go routine will not recover from a panic.
// The effect can defer [dispatch.RecoverEffectPanic] inside the go routine to recover any panic.
func dispatchWrapper(dispatchFunc hypp.Dispatch) hypp.Dispatch {
	return func(dispatchable hypp.Dispatchable, payload hypp.Payload) {
		switch v := dispatchable.(type) {
		case hypp.StateAndEffects[*state.State]:
			dispatchable = dispatch.RecoverWrapStateAndEffects(v)
		case hypp.Action[*state.State]:
			dispatchable = dispatch.RecoverWrapAction(v)
		case hypp.ActionAndPayload[*state.State]:
			dispatchable = hypp.ActionAndPayload[*state.State]{
				Action:  dispatch.RecoverWrapAction(v.Action),
				Payload: v.Payload,
			}
		}

		dispatchFunc(dispatchable, payload)
	}
}

func Run(element window.Element) {
	hypp.App(hypp.AppProps[*state.State]{
		Init: &state.State{
			Page: state.StartPage,
		},
		View: func(s *state.State) *hypp.VNode {
			return component.RecoverPanic(component.Senet, s)
		},
		Node: element,
		Subscriptions: func(s *state.State) []hypp.Subscription {
			initialized := s.Signaling != nil && s.Signaling.Initialized
			return []hypp.Subscription{
				{
					Subscriber: dispatch.OnICEConnectionStateChangeSubscriber,
					Disabled:   !initialized,
				},
				{
					Subscriber: dispatch.OnConnectionStateChangeSubscriber,
					Disabled:   !initialized,
				},
				{
					Subscriber: dispatch.OnDataChannelOpenSubscriber,
					Disabled:   !initialized,
				},
				{
					Subscriber: dispatch.OnDataChannelMessageSubscriber,
					Disabled:   !initialized,
				},
			}
		},
		DispatchWrapper: dispatchWrapper,
	})

	select {} // keep app running
}
