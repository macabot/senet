package dispatch

import (
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

var onThrowSticks = []func(s, newState *state.State){}

func ThrowSticks(s *state.State, _ hypp.Payload) hypp.Dispatchable {
	newState := s.Clone()
	newState.Game.ThrowSticks()
	for _, f := range onThrowSticks {
		f(s, newState)
	}
	return newState
}
