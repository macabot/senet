package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/model"
)

func Board(board model.Board) *hypp.VNode {
	children := make([]*hypp.VNode, 30+len(board.Pieces[0])+len(board.Pieces[1]))
	i := 0
	for row := 0; row < 3; row++ {
		for column := 0; column < 10; column++ {
			position := model.Position{row, column}
			children[i] = Square(SquareProps{
				Position:    position,
				Selected:    board.Selected != nil && *board.Selected == position,
				Highlighted: board.Highlighted.Has(position),
				Protected:   board.IsProtected(position),
				Blocking:    board.IsBlocking(position),
			})
			i++
		}
	}
	for player, pieces := range board.Pieces {
		for piecePosition := range pieces {
			children[i] = Piece(PieceProps{
				Position:  piecePosition,
				Player:    player,
				CanSelect: player == board.You,
			})
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
