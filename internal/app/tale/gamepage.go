package tale

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/state"
	mycontrol "github.com/macabot/senet/internal/app/tale/control"
)

func GamePage() *fairytale.Tale[*state.State] {
	game := state.NewGame()
	game.TurnMode = state.IsPlayer1
	return fairytale.New(
		"GamePage",
		&state.State{
			Game: game,
			Page: state.GamePage,
		},
		component.Senet,
	).WithControls(
		mycontrol.Configuration(),
		mycontrol.Steps(),
		mycontrol.PlayerTurn(),
		control.NewSelect(
			"Winner",
			func(s *state.State, option int) hypp.Dispatchable {
				if option == 0 {
					s.Game.Winner = nil
				} else if option == 1 {
					player := 0
					s.Game.Winner = &player
				} else {
					player := 1
					s.Game.Winner = &player
				}
				return s
			},
			func(s *state.State) int {
				if s.Game == nil {
					return -1
				} else if s.Game.Winner == nil {
					return 0
				} else if *s.Game.Winner == 0 {
					return 1
				} else {
					return 2
				}
			},
			[]control.SelectOption[int]{
				{Label: "No winner", Value: 0},
				{Label: "Player 1", Value: 1},
				{Label: "Player 2", Value: 2},
			},
		),
	).WithSettings(fairytale.TaleSettings{
		Target: fairytale.TaleAsHTML,
	})
}
