package model

import (
	"errors"

	"github.com/macabot/senet/internal/pkg/set"
	"github.com/macabot/senet/internal/pkg/stack"
)

type Position [2]int // [Row, Column]

func (p Position) Move(steps int) Position {
	row, column := p
	rowsUp := steps / 10
	remainder := steps % 10
	row -= rowsUp


	rowsUp := steps / 20
	row := p[0] - rowsUp
	remainder := steps % 20
	moveLeft := row % 2 == 0
	oppositeSteps := steps < 0
	if moveLeft == oppositeSteps {
		//    on row that moves left and steps to right
		// or on row that moves right and steps to right
		if p[1] + remainder
	} else {
		//    on row that moves left and steps to left
		// or on row that oves right and steps to left
	}
}

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

func (b Board) piecesForPosition(position Position) set.Set[Position] {
	if b.Pieces[0].Has(position) {
		return b.Pieces[0]
	} else if b.Pieces[1].Has(position) {
		return b.Pieces[1]
	} else {
		return nil
	}
}

func (b Board) IsProtected(position Position) bool {
	if special, ok := SpecialPositions[position]; ok && special.Protects {
		return true
	}
	pieces := b.piecesForPosition(position)
	if pieces == nil {
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

func (b Board) IsBlocking(position Position) bool {
	pieces := b.piecesForPosition(position)
	if pieces == nil {
		return false
	}
	toSee := stack.NewStack(position)
	seen := set.Set[Position]{}
	for toSee.Len() > 0 {
		current := toSee.Pop()
		seen.Add(current)
		neighbours := b.Neighbours(current)
		for _, neighbour := range neighbours {
			if seen.Has(neighbour) {
				continue
			}
			if pieces.Has(position) {
				toSee.Push(position)
			}
		}
	}
	return len(seen) >= 3
}

var (
	ErrPieceNotFound = errors.New("piece not found")
)

func (b *Board) Move(from Position, steps int) error {
	pieces := b.piecesForPosition(from)
	if pieces == nil {
		return ErrPieceNotFound
	}

}
