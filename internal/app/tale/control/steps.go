package control

import (
	"github.com/macabot/fairytale/control"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

func Steps() *control.Select[*state.State, int] {
	return control.NewSelect(
		"Steps",
		func(s *state.State, steps int) hypp.Dispatchable {
			if steps == 0 {
				s.Game.Sticks.HasThrown = false
			} else {
				s.Game.Sticks.SetSteps(steps)
			}
			return s
		},
		func(s *state.State) int {
			if s.Game == nil {
				return -1
			}
			if !s.Game.Sticks.HasThrown {
				return 0
			}
			return s.Game.Sticks.Steps()
		},
		[]control.SelectOption[int]{
			{Label: "Not thrown", Value: 0},
			{Label: "1", Value: 1},
			{Label: "2", Value: 2},
			{Label: "3", Value: 3},
			{Label: "4", Value: 4},
			{Label: "6", Value: 6},
		},
	)
}
