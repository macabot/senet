package state

import (
	"fmt"

	"github.com/macabot/senet/internal/pkg/set"
	"golang.org/x/exp/maps"
)

type TurnMode int

const (
	IsBothPlayers TurnMode = iota
	IsPlayer0
	IsPlayer1
)

type Game struct {
	Players               [2]*Player
	Board                 *Board
	Selected              *Piece
	SelectedChangeCounter int
	Sticks                *Sticks
	Turn                  int
	TurnMode              TurnMode
	OverwriteHasTurn      *bool
	ValidMoves            map[Position]Position
	InvalidMoves          map[Position]set.Set[Position]
	HasMoved              bool
	Winner                *int
}

func (g *Game) Clone() *Game {
	if g == nil {
		return nil
	}
	var winnerClone *int
	if g.Winner != nil {
		w := *g.Winner
		winnerClone = &w
	}
	return &Game{
		Players: [2]*Player{
			g.Players[0].Clone(),
			g.Players[1].Clone(),
		},
		Board:                 g.Board.Clone(),
		Selected:              g.Selected.Clone(),
		SelectedChangeCounter: g.SelectedChangeCounter,
		Sticks:                g.Sticks.Clone(),
		Turn:                  g.Turn,
		TurnMode:              g.TurnMode,
		OverwriteHasTurn:      g.OverwriteHasTurn,
		ValidMoves:            maps.Clone(g.ValidMoves),
		InvalidMoves:          maps.Clone(g.InvalidMoves),
		HasMoved:              g.HasMoved,
		Winner:                winnerClone,
	}
}

func NewGame() *Game {
	g := &Game{
		Players: [2]*Player{{Name: "Player 1"}, {Name: "Player 2"}},
		Board:   NewBoard(),
		Sticks:  NewSticks(),
	}
	g.CalcValidMoves()
	g.HasMoved = false
	return g
}

func (g Game) HasTurn() bool {
	if g.OverwriteHasTurn != nil {
		return *g.OverwriteHasTurn
	}
	if g.Winner != nil {
		return false
	}
	switch g.TurnMode {
	case IsBothPlayers:
		return true
	case IsPlayer0:
		return g.Turn == 0
	case IsPlayer1:
		return g.Turn == 1
	default:
		panic(fmt.Errorf("Invalid TurnMode %v", g.TurnMode))
	}
}

func (g *Game) SetBoard(board *Board) {
	g.Board = board
	g.CalcValidMoves()
}

func (g *Game) SetSelected(selected *Piece) {
	nilToNil := selected == nil && g.Selected == nil
	sameID := selected != nil && g.Selected != nil && selected.ID == g.Selected.ID
	if !nilToNil && !sameID {
		g.SelectedChangeCounter++
	}
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
	return g.PieceDrawsAttention(player, piece.Position) || g.PieceIsSelected(piece)
}

func (g Game) PieceDrawsAttention(player int, pos Position) bool {
	return g.HasTurn() &&
		g.Sticks.HasThrown &&
		player == g.Turn &&
		pos >= 0 && pos < 30
}

func (g Game) PieceIsSelected(piece *Piece) bool {
	return g.Selected != nil && g.Selected.Position == piece.Position
}

func (g Game) CanThrow() bool {
	return !g.Sticks.HasThrown && g.HasTurn()
}

func (g Game) StartPosition() Position {
	var pos Position
	piecesByPos := g.Board.PlayerPieces[g.Turn]
	otherPiecesByPos := g.Board.PlayerPieces[(g.Turn+1)%2]
	for ; piecesByPos.Has(pos) || otherPiecesByPos.Has(pos); pos++ {
		// no-op
	}
	return pos
}

func (g Game) AllOnTopRow(player int) bool {
	piecesByPos := g.Board.PlayerPieces[player]
	for pos := range piecesByPos {
		if pos < 20 {
			return false
		}
	}
	return true
}

