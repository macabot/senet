package app

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/window"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
)

func Run(element window.Element) {
	hypp.App(hypp.AppProps[*state.State]{
		Init: &state.State{
			Page: state.StartPage,
		},
		View: component.Senet,
		Node: element,
		Subscriptions: func(s *state.State) []hypp.Subscription {
			initialize := s.Signaling != nil && s.Signaling.Initialize
			return []hypp.Subscription{
				{
					Subscriber: dispatch.OnICEConnectionStateChangeSubscriber,
					Disabled:   !initialize,
				},
				{
					Subscriber: dispatch.OnConnectionStateChangeSubscriber,
					Disabled:   !initialize,
				},
				{
					Subscriber: dispatch.OnDataChannelOpenSubscriber,
					Disabled:   !initialize,
				},
				{
					Subscriber: dispatch.OnDataChannelMessageSubscriber,
					Disabled:   !initialize,
				},
			}
		},
	})
}
