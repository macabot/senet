package tale

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
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
				NoValidMoves:  len(s.Game.ValidMoves) == 0,
			})
		},
	).WithControls(
		mycontrol.Steps(),
		control.NewCheckbox(
			"No valid moves",
			func(s *state.State, noMoves bool) *state.State {
				if noMoves {
					s.Game.SetBoard(mycontrol.NoValidMovesBoard)
				} else {
					s.Game.SetBoard(state.NewBoard())
				}
				return s
			},
			func(s *state.State) bool {
				return len(s.Game.ValidMoves) == 0
			},
		),
	)
}
