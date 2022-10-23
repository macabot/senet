package state

import (
	"sort"

	"github.com/macabot/senet/internal/pkg/set"
)

type Position int

type Coordinate struct {
	Row    int
	Column int
}

func (p Position) Coordinate() Coordinate {
	row := 2 - (int(p) / 10)
	column := int(p) % 10
	if row%2 == 0 {
		column = 9 - column
	}
	return Coordinate{
		Row:    row,
		Column: column,
	}
}

func PositionFromCoordinate(coordinate Coordinate) Position {
	row := coordinate.Row
	column := coordinate.Column

	p := (2 - row) * 10
	if row%2 == 0 {
		column = 9 - column
	}
	p += column

	return Position(p)
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
	Portal   bool
}

var SpecialPositions = map[Position]SpecialPosition{
	28: {Icon: Two, Protects: true, Portal: false},
	27: {Icon: Three, Protects: true, Portal: false},
	26: {Icon: Cross, Protects: false, Portal: true},
	25: {Icon: Ankh, Protects: true, Portal: false},
}

type Piece struct {
	ID       int
	Position Position
}

type PiecesByPosition map[Position]Piece

func NewPiecesByPosition(pieces ...Piece) PiecesByPosition {
	m := map[Position]Piece{}
	for _, piece := range pieces {
		m[piece.Position] = piece
	}
	return m
}

func (p PiecesByPosition) Has(pos Position) bool {
	_, ok := p[pos]
	return ok
}

func (p PiecesByPosition) OrderedByID() []Piece {
	pieces := make([]Piece, len(p))
	i := 0
	for _, piece := range p {
		pieces[i] = piece
		i++
	}
	sort.Sort(byID(pieces))
	return pieces
}

type byID []Piece

func (s byID) Len() int {
	return len(s)
}

func (s byID) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byID) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

type Board struct {
	PlayerPieces [2]PiecesByPosition
	Selected     *Piece
}

func NewBoard() Board {
	return Board{
		PlayerPieces: [2]PiecesByPosition{
			NewPiecesByPosition(
				Piece{ID: 1, Position: 9},
				Piece{ID: 2, Position: 7},
				Piece{ID: 3, Position: 5},
				Piece{ID: 4, Position: 3},
				Piece{ID: 5, Position: 1},
			),
			NewPiecesByPosition(
				Piece{ID: 6, Position: 8},
				Piece{ID: 7, Position: 6},
				Piece{ID: 8, Position: 4},
				Piece{ID: 9, Position: 2},
				Piece{ID: 10, Position: 0},
			),
		},
	}
}

// NeighourSquares returns the positions of the neighbouring squares.
// The squares are layed out as followes:
//     29 28 27 26 25 24 23 22 21 20
//     10 11 12 13 14 15 16 17 18 19
//      9  8  7  6  5  4  3  2  1  0
//
// For example, position 12 has neighbours 27 (north), 13 (east), 7 (south) and
// 11 (west).
// Note that some position can have fewer than four neighbours. For example,
// postion 0 has two neighbours: 19 (north) and 1 (west).
func (b Board) NeighbourSquares(position Position) set.Set[Position] {
	neighbours := set.Set[Position]{}
	if position < 20 {
		neighbours.Add(position + 19 - 2*(position%10))
	}
	if position >= 10 {
		neighbours.Add(position - 1 - 2*(position%10))
	}
	if position > 0 {
		neighbours.Add(position - 1)
	}
	if position < 30 {
		neighbours.Add(position + 1)
	}
	return neighbours
}

func (b Board) FindGroups(piecesByPosition PiecesByPosition) map[Position]set.Set[Position] {
	groups := map[Position]set.Set[Position]{}
	for pos, piece := range piecesByPosition {
		neighbourSquares := b.NeighbourSquares(piece.Position)
		posGroup := set.New(pos)
		for neighbourPos := range neighbourSquares {
			if neighbourGroup, ok := groups[neighbourPos]; ok {
				posGroup.AddSet(neighbourGroup)
				groups[pos] = posGroup
			}
		}
		groups[pos] = posGroup
	}

	return groups
}

func (b Board) FindPieceByID(id int) *Piece {
	var piecesByPos PiecesByPosition
	if id <= 5 {
		piecesByPos = b.PlayerPieces[0]
	} else {
		piecesByPos = b.PlayerPieces[1]
	}
	for _, piece := range piecesByPos {
		if piece.ID == id {
			return &piece
		}
	}
	return nil
}
