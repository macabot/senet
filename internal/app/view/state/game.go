package state

type Status int

const (
	Created Status = iota
	Ready
	InProgress
	Finished
)

type Game struct {
	board        *Board
	sticks       Sticks
	turn         int
	hasTurn      bool
	status       Status
	validMoves   map[Position]Position
	invalidMoves map[Position]Position
}

func NewGame() *Game {
	return &Game{
		board: NewBoard(),
	}
}

func (g Game) Board() *Board {
	return g.board
}

func (g *Game) SetBoard(board *Board) {
	g.board = board
	g.CalcValidMoves()
}

func (g Game) Sticks() Sticks {
	return g.sticks
}

func (g *Game) SetSticks(sticks Sticks) {
	g.sticks = sticks
	g.CalcValidMoves()
}

func (g Game) Turn() int {
	return g.turn
}

func (g *Game) SetTurn(turn int) {
	g.turn = turn
	g.CalcValidMoves()
}

func (g Game) HasTurn() bool {
	return g.hasTurn
}

func (g *Game) SetHasTurn(hasTurn bool) {
	g.hasTurn = hasTurn
}

func (g Game) ValidMoves() map[Position]Position {
	return g.validMoves
}

func (g Game) InvalidMoves() map[Position]Position {
	return g.invalidMoves
}

func (g Game) CanSelect(player int) bool {
	return g.hasTurn && g.sticks.HasThrown && player == g.turn
}

func (g *Game) CalcValidMoves() {
	g.validMoves = map[Position]Position{}
	g.invalidMoves = map[Position]Position{}

	piecesByPos := g.board.PlayerPieces[g.turn]
	otherPiecesByPos := g.board.PlayerPieces[(g.turn+1)%2]

	otherGroups := g.board.FindGroups(otherPiecesByPos)

	findMoves := func(steps int, protectedSize int) {
		for pos := range piecesByPos {
			toPos := pos + Position(steps)
			if toPos >= 30 || toPos < 0 {
				g.invalidMoves[pos] = toPos
				continue
			}
			if _, ok := piecesByPos[toPos]; ok {
				g.invalidMoves[pos] = toPos
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
				g.invalidMoves[pos] = toPos
				continue
			}
			if group, ok := otherGroups[toPos]; ok {
				if len(group) >= protectedSize {
					g.invalidMoves[pos] = toPos
					continue
				}
				if special, ok := SpecialPositions[toPos]; ok && special.Protects {
					g.invalidMoves[pos] = toPos
					continue
				}
			}

			g.validMoves[pos] = toPos

			if special, ok := SpecialPositions[toPos]; ok && special.Portal {
				for portalPos := Position(0); portalPos < 30; portalPos++ {
					if !piecesByPos.Has(portalPos) && !otherPiecesByPos.Has(Position(portalPos)) {
						g.validMoves[toPos] = portalPos
						break
					}
				}
			}
		}
	}

	steps := g.sticks.Steps()
	findMoves(steps, 2)
	if len(g.validMoves) == 0 {
		findMoves(-steps, 1)
	}
}

func (g Game) CanMove(player int, from Position) bool {
	piecesByPos := g.board.PlayerPieces[player]
	if _, ok := piecesByPos[from]; !ok {
		return false
	}
	if _, ok := g.validMoves[from]; !ok {
		return false
	}
	return true
}

func (g *Game) Move(player int, from Position) {
	piecesByPos := g.board.PlayerPieces[player]
	piece, ok := piecesByPos[from]
	if !ok {
		return
	}
	to, ok := g.validMoves[from]
	if !ok {
		return
	}
	piece.Position = to
	delete(piecesByPos, from)
	piecesByPos[to] = piece
}
