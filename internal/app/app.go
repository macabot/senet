package app

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/window"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/state"
)

func Run(element window.Element) {
	hypp.App(hypp.AppProps[*state.State]{
		Init: &state.State{
			Page: state.StartPage,
		},
		View: component.Senet,
		Node: element,
		// FIXME this causes error: Uncaught InternalError: too much recursion
		// Subscriptions: func(s *state.State) []hypp.Subscription {
		// 	initialized := s.Signaling != nil && s.Signaling.Initialized
		// 	fmt.Println("Signaling.Initialized", initialized)
		// 	return []hypp.Subscription{
		// 		{
		// 			Subscriber: dispatch.OnICEConnectionStateChangeSubscriber,
		// 			Disabled:   !initialized,
		// 		},
		// 		{
		// 			Subscriber: dispatch.OnConnectionStateChangeSubscriber,
		// 			Disabled:   !initialized,
		// 		},
		// 		{
		// 			Subscriber: dispatch.OnDataChannelOpenSubscriber,
		// 			Disabled:   !initialized,
		// 		},
		// 		{
		// 			Subscriber: dispatch.OnDataChannelMessageSubscriber,
		// 			Disabled:   !initialized,
		// 		},
		// 	}
		// },
	})
}
