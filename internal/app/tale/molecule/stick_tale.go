package molecule

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/component/molecule"
	"github.com/macabot/senet/internal/app/state"
)

func TaleStick() *fairytale.Tale[*state.State] {
	return fairytale.New(
		"Stick",
		&state.State{Game: state.NewGame()},
		func(s *state.State) *hypp.VNode {
			return molecule.Stick(s.Game.Sticks.Flips[0])
		},
	).WithControls(
		control.NewNumberInput(
			"Flips",
			func(s *state.State, flips int) hypp.Dispatchable {
				s.Game.SetSticks(&state.Sticks{Flips: [4]int{flips, 0, 0, 0}})
				return s
			},
			func(s *state.State) int {
				return s.Game.Sticks.Flips[0]
			},
		),
	)
}
