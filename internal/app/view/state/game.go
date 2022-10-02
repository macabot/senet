package state

type Status int

const (
	Created Status = iota
	Ready
	InProgress
	Finished
)

type Game struct {
	Board      Board
	Sticks     Sticks
	You        int
	Turn       int
	Status     Status
	ValidMoves map[Position]Position
}

func (g *Game) CalcValidMoves(player int) {
	g.ValidMoves = map[Position]Position{}

	piecesByPos := g.Board.PlayerPieces[player]
	otherPiecesByPos := g.Board.PlayerPieces[(player+1)%2]

	otherGroups := g.Board.FindGroups(otherPiecesByPos)

	findMoves := func(steps int, protectedSize int) {
		for pos := range piecesByPos {
			toPos := pos + Position(steps)
			if toPos >= 30 || toPos < 0 {
				continue
			}
			if _, ok := piecesByPos[toPos]; ok {
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
				continue
			}
			if group, ok := otherGroups[toPos]; ok {
				if len(group) >= protectedSize {
					continue
				}
				if special, ok := SpecialPositions[toPos]; ok && special.Protects {
					continue
				}
			}

			g.ValidMoves[pos] = toPos

			if special, ok := SpecialPositions[toPos]; ok && special.Portal {
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
