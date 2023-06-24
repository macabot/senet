package tale

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/state"
	mycontrol "github.com/macabot/senet/internal/app/tale/control"
)

func Sticks() *fairytale.Tale[*state.State] {
	game := state.NewGame()
	game.HasTurn = true
	return fairytale.New(
		"Sticks",
		&state.State{Game: game},
		func(s *state.State) *hypp.VNode {
			return component.Sticks(component.SticksProps{
				Sticks:        s.Game.Sticks,
				DrawAttention: s.Game.SticksDrawAttention(),
			})
		},
	).WithControls(
		mycontrol.Steps(),
	)
}
