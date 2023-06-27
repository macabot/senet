package control

import (
	"github.com/macabot/fairytale/control"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

type BoardConfiguration struct {
	Label string
	Board *state.Board
}

type BoardConfigurations []BoardConfiguration

func (b BoardConfigurations) SelectOptions() []control.SelectOption[int] {
	options := make([]control.SelectOption[int], len(b))
	for i, configuration := range b {
		options[i] = control.SelectOption[int]{
			Label: configuration.Label,
			Value: i,
		}
	}
	return options
}

var boardConfigurations = BoardConfigurations{
	{
		Label: "New game",
		Board: state.NewBoard(),
	},
	{
		Label: "P1 - Protecting",
		Board: &state.Board{
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
		},
	},
	{
		Label: "P2 - Protecting",
		Board: &state.Board{
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
		},
	},
	{
		Label: "P1 - Blocking",
		Board: &state.Board{
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
		},
	},
	{
		Label: "P2 - Blocking",
		Board: &state.Board{
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
		},
	},
	{
		Label: "P1 - 2nd piece up",
		Board: &state.Board{
			PlayerPieces: [2]state.PiecesByPosition{
				state.NewPiecesByPosition(
					&state.Piece{ID: 1, Position: 9},
					&state.Piece{ID: 2, Position: 12},
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
		},
	},
	{
		Label: "P2 - 2nd piece up",
		Board: &state.Board{
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
					&state.Piece{ID: 7, Position: 13},
					&state.Piece{ID: 8, Position: 4},
					&state.Piece{ID: 9, Position: 2},
					&state.Piece{ID: 10, Position: 0},
				),
			},
		},
	},
	{
		Label: "P1 - Protected by square",
		Board: &state.Board{
			PlayerPieces: [2]state.PiecesByPosition{
				state.NewPiecesByPosition(
					&state.Piece{ID: 1, Position: 9},
					&state.Piece{ID: 2, Position: 7},
					&state.Piece{ID: 3, Position: 25},
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
		},
	},
	{
		Label: "P1 - Remove other piece by moving to top row",
		Board: &state.Board{
			PlayerPieces: [2]state.PiecesByPosition{
				state.NewPiecesByPosition(
					&state.Piece{ID: 1, Position: 29},
					&state.Piece{ID: 2, Position: 28},
					&state.Piece{ID: 3, Position: 27},
					&state.Piece{ID: 4, Position: 25},
					&state.Piece{ID: 5, Position: 19},
				),
				state.NewPiecesByPosition(
					&state.Piece{ID: 6, Position: 8},
					&state.Piece{ID: 7, Position: 6},
					&state.Piece{ID: 8, Position: 4},
					&state.Piece{ID: 9, Position: 2},
					&state.Piece{ID: 10, Position: 0},
				),
			},
		},
	},
	{
		Label: "P1 - No possible moves",
		Board: &state.Board{
			PlayerPieces: [2]state.PiecesByPosition{
				state.NewPiecesByPosition(
					&state.Piece{ID: 1, Position: 4},
					&state.Piece{ID: 2, Position: 3},
					&state.Piece{ID: 3, Position: 2},
					&state.Piece{ID: 4, Position: 1},
					&state.Piece{ID: 5, Position: 0},
				),
				state.NewPiecesByPosition(
					&state.Piece{ID: 6, Position: 9},
					&state.Piece{ID: 7, Position: 8},
					&state.Piece{ID: 8, Position: 7},
					&state.Piece{ID: 9, Position: 6},
					&state.Piece{ID: 10, Position: 5},
				),
			},
		},
	},
}

func Configuration() *control.Select[*state.State, int] {
	return control.NewSelect(
		"Configuration",
		func(s *state.State, option int) hypp.Dispatchable {
			if option == -1 {
				return s
			}
			board := boardConfigurations[option].Board
			s.Game.SetBoard(board)
			return s
		},
		func(s *state.State) int {
			for i, configuration := range boardConfigurations {
				if configuration.Board.Equal(s.Game.Board) {
					return i
				}
			}
			return -1
		},
		boardConfigurations.SelectOptions(),
	)
}
