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
	Two Icon = iota
	Three
	Cross
	Ankh
)

type SpecialPosition struct {
	Icon     Icon
	Protects bool
}

var SpecialPositions = map[Position]SpecialPosition{
	{0, 1}: {Icon: Two, Protects: true},
	{0, 2}: {Icon: Three, Protects: true},
	{0, 3}: {Icon: Cross, Protects: false},
	{0, 4}: {Icon: Ankh, Protects: true},
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

func (b Board) Neighbours(position Position) []Position {
	var neighbours []Position
	if position[0] > 0 {
		neighbours = append(neighbours, Position{position[0] - 1, position[1]})
	}
	if position[0] < 2 {
		neighbours = append(neighbours, Position{position[0] + 1, position[1]})
	}
	if position[1] > 0 {
		neighbours = append(neighbours, Position{position[0], position[1] - 1})
	}
	if position[1] < 2 {
		neighbours = append(neighbours, Position{position[0], position[1] + 1})
	}
	return neighbours
}

func (b Board) IsProtected(position Position) bool {
	if special, ok := SpecialPositions[position]; ok && special.Protects {
		return true
	}
	var pieces set.Set[Position]
	if b.Pieces[0].Has(position) {
		pieces = b.Pieces[0]
	} else if b.Pieces[1].Has(position) {
		pieces = b.Pieces[1]
	} else {
		return false
	}
	neighbours := b.Neighbours(position)
	for _, neighbour := range neighbours {
		if pieces.Has(neighbour) {
			return true
		}
	}
	return false
}
