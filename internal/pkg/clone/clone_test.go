package clone_test

import (
	"testing"

	"github.com/macabot/senet/internal/app/state"
	"github.com/macabot/senet/internal/pkg/clone"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	pos := state.Position(1)
	piece := &state.Piece{}
	byPos := state.PiecesByPosition{
		pos: piece,
	}
	byPosClone := clone.Map(byPos)
	assert.False(t, piece == byPosClone[pos])
}
