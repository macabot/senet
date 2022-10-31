package component

import (
	"github.com/macabot/fairytale/fairy"
	"github.com/macabot/senet/internal/app/view/render/component"
	"github.com/macabot/senet/internal/app/view/state"
)

func BoardTale() *fairy.Tale {
	configuration := 0
	return fairy.NewTale(
		"Board",
		&state.State{
			Game: state.NewGame(),
		},
		component.Board,
	).WithControls(
		fairy.NewSelectControl(
			"Configuration",
			func(props *state.State, option int) *state.State {
				configuration = option
				switch option {
				case 0:
					props.Game.SetBoard(state.NewBoard())
				case 1:
					props.Game.SetBoard(&state.Board{
						PlayerPieces: [2]state.PiecesByPosition{
							state.NewPiecesByPosition(
								state.Piece{ID: 1, Position: 9},
								state.Piece{ID: 2, Position: 10},
								state.Piece{ID: 3, Position: 5},
								state.Piece{ID: 4, Position: 3},
								state.Piece{ID: 5, Position: 1},
							),
							state.NewPiecesByPosition(
								state.Piece{ID: 6, Position: 8},
								state.Piece{ID: 7, Position: 6},
								state.Piece{ID: 8, Position: 4},
								state.Piece{ID: 9, Position: 2},
								state.Piece{ID: 10, Position: 0},
							),
						},
					})
				case 2:
					props.Game.SetBoard(&state.Board{
						PlayerPieces: [2]state.PiecesByPosition{
							state.NewPiecesByPosition(
								state.Piece{ID: 1, Position: 9},
								state.Piece{ID: 2, Position: 7},
								state.Piece{ID: 3, Position: 5},
								state.Piece{ID: 4, Position: 3},
								state.Piece{ID: 5, Position: 1},
							),
							state.NewPiecesByPosition(
								state.Piece{ID: 6, Position: 8},
								state.Piece{ID: 7, Position: 11},
								state.Piece{ID: 8, Position: 4},
								state.Piece{ID: 9, Position: 2},
								state.Piece{ID: 10, Position: 0},
							),
						},
					})
				case 3:
					props.Game.SetBoard(&state.Board{
						PlayerPieces: [2]state.PiecesByPosition{
							state.NewPiecesByPosition(
								state.Piece{ID: 1, Position: 9},
								state.Piece{ID: 2, Position: 10},
								state.Piece{ID: 3, Position: 11},
								state.Piece{ID: 4, Position: 3},
								state.Piece{ID: 5, Position: 1},
							),
							state.NewPiecesByPosition(
								state.Piece{ID: 6, Position: 8},
								state.Piece{ID: 7, Position: 6},
								state.Piece{ID: 8, Position: 4},
								state.Piece{ID: 9, Position: 2},
								state.Piece{ID: 10, Position: 0},
							),
						},
					})
				case 4:
					props.Game.SetBoard(&state.Board{
						PlayerPieces: [2]state.PiecesByPosition{
							state.NewPiecesByPosition(
								state.Piece{ID: 1, Position: 9},
								state.Piece{ID: 2, Position: 7},
								state.Piece{ID: 3, Position: 5},
								state.Piece{ID: 4, Position: 3},
								state.Piece{ID: 5, Position: 1},
							),
							state.NewPiecesByPosition(
								state.Piece{ID: 6, Position: 8},
								state.Piece{ID: 7, Position: 11},
								state.Piece{ID: 8, Position: 12},
								state.Piece{ID: 9, Position: 2},
								state.Piece{ID: 10, Position: 0},
							),
						},
					})
				}
				return props
			},
			func(_ *state.State) int {
				return configuration
			},
			[]fairy.SelectOption[int]{
				{Label: "New game", Value: 0},
				{Label: "P1- Protecting", Value: 1},
				{Label: "P2 - Protecting", Value: 2},
				{Label: "P1 - Blocking", Value: 3},
				{Label: "P2 - Blocking", Value: 4},
			},
		),
		fairy.NewCheckboxControl(
			"Has turn",
			func(props *state.State, hasTurn bool) *state.State {
				props.Game.SetHasTurn(hasTurn)
				return props
			},
			func(props *state.State) bool {
				return props.Game.HasTurn()
			},
		),
		fairy.NewSelectControl(
			"Turn",
			func(props *state.State, turn int) *state.State {
				props.Game.SetTurn(turn)
				return props
			},
			func(props *state.State) int {
				return props.Game.Turn()
			},
			[]fairy.SelectOption[int]{
				{Label: "Player 1", Value: 0},
				{Label: "Player 2", Value: 1},
			},
		),
		fairy.NewSelectControl(
			"Steps",
			func(props *state.State, steps int) *state.State {
				props.Game.SetSticks(state.SticksFromSteps(steps, steps != 0))
				return props
			},
			func(props *state.State) int {
				if !props.Game.Sticks().HasThrown {
					return 0
				}
				return props.Game.Sticks().Steps()
			},
			[]fairy.SelectOption[int]{
				{Label: "Not thrown", Value: 0},
				{Label: "1", Value: 1},
				{Label: "2", Value: 2},
				{Label: "3", Value: 3},
				{Label: "4", Value: 4},
				{Label: "6", Value: 6},
			},
		),
		fairy.NewSelectControl(
			"Selected",
			func(props *state.State, id int) *state.State {
				if id <= 0 {
					props.Game.Board().Selected = nil
				} else {
					props.Game.Board().Selected = props.Game.Board().FindPieceByID(id)
				}
				return props
			},
			func(props *state.State) int {
				if props.Game.Board().Selected == nil {
					return 0
				}
				return props.Game.Board().Selected.ID
			},
			[]fairy.SelectOption[int]{
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
