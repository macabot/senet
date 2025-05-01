package state_test

import (
	"testing"

	"github.com/macabot/senet/internal/app/state"
	"github.com/stretchr/testify/assert"
)

func TestNoMovePossible(t *testing.T) {
	game := state.NewGame()
	game.SetBoard(&state.Board{
		PlayerPieces: [2]state.PiecesByPosition{
			state.NewPiecesByPosition(
				&state.Piece{ID: 1, Position: 10},
				&state.Piece{ID: 2, Position: 9},
				&state.Piece{ID: 3, Position: 30},
				&state.Piece{ID: 4, Position: 31},
				&state.Piece{ID: 5, Position: 32},
			),
			state.NewPiecesByPosition(
				&state.Piece{ID: 6, Position: 27},
				&state.Piece{ID: 7, Position: 12},
				&state.Piece{ID: 8, Position: 7},
				&state.Piece{ID: 9, Position: 33},
				&state.Piece{ID: 10, Position: 34},
			),
		},
	})
	sticks := state.NewSticks()
	sticks.SetSteps(3)
	game.SetSticks(sticks)
	assert.Empty(t, game.ValidMoves)
}
