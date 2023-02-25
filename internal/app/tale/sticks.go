package tale

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/state"
)

func Sticks() *fairytale.Tale[*state.State] {
	return fairytale.New(
		"Sticks",
		&state.State{Game: state.NewGame()},
		func(s *state.State) *hypp.VNode {
			return component.Sticks(s.Game.Sticks())
		},
	).WithControls(
		control.NewButton(
			"Throw",
			func(s *state.State) *state.State {
				s.Game.SetSticks(s.Game.Sticks().Throw())
				return s
			},
		),
		control.NewCheckbox(
			"Has thrown",
			func(s *state.State, hasThrown bool) *state.State {
				sticks := s.Game.Sticks()
				sticks.HasThrown = hasThrown
				s.Game.SetSticks(sticks)
				return s
			},
			func(s *state.State) bool {
				return s.Game.Sticks().HasThrown
			},
		),
	)
}
