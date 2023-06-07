package state

import (
	"github.com/macabot/senet/internal/pkg/set"
	"golang.org/x/exp/maps"
)

type Status int

const (
	Created Status = iota
	Ready
	InProgress
	Finished
)

type Game struct {
	Board        *Board
	Selected     *Piece
	Sticks       *Sticks
	Turn         int
	HasTurn      bool
	Status       Status
	ValidMoves   map[Position]Position
	InvalidMoves map[Position]set.Set[Position]
	HasMoved     bool
}

func (g *Game) Clone() *Game {
	if g == nil {
		return nil
	}
	return &Game{
		Board:        g.Board.Clone(),
		Selected:     g.Selected.Clone(),
		Sticks:       g.Sticks.Clone(),
		Turn:         g.Turn,
		HasTurn:      g.HasTurn,
		Status:       g.Status,
		ValidMoves:   maps.Clone(g.ValidMoves),
		InvalidMoves: maps.Clone(g.InvalidMoves),
	}
}

func NewGame() *Game {
	g := &Game{
		Board:  NewBoard(),
		Sticks: NewSticks(),
	}
	g.CalcValidMoves()
	g.HasMoved = false
	return g
}

func (g *Game) SetBoard(board *Board) {
	g.Board = board
	g.CalcValidMoves()
}

func (g *Game) SetSelected(selected *Piece) {
	g.Selected = selected
	g.CalcValidMoves()
}

func (g *Game) SetSticks(sticks *Sticks) {
	g.Sticks = sticks
	g.CalcValidMoves()
}

func (g *Game) SetTurn(turn int) {
	g.Turn = turn
	g.CalcValidMoves()
}

func (g *Game) addInvalidMove(from, to Position) {
	if _, ok := g.InvalidMoves[from]; !ok {
		g.InvalidMoves[from] = set.New[Position]()
	}
	g.InvalidMoves[from].Add(to)
}

func (g Game) CanClickOnPiece(player int, piece *Piece) bool {
	return g.PiecesDrawAttention(player) || g.PieceIsSelected(piece)
}

func (g Game) PiecesDrawAttention(player int) bool {
	return g.HasTurn &&
		g.Sticks.HasThrown &&
		player == g.Turn
}

func (g Game) PieceIsSelected(piece *Piece) bool {
	return g.Selected != nil && g.Selected.Position == piece.Position
}

func (g Game) SticksDrawAttention() bool {
	return !g.Sticks.HasThrown && g.HasTurn
}

func (g *Game) CalcValidMoves() {
	g.HasMoved = true
	g.Board.UpdatePieceAbilities()
	g.ValidMoves = map[Position]Position{}
	g.InvalidMoves = map[Position]set.Set[Position]{}

	piecesByPos := g.Board.PlayerPieces[g.Turn]
	otherPiecesByPos := g.Board.PlayerPieces[(g.Turn+1)%2]

	otherGroups := g.Board.FindGroups(otherPiecesByPos)

	findMoves := func(steps int, protectedSize int) {
		for pos := range piecesByPos {
			toPos := pos + Position(steps)
			if toPos >= 30 || toPos < 0 {
				g.addInvalidMove(pos, toPos)
				continue
			}
			if _, ok := piecesByPos[toPos]; ok {
				g.addInvalidMove(pos, toPos)
				continue
			}
			isBlocked := false
			for checkPos := pos + 1; checkPos <= toPos; checkPos++ {
				if group, ok := otherGroups[checkPos]; ok && len(group) >= 3 {
					isBlocked = true
					break
				}
			}
			if isBlocked {
				g.addInvalidMove(pos, toPos)
				continue
			}
			if group, ok := otherGroups[toPos]; ok {
				if len(group) >= protectedSize {
					g.addInvalidMove(pos, toPos)
					continue
				}
				if special, ok := SpecialPositions[toPos]; ok && special.Protects {
					g.addInvalidMove(pos, toPos)
					continue
				}
			}

			g.ValidMoves[pos] = toPos

			if special, ok := SpecialPositions[toPos]; ok && special.ReturnToStart {
				for portalPos := Position(0); portalPos < 30; portalPos++ {
					if !piecesByPos.Has(portalPos) && !otherPiecesByPos.Has(Position(portalPos)) {
						g.ValidMoves[toPos] = portalPos
						break
					}
				}
			}
		}
	}

	steps := g.Sticks.Steps()
	findMoves(steps, 2)
	if len(g.ValidMoves) == 0 {
		findMoves(-steps, 1)
	}
}

func (g Game) CanMove(player int, from Position) bool {
	piecesByPos := g.Board.PlayerPieces[player]
	if _, ok := piecesByPos[from]; !ok {
		return false
	}
	if _, ok := g.ValidMoves[from]; !ok {
		return false
	}
	return true
}

func (g *Game) Move(player int, from Position) {
	piecesByPos := g.Board.PlayerPieces[player]
	piece, ok := piecesByPos[from]
	if !ok {
		return
	}
	to, ok := g.ValidMoves[from]
	if !ok {
		return
	}
	piece.Position = to
	delete(piecesByPos, from)
	piecesByPos[to] = piece
}
