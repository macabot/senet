package tale

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/state"
)

func Stick() *fairytale.Tale[*state.State] {
	return fairytale.New(
		"Stick",
		&state.State{Game: state.NewGame()},
		func(s *state.State) *hypp.VNode {
			return component.Stick(s.Game.Sticks().Flips[0])
		},
	).WithControls(
		control.NewNumberInput(
			"Flips",
			func(s *state.State, flips int) *state.State {
				s.Game.SetSticks(state.Sticks{Flips: [4]int{flips, 0, 0, 0}})
				return s
			},
			func(s *state.State) int {
				return s.Game.Sticks().Flips[0]
			},
		),
	)
}
