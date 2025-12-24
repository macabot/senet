package organism

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/component/organism"
	"github.com/macabot/senet/internal/app/state"
	mycontrol "github.com/macabot/senet/internal/app/tale/control"
)

func TaleBoard() *fairytale.Tale[*state.State] {
	game := state.NewGame()
	game.TurnMode = state.IsPlayer0
	return fairytale.New(
		"Board",
		&state.State{
			Game: game,
		},
		organism.Board,
	).WithControls(
		mycontrol.Configuration(),
		control.NewCheckbox(
			"Has turn",
			func(s *state.State, hasTurn bool) hypp.Dispatchable {
				if s.Game.Turn == 0 {
					s.Game.TurnMode = state.IsPlayer0
				} else if s.Game.Turn == 1 {
					s.Game.TurnMode = state.IsPlayer1
				}
				return s
			},
			func(s *state.State) bool {
				return s.Game.HasTurn()
			},
		),
		mycontrol.PlayerTurn(),
		mycontrol.Steps(),
		control.NewSelect(
			"Selected",
			func(s *state.State, id int) hypp.Dispatchable {
				if id <= 0 {
					s.Game.SetSelected(nil)
				} else {
					s.Game.SetSelected(s.Game.Board.FindPieceByID(id))
				}
				return s
			},
			func(s *state.State) int {
				if s.Game.Selected == nil {
					return 0
				}
				return s.Game.Selected.ID
			},
			[]control.SelectOption[int]{
				{Label: "Not Selected", Value: 0},
				{Label: "1", Value: 1},
				{Label: "2", Value: 2},
				{Label: "3", Value: 3},
				{Label: "4", Value: 4},
				{Label: "5", Value: 5},
				{Label: "6", Value: 6},
				{Label: "7", Value: 7},
				{Label: "8", Value: 8},
				{Label: "9", Value: 9},
				{Label: "10", Value: 10},
			},
		),
		mycontrol.ShowDirections(),
	)
}
