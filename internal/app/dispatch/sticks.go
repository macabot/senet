package dispatch

import (
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

func ThrowSticks(s *state.State, _ hypp.Payload) hypp.Dispatchable {
	newState := s.Clone()
	newState.Game.SetSticks(s.Game.Sticks.Throw())
	return newState
}
