package app

import (
	"encoding/json"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/js"
	"github.com/macabot/hypp/window"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
	"github.com/macabot/senet/internal/pkg/sessionstorage"
)

var debug = false

// dispatchWrapper recovers panics throws in effects and actions.
// An effect that starts a new go routine will not recover from a panic.
// The effect can defer [dispatch.RecoverEffectPanic] inside the go routine to recover any panic.
//
// If debug=true, the state will be stored in the session storage.
func dispatchWrapper(dispatchFunc hypp.Dispatch) hypp.Dispatch {
	return func(dispatchable hypp.Dispatchable, payload hypp.Payload) {
		var s *state.State
		stateFound := false
		switch v := dispatchable.(type) {
		case hypp.StateAndEffects[*state.State]:
			s = v.State
			stateFound = true
			dispatchable = dispatch.RecoverWrapStateAndEffects(v)
		case hypp.Action[*state.State]:
			dispatchable = dispatch.RecoverWrapAction(v)
		case hypp.ActionAndPayload[*state.State]:
			dispatchable = hypp.ActionAndPayload[*state.State]{
				Action:  dispatch.RecoverWrapAction(v.Action),
				Payload: v.Payload,
			}
		case *state.State:
			s = v
			stateFound = true
		}

		dispatchFunc(dispatchable, payload)

		if stateFound && debug {
			prevState := sessionstorage.GetItem("state")
			if prevState != nil {
				sessionstorage.SetItem("prevState", *prevState)
			}

			b, err := json.Marshal(s)
			if err != nil {
				b = []byte(err.Error())
			}
			sessionstorage.SetItem("state", string(b))
		}
	}
}

func Run(element window.Element) {
	urlSearchParams := js.Global().Get("URLSearchParams")
	urlParams := urlSearchParams.New(js.Global().Get("location").Get("search"))
	debugParam := urlParams.Call("get", "debug")
	debug = !debugParam.IsNull() && debugParam.String() == "true"

	if !debug {
		window.AddEventListener("beforeunload", func(e window.Event) {
			e.PreventDefault()
		})
	}

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
