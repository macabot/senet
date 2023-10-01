package control

import (
	"github.com/macabot/fairytale/control"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

var NoValidMovesBoard = &state.Board{
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
}

var boardConfigurations = LabeledSlice[*state.Board]{
	{
		Label: "New game",
		V:     state.NewBoard(),
	},
	{
		Label: "P1 - Protecting",
		V: &state.Board{
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
		V: &state.Board{
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
		V: &state.Board{
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
		V: &state.Board{
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
		V: &state.Board{
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
		V: &state.Board{
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
		V: &state.Board{
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
		V: &state.Board{
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
		Label: "P1 - No valid moves",
		V:     NoValidMovesBoard,
	},
	{
		Label: "No pieces",
		V:     &state.Board{},
	},
}

func Configuration() *control.Select[*state.State, int] {
	return control.NewSelect(
		"Configuration",
		func(s *state.State, option int) hypp.Dispatchable {
			if option == -1 {
				return s
			}
			optionBoard := boardConfigurations[option].V
			board := &state.Board{
				PlayerPieces:   optionBoard.PlayerPieces,
				ShowDirections: s.Game.Board.ShowDirections,
			}
			s.Game.SetBoard(board)
			return s
		},
		func(s *state.State) int {
			for i, configuration := range boardConfigurations {
				if s.Game != nil && configuration.V.Equal(s.Game.Board) {
					return i
				}
			}
			return -1
		},
		boardConfigurations.SelectOptions(),
	)
}
