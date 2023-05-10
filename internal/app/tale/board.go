package tale

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/state"
)

func Board() *fairytale.Tale[*state.State] {
	configuration := 0
	return fairytale.New(
		"Board",
		&state.State{
			Game: state.NewGame(),
		},
		component.Board,
	).WithControls(
		control.NewSelect(
			"Configuration",
			func(s *state.State, option int) hypp.Dispatchable {
				configuration = option
				switch option {
				case 0:
					s.Game.SetBoard(state.NewBoard())
				case 1:
					s.Game.SetBoard(&state.Board{
						PlayerPieces: [2]state.PiecesByPosition{
							state.NewPiecesByPosition(
								&state.Piece{ID: 1, Position: 9},
								&state.Piece{ID: 2, Position: 10},
								&state.Piece{ID: 3, Position: 5},
								&state.Piece{ID: 4, Position: 3},
								&state.Piece{ID: 5, Position: 1},
							),
							state.NewPiecesByPosition(
								&state.Piece{ID: 6, Position: 8},
								&state.Piece{ID: 7, Position: 6},
								&state.Piece{ID: 8, Position: 4},
								&state.Piece{ID: 9, Position: 2},
								&state.Piece{ID: 10, Position: 0},
							),
						},
					})
				case 2:
					s.Game.SetBoard(&state.Board{
						PlayerPieces: [2]state.PiecesByPosition{
							state.NewPiecesByPosition(
								&state.Piece{ID: 1, Position: 9},
								&state.Piece{ID: 2, Position: 7},
								&state.Piece{ID: 3, Position: 5},
								&state.Piece{ID: 4, Position: 3},
								&state.Piece{ID: 5, Position: 1},
							),
							state.NewPiecesByPosition(
								&state.Piece{ID: 6, Position: 8},
								&state.Piece{ID: 7, Position: 11},
								&state.Piece{ID: 8, Position: 4},
								&state.Piece{ID: 9, Position: 2},
								&state.Piece{ID: 10, Position: 0},
							),
						},
					})
				case 3:
					s.Game.SetBoard(&state.Board{
						PlayerPieces: [2]state.PiecesByPosition{
							state.NewPiecesByPosition(
								&state.Piece{ID: 1, Position: 9},
								&state.Piece{ID: 2, Position: 10},
								&state.Piece{ID: 3, Position: 11},
								&state.Piece{ID: 4, Position: 3},
								&state.Piece{ID: 5, Position: 1},
							),
							state.NewPiecesByPosition(
								&state.Piece{ID: 6, Position: 8},
								&state.Piece{ID: 7, Position: 6},
								&state.Piece{ID: 8, Position: 4},
								&state.Piece{ID: 9, Position: 2},
								&state.Piece{ID: 10, Position: 0},
							),
						},
					})
				case 4:
					s.Game.SetBoard(&state.Board{
						PlayerPieces: [2]state.PiecesByPosition{
							state.NewPiecesByPosition(
								&state.Piece{ID: 1, Position: 9},
								&state.Piece{ID: 2, Position: 7},
								&state.Piece{ID: 3, Position: 5},
								&state.Piece{ID: 4, Position: 3},
								&state.Piece{ID: 5, Position: 1},
							),
							state.NewPiecesByPosition(
								&state.Piece{ID: 6, Position: 8},
								&state.Piece{ID: 7, Position: 11},
								&state.Piece{ID: 8, Position: 12},
								&state.Piece{ID: 9, Position: 2},
								&state.Piece{ID: 10, Position: 0},
							),
						},
					})
				}
				return s
			},
			func(_ *state.State) int {
				return configuration
			},
			[]control.SelectOption[int]{
				{Label: "New game", Value: 0},
				{Label: "P1- Protecting", Value: 1},
				{Label: "P2 - Protecting", Value: 2},
				{Label: "P1 - Blocking", Value: 3},
				{Label: "P2 - Blocking", Value: 4},
			},
		),
		control.NewCheckbox(
			"Has turn",
			func(s *state.State, hasTurn bool) *state.State {
				s.Game.SetHasTurn(hasTurn)
				return s
			},
			func(s *state.State) bool {
				return s.Game.HasTurn
			},
		),
		control.NewSelect(
			"Turn",
			func(s *state.State, turn int) hypp.Dispatchable {
				s.Game.SetTurn(turn)
				return s
			},
			func(s *state.State) int {
				return s.Game.Turn
			},
			[]control.SelectOption[int]{
				{Label: "Player 1", Value: 0},
				{Label: "Player 2", Value: 1},
			},
		),
		control.NewSelect(
			"Steps",
			func(s *state.State, steps int) hypp.Dispatchable {
				s.Game.SetSticks(state.SticksFromSteps(steps, steps != 0))
				return s
			},
			func(s *state.State) int {
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
		),
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
	)
}
