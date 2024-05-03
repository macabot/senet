package dispatch

import (
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

func ToggleMenu(s *state.State, _ hypp.Payload) hypp.Dispatchable {
	newState := s.Clone()
	newState.ShowMenu = !newState.ShowMenu
	return newState
}

func ToggleOrientationTip(s *state.State, _ hypp.Payload) hypp.Dispatchable {
	newState := s.Clone()
	newState.HideOrientationTip = !newState.HideOrientationTip
	return newState
}
