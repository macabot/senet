package app

import (
	"encoding/json"
	"fmt"
	"runtime/debug"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/window"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
	"github.com/macabot/senet/internal/pkg/localstorage"
)

func loadState() *state.State {
	defaultState := &state.State{
		Page: state.StartPage,
	}
	v, ok := localstorage.GetItem("state")
	if !ok {
		return defaultState
	}
	s := &state.State{}
	if err := json.Unmarshal([]byte(v), s); err != nil {
		window.Console().Error("Could not JSON decode state from localstorage.\nResetting to default state.")
		return defaultState
	}
	if s.PanicTrace != nil {
		return defaultState
	}

	if s.Game != nil && s.Game.Sticks.GeneratorKind == state.TutorialSticksGeneratorKind {
		dispatch.RegisterTutorial()
	}

	if s.Signaling != nil && s.Signaling.Initialized {
		s.Signaling.Initialized = false
		s.Signaling.ICEConnectionState = "disconnected"
		s.Signaling.ConnectionState = "failed"
		s.Signaling.ReadyState = "closed"
	}

	return s
}

func dispatchWrapper(dispatchFunc hypp.Dispatch) hypp.Dispatch {
	return func(dispatchable hypp.Dispatchable, payload hypp.Payload) {
		var s *state.State
		stateFound := false
		switch v := dispatchable.(type) {
		case hypp.StateAndEffects[*state.State]:
			s = v.State
			stateFound = true
			wrappedEffects := make([]hypp.Effect, len(v.Effects))
			for i, e := range v.Effects {
				wrappedEffects[i] = hypp.Effect{
					Effecter: func(dispatchFunc hypp.Dispatch, payload hypp.Payload) {
						defer dispatch.RecoverPanic(dispatchFunc)
						e.Effecter(dispatchFunc, payload)
					},
					Payload: e.Payload,
				}
			}
			dispatchable = hypp.StateAndEffects[*state.State]{
				State:   v.State,
				Effects: wrappedEffects,
			}
		case *state.State:
			s = v
			stateFound = true
		}

		dispatchFunc(dispatchable, payload)

		if stateFound {
			b, err := json.Marshal(s)
			if err != nil {
				panic(err)
			}
			localstorage.SetItem("state", string(b))
		}
	}
}

func Run(element window.Element) {
	hypp.App(hypp.AppProps[*state.State]{
		Init: loadState(),
		View: func(s *state.State) (out *hypp.VNode) {
			defer func() {
				if r := recover(); r != nil {
					panicTrace := fmt.Sprintf("%v\n%s", r, string(debug.Stack()))
					window.Console().Error(panicTrace)
					s.PanicTrace = &panicTrace
					out = component.Senet(s)
				}
			}()

			return component.Senet(s)
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
