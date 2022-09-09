package model_test

import (
	"testing"

	"github.com/macabot/senet/internal/app/model"
	"github.com/macabot/senet/internal/pkg/set"
	"github.com/stretchr/testify/assert"
)

func TestPositionMove(t *testing.T) {
	assert.Equal(t, model.Position{0, 0}, model.Position{0, 0}.Move(0))
	assert.Equal(t, model.Position{2, 3}, model.Position{2, 4}.Move(1))
	assert.Equal(t, model.Position{2, 5}, model.Position{2, 4}.Move(-1))
	assert.Equal(t, model.Position{1, 0}, model.Position{2, 4}.Move(5))
	assert.Equal(t, model.Position{2, 4}, model.Position{1, 0}.Move(-5))
	assert.Equal(t, model.Position{0, 9}, model.Position{2, 3}.Move(14))
}

func TestBoardNeighbours(t *testing.T) {
	assert.True(
		t,
		set.New(
			model.Position{0, 4},
			model.Position{2, 4},
			model.Position{1, 3},
			model.Position{1, 5},
		).Equal(model.Board{}.Neighbours(model.Position{1, 4})),
	)
	assert.True(
		t,
		set.New(
			model.Position{0, 9},
			model.Position{2, 9},
			model.Position{1, 8},
		).Equal(model.Board{}.Neighbours(model.Position{1, 9})),
	)
	assert.True(
		t,
		set.New(
			model.Position{1, 9},
			model.Position{2, 8},
		).Equal(model.Board{}.Neighbours(model.Position{2, 9})),
	)
}
