package dispatch

import (
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

func ToggleMenuAction() hypp.Action[*state.State] {
	return func(s *state.State, _ hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		newState.ShowMenu = !newState.ShowMenu
		return newState
	}
}

func ToggleOrientationTipAction() hypp.Action[*state.State] {
	return func(s *state.State, _ hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		newState.HideOrientationTip = !newState.HideOrientationTip
		return newState
	}
}
