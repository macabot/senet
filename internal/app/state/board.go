package state

import (
	"fmt"
	"sort"

	"github.com/macabot/senet/internal/pkg/clone"
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
	Protected Icon = iota
	ReturnToStart
)

type SpecialPosition struct {
	Icon          Icon
	Protects      bool
	ReturnToStart bool
}

var SpecialPositions = map[Position]SpecialPosition{
	28: {Icon: Protected, Protects: true, ReturnToStart: false},
	27: {Icon: Protected, Protects: true, ReturnToStart: false},
	26: {Icon: ReturnToStart, Protects: false, ReturnToStart: true},
	25: {Icon: Protected, Protects: true, ReturnToStart: false},
}

var ReturnToStartPosition Position = 26

type PieceAbility int

const (
	NormalPiece PieceAbility = iota
	ProtectedPiece
	BlockingPiece
)

func (a PieceAbility) String() string {
	switch a {
	case NormalPiece:
		return "Normal"
	case ProtectedPiece:
		return "Protected"
	case BlockingPiece:
		return "Blocking"
	default:
		panic(fmt.Errorf("invalid PieceAbility %d", a))
	}
}

func (a PieceAbility) IsProtected() bool {
	return a == ProtectedPiece || a == BlockingPiece
}

func (a PieceAbility) IsBlocking() bool {
	return a == BlockingPiece
}

type Piece struct {
	ID       int
	Position Position
	Ability  PieceAbility
}

func (p *Piece) Clone() *Piece {
	if p == nil {
		return nil
	}
	c := *p
	return &c
}

type PiecesByPosition map[Position]*Piece

func NewPiecesByPosition(pieces ...*Piece) PiecesByPosition {
	m := map[Position]*Piece{}
	for _, piece := range pieces {
		m[piece.Position] = piece
	}
	return m
}

func (p PiecesByPosition) Equal(other PiecesByPosition) bool {
	if len(p) != len(other) {
		return false
	}
	for pos, piece := range p {
		if otherPiece, ok := other[pos]; !ok {
			return false
		} else if otherPiece.ID != piece.ID {
			return false
		}
	}
	return true
}

func (p PiecesByPosition) Has(pos Position) bool {
	_, ok := p[pos]
	return ok
}

func (p PiecesByPosition) OrderedByID() []*Piece {
	pieces := make([]*Piece, len(p))
	i := 0
	for _, piece := range p {
		pieces[i] = piece
		i++
	}
	sort.Sort(byID(pieces))
	return pieces
}

type byID []*Piece

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
}

func (b *Board) Clone() *Board {
	if b == nil {
		return nil
	}
	return &Board{
		PlayerPieces: [2]PiecesByPosition{
			clone.Map(b.PlayerPieces[0]),
			clone.Map(b.PlayerPieces[1]),
		},
	}
}

func (b Board) Equal(other *Board) bool {
	if other == nil {
		return false
	}
	for i, piecesByPosition := range b.PlayerPieces {
		if !piecesByPosition.Equal(other.PlayerPieces[i]) {
			return false
		}
	}
	return true
}

func NewBoard() *Board {
	return &Board{
		PlayerPieces: [2]PiecesByPosition{
			NewPiecesByPosition(
				&Piece{ID: 1, Position: 9},
				&Piece{ID: 2, Position: 7},
				&Piece{ID: 3, Position: 5},
				&Piece{ID: 4, Position: 3},
				&Piece{ID: 5, Position: 1},
			),
			NewPiecesByPosition(
				&Piece{ID: 6, Position: 8},
				&Piece{ID: 7, Position: 6},
				&Piece{ID: 8, Position: 4},
				&Piece{ID: 9, Position: 2},
				&Piece{ID: 10, Position: 0},
			),
		},
	}
}

// NeighborSquares returns the positions of the neighboring squares.
// The squares are layed out as follows:
//     29 28 27 26 25 24 23 22 21 20
//     10 11 12 13 14 15 16 17 18 19
//      9  8  7  6  5  4  3  2  1  0
//
// For example, position 12 has neighbors 27 (north), 13 (east), 7 (south) and
// 11 (west).
// Note that some position can have fewer than four neighbors. For example,
// position 0 has two neighbors: 19 (north) and 1 (west).
func (b Board) NeighborSquares(position Position) set.Set[Position] {
	neighbors := set.Set[Position]{}
	if position < 20 {
		neighbors.Add(position + 19 - 2*(position%10))
	}
	if position >= 10 {
		neighbors.Add(position - 1 - 2*(position%10))
	}
	if position > 0 {
		neighbors.Add(position - 1)
	}
	if position < 30 {
		neighbors.Add(position + 1)
	}
	return neighbors
}

func (b Board) FindGroups(piecesByPosition PiecesByPosition) map[Position]set.Set[Position] {
	groups := map[Position]set.Set[Position]{}
	for pos, piece := range piecesByPosition {
		neighborSquares := b.NeighborSquares(piece.Position)
		posGroup := set.New(pos)
		for neighborPos := range neighborSquares {
			if neighborGroup, ok := groups[neighborPos]; ok {
				posGroup.AddSet(neighborGroup)
				groups[pos] = posGroup
			}
		}
		for pos := range posGroup {
			groups[pos] = posGroup
		}
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
			return piece
		}
	}
	return nil
}

// UpdatePieceAbilities updates the abilities of the pieces given the layout of the board.
func (b *Board) UpdatePieceAbilities() {
	update := func(piecesByPosition PiecesByPosition) {
		groups := b.FindGroups(piecesByPosition)
		for pos, piece := range piecesByPosition {
			if len(groups[pos]) >= 3 {
				piece.Ability = BlockingPiece
			} else if len(groups[pos]) == 2 {
				piece.Ability = ProtectedPiece
			} else if SpecialPositions[pos].Protects {
				piece.Ability = ProtectedPiece
			} else {
				piece.Ability = NormalPiece
			}
		}
	}
	for _, piecesByPosition := range b.PlayerPieces {
		update(piecesByPosition)
	}
}
