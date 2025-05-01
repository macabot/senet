package dispatch

import (
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

var onThrowSticks = []func(s, newState *state.State) []hypp.Effect{}

func ThrowSticks(s *state.State, _ hypp.Payload) hypp.Dispatchable {
	newState := s.Clone()
	newState.Game.ThrowSticks(newState)
	var effects []hypp.Effect
	for _, f := range onThrowSticks {
		fEffects := f(s, newState)
		effects = append(effects, fEffects...)
	}
	return hypp.StateAndEffects[*state.State]{
		State:   newState,
		Effects: effects,
	}
}
