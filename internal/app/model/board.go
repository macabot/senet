package model

import "github.com/macabot/senet/internal/pkg/set"

type Position [2]int

type Board struct {
	Pieces      [2]set.Set[Position]
	You         int
	Selected    *Position
	Highlighted set.Set[Position]
}

type Icon int

const (
	Foo Icon = iota // TODO rename
	Bar
	Pit
	Buz
)

var SpecialPositions = map[Position]Icon{
	{0, 1}: Foo,
	{0, 2}: Bar,
	{0, 3}: Pit,
	{0, 4}: Buz,
}

func NewBoard() Board {
	return Board{
		Pieces: [2]set.Set[Position]{
			set.New(
				Position{2, 0},
				Position{2, 2},
				Position{2, 4},
				Position{2, 6},
				Position{2, 8},
			),
			set.New(
				Position{2, 1},
				Position{2, 3},
				Position{2, 5},
				Position{2, 7},
				Position{2, 9},
			),
		},
	}
}
