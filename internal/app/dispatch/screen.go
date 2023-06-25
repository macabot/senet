package dispatch

import (
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

func RotateScreenAction() hypp.Action[*state.State] {
	return func(s *state.State, _ hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		newState.Rotated = !newState.Rotated
		return newState
	}
}