func (g Game) NextPositionOffBoard() Position {
	alreadyOffBoard := 0
	for _, piecesByPos := range g.Board.PlayerPieces {
		for pos := range piecesByPos {
			if pos >= 30 {
				alreadyOffBoard++
			}
		}
	}
	return Position(30 + alreadyOffBoard)
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
			if pos >= 30 {
				continue
			}
			if pos == ReturnToStartPosition {
				g.ValidMoves[pos] = g.StartPosition()
				continue
			}
			if pos == MoveOffBoardPosition && g.AllOnTopRow(g.Turn) {
				g.ValidMoves[pos] = g.NextPositionOffBoard()
				continue
			}

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
			loopStep := Position(1)
			loopCheck := func(checkPos, toPos Position) bool {
				return checkPos <= toPos
			}
			if toPos < pos {
				loopStep = -1
				loopCheck = func(checkPos, toPos Position) bool {
					return checkPos >= toPos
				}
			}
			for checkPos := pos + loopStep; loopCheck(checkPos, toPos); checkPos += loopStep {
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
				startPos := g.Board.StartPosition()
				g.ValidMoves[toPos] = startPos
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

type NextMove struct {
	Player int
	From   Position
	To     Position
}

func (g Game) FindValidFromPosition(toPosition Position) (Position, bool) {
	fromPositionFound := false
	var fromPosition Position
	for from, to := range g.ValidMoves {
		if to == toPosition {
			fromPosition = from
			fromPositionFound = true
		}
	}
	return fromPosition, fromPositionFound
}

func (g *Game) Move(player int, from, to Position) (*NextMove, error) {
	piecesByPos := g.Board.PlayerPieces[player]
	piece, ok := piecesByPos[from]
	if !ok {
		return nil, fmt.Errorf("Cannot move. Piece not found on '%d' for player '%d'", from, player)
	}

	if validMovesTo, ok := g.ValidMoves[from]; !ok || to != validMovesTo {
		return nil, fmt.Errorf("Cannot move. Move is not valid from '%d' to '%d' for player '%d'", from, to, player)
	}
	piece.Position = to
	delete(piecesByPos, from)
	piecesByPos[to] = piece

	otherPiecesByPos := g.Board.PlayerPieces[(player+1)%2]
	if otherPiece, ok := otherPiecesByPos[to]; ok {
		otherPiece.Position = from
		delete(otherPiecesByPos, to)
		otherPiecesByPos[from] = otherPiece
	}

	g.Selected = nil

	var nextMove *NextMove
	if SpecialPositions[to].ReturnToStart {
		nextMove = &NextMove{
			Player: player,
			From:   to,
			To:     g.StartPosition(),
		}
	} else if to >= 20 && to < 30 && piecesByPos.Has(MoveOffBoardPosition) && g.AllOnTopRow(player) {
		nextMove = &NextMove{
			Player: player,
			From:   MoveOffBoardPosition,
			To:     g.NextPositionOffBoard(),
		}
	}

	if nextMove == nil && !g.Sticks.CanGoAgain() {
		g.Turn = (g.Turn + 1) % 2
	}
	g.Sticks.HasThrown = false

	g.UpdateWinner()
	g.CalcValidMoves()

	return nextMove, nil
}

func (g *Game) UpdateWinner() {
	for i, piecesByPos := range g.Board.PlayerPieces {
		allOffBoard := true
		for pos := range piecesByPos {
			if pos < 30 {
				allOffBoard = false
				break
			}
		}
		if allOffBoard {
			player := i
			g.Winner = &player
			return
		}
	}
	g.Winner = nil
}

func (g *Game) NoMove(player int) error {
	if len(g.ValidMoves) > 0 {
		return fmt.Errorf("Cannot perform no-move. There are valid moves.")
	}

	g.Selected = nil
	if !g.Sticks.CanGoAgain() {
		g.Turn = (g.Turn + 1) % 2
	}
	g.Sticks.HasThrown = false
	g.CalcValidMoves()

	return nil
}

func (g *Game) ThrowSticks(s *State) {
	g.Sticks.Throw(s)
	g.CalcValidMoves()
}
