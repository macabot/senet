package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/view/render/hoc"
	"github.com/macabot/senet/internal/app/view/state"
)

func Board(props *state.State) *hypp.VNode {
	board := props.Game.Board
	children := make([]*hypp.VNode, 30+len(board.PlayerPieces[0])+len(board.PlayerPieces[1]))
	i := 0
	for row := 0; row < 3; row++ {
		for column := 0; column < 10; column++ {
			coordinate := state.Coordinate{Row: row, Column: column}
			position := state.PositionFromCoordinate(coordinate)
			children[i] = hoc.With(
				Square(SquareProps{
					Position:    position,
					Highlighted: false, // TODO
					// Protected:   board.IsProtected(position),
					// Blocking:    board.IsBlocking(position),
				}),
				hoc.Key(fmt.Sprintf("square-%d", position)),
			)
			i++
		}
	}
	selected := props.Game.Board.Selected
	for player, piecesByPos := range board.PlayerPieces {
		pieces := piecesByPos.OrderedByID()
		for _, piece := range pieces {
			children[i] = hoc.With(
				Piece(PieceProps{
					Piece:     piece,
					Player:    player,
					CanSelect: selected == nil && player == props.Game.You && props.Game.Sticks.HasThrown,
					Selected:  selected != nil && selected.Position == piece.Position,
				}),
				hoc.Key(fmt.Sprintf("piece-%d", piece.ID)),
			)
			i++
		}
	}
	return html.Main(
		hypp.HProps{
			"class": "board",
		},
		children...,
	)
}
