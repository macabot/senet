package state_test

import (
	"testing"

	"github.com/macabot/senet/internal/app/state"
	"github.com/macabot/senet/internal/pkg/set"
	"github.com/stretchr/testify/assert"
)

func TestBoardNeighbourSquares(t *testing.T) {
	assert.Equal(
		t,
		set.New[state.Position](
			25,
			5,
			15,
			13,
		),
		state.Board{}.NeighbourSquares(14),
	)
	assert.Equal(
		t,
		set.New[state.Position](
			23,
			3,
			17,
			15,
		),
		state.Board{}.NeighbourSquares(16),
	)
	assert.Equal(
		t,
		set.New[state.Position](
			0,
			20,
			18,
		),
		state.Board{}.NeighbourSquares(19),
	)
	assert.Equal(
		t,
		set.New[state.Position](
			19,
			21,
		),
		state.Board{}.NeighbourSquares(20),
	)
}

func TestBoardFindGroups(t *testing.T) {
	assert.Equal(
		t,
		map[state.Position]set.Set[state.Position]{
			9:  set.New[state.Position](9, 10),
			10: set.New[state.Position](9, 10),
		},
		state.Board{}.FindGroups(state.PiecesByPosition{
			9:  &state.Piece{ID: 1, Position: 9},
			10: &state.Piece{ID: 2, Position: 10},
		}),
	)
}

func TestUpdatePieceAbilitiesSetProtected(t *testing.T) {
	board := state.Board{
		PlayerPieces: [2]state.PiecesByPosition{
			state.NewPiecesByPosition(
				&state.Piece{ID: 1, Position: 7},
				&state.Piece{ID: 2, Position: 12},
			),
		},
	}
	board.UpdatePieceAbilities()
	assert.True(t, board.PlayerPieces[0][7].Ability.IsProtected())
	assert.True(t, board.PlayerPieces[0][12].Ability.IsProtected())
}

func TestUpdatePieceAbilitiesSetBlocking(t *testing.T) {
	board := state.Board{
		PlayerPieces: [2]state.PiecesByPosition{
			state.NewPiecesByPosition(
				&state.Piece{ID: 1, Position: 7},
				&state.Piece{ID: 2, Position: 12},
				&state.Piece{ID: 3, Position: 27},
			),
		},
	}
	board.UpdatePieceAbilities()
	assert.True(t, board.PlayerPieces[0][7].Ability.IsBlocking())
	assert.True(t, board.PlayerPieces[0][12].Ability.IsBlocking())
	assert.True(t, board.PlayerPieces[0][27].Ability.IsBlocking())
}

func TestUpdatePieceAbilitiesSetNormal(t *testing.T) {
	board := state.Board{
		PlayerPieces: [2]state.PiecesByPosition{
			state.NewPiecesByPosition(
				&state.Piece{ID: 1, Position: 7, Ability: state.ProtectedPiece},
			),
		},
	}
	board.UpdatePieceAbilities()
	assert.Equal(t, state.NormalPiece, board.PlayerPieces[0][7].Ability)
}
