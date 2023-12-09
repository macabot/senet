package app

import (
	"encoding/json"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/window"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
	"github.com/macabot/senet/internal/pkg/localstorage"
)

func Run(element window.Element) {
	s := &state.State{
		Page: state.StartPage,
	}
	v, ok := localstorage.GetItem("state")
	if ok {
		if err := json.Unmarshal([]byte(v), s); err != nil {
			panic(err)
		}
	}

	hypp.App(hypp.AppProps[*state.State]{
		Init: s,
		View: component.Senet,
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
		DispatchWrapper: func(dispatch hypp.Dispatch) hypp.Dispatch {
			return func(dispatchable hypp.Dispatchable, payload hypp.Payload) {
				dispatch(dispatchable, payload)
				var s *state.State
				switch v := dispatchable.(type) {
				case hypp.StateAndEffects[*state.State]:
					s = v.State
				case *state.State: // State
					s = v
				default:
					return
				}
				b, err := json.Marshal(s)
				if err != nil {
					panic(err)
				}
				localstorage.SetItem("state", string(b))
			}
		},
	})
}
