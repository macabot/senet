package model

import (
	"github.com/macabot/senet/internal/pkg/set"
	"github.com/macabot/senet/internal/pkg/stack"
)

type Position [2]int // [Row, Column]

func (p Position) Move(steps int) Position {
	var sign int
	if steps >= 0 {
		sign = 1
	} else {
		sign = -1
	}
	position := p
	for ; steps*sign > 0; steps -= sign {
		position = position.step(sign)
	}
	return position
}

// step returns the next Position if sign=1 or the previous position
// if sign=-1.
// The behaviour is undefined if sign has a value other than 1 or -1.
func (p Position) step(sign int) Position {
	row, column := p[0], p[1]
	if row%2 == 0 {
		column -= sign
	} else {
		column += sign
	}
	if column < 0 {
		row -= sign
		column = 0
	} else if column > 9 {
		row -= sign
		column = 9
	}
	return Position{row, column}
}

func (p Position) Next() Position {
	return p.step(1)
}

func (p Position) Previous() Position {
	return p.step(-1)
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

type Piece struct {
	ID int
}

type Pieces map[Position]Piece

func (p Pieces) Has(position Position) bool {
	_, ok := p[position]
	return ok
}

type Board struct {
	Pieces [2]Pieces
	You    int
}

func NewBoard() Board {
	return Board{
		Pieces: [2]Pieces{
			{
				Position{2, 0}: {ID: 1},
				Position{2, 2}: {ID: 2},
				Position{2, 4}: {ID: 3},
				Position{2, 6}: {ID: 4},
				Position{2, 8}: {ID: 5},
			},
			{
				Position{2, 1}: {ID: 6},
				Position{2, 3}: {ID: 7},
				Position{2, 5}: {ID: 8},
				Position{2, 7}: {ID: 9},
				Position{2, 9}: {ID: 10},
			},
		},
	}
}

func (b Board) Neighbours(position Position) set.Set[Position] {
	neighbours := set.Set[Position]{}
	if position[0] > 0 {
		neighbours.Add(Position{position[0] - 1, position[1]})
	}
	if position[0] < 2 {
		neighbours.Add(Position{position[0] + 1, position[1]})
	}
	if position[1] > 0 {
		neighbours.Add(Position{position[0], position[1] - 1})
	}
	if position[1] < 9 {
		neighbours.Add(Position{position[0], position[1] + 1})
	}
	return neighbours
}

func (b Board) piecesForPosition(position Position) Pieces {
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
	for neighbour := range neighbours {
		if _, ok := pieces[neighbour]; ok {
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
		for neighbour := range neighbours {
			if seen.Has(neighbour) {
				continue
			}
			if pieces.Has(neighbour) {
				toSee.Push(neighbour)
			}
		}
	}
	return len(seen) >= 3
}

// Move moves a piece at the given position the given number of steps.
// If no piece is on this position, nothing happens.
// If the piece is moved unto the cross, then the piece is moved instead to the
// first open square starting at [2,9].
// If the piece is moved on top of one of her own pieces, then nothing happens.
// If the piece is moved on top an opponent's piece, then the pieces swap
// places.
func (b *Board) Move(from Position, steps int) {
	var pieces Pieces
	var otherPieces Pieces
	if b.Pieces[0].Has(from) {
		pieces = b.Pieces[0]
		otherPieces = b.Pieces[1]
	} else if b.Pieces[1].Has(from) {
		pieces = b.Pieces[1]
		otherPieces = b.Pieces[0]
	} else {
		return
	}
	to := from.Move(steps)
	if to == (Position{0, 3}) {
		for to = (Position{2, 9}); b.Pieces[0].Has(to) || b.Pieces[1].Has(to); to = to.Next() {
			// no-op
		}
	}
	if pieces.Has(to) {
		return
	}
	pieces[to] = pieces[from]
	delete(pieces, from)
	if otherPieces.Has(to) {
		otherPieces[from] = otherPieces[to]
		delete(otherPieces, to)
	}
}
