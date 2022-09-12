package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/model"
	"github.com/macabot/senet/internal/app/view/render/hoc"
	"github.com/macabot/senet/internal/app/view/state"
)

type BoardProps struct {
	Board model.Board
	Meta  state.Meta
}

func Board(props BoardProps) *hypp.VNode {
	board := props.Board
	children := make([]*hypp.VNode, 30+len(board.Pieces[0])+len(board.Pieces[1]))
	i := 0
	for row := 0; row < 3; row++ {
		for column := 0; column < 10; column++ {
			position := model.Position{row, column}
			children[i] = hoc.With(
				Square(SquareProps{
					Position:    position,
					Selected:    props.Meta.Selected != nil && *props.Meta.Selected == position,
					Highlighted: props.Meta.Highlighted.Has(position),
					Protected:   board.IsProtected(position),
					Blocking:    board.IsBlocking(position),
				}),
				hoc.Key(fmt.Sprintf("square-%d-%d", position[0], position[1])),
			)
			i++
		}
	}
	for player, pieces := range board.Pieces {
		for _, piece := range pieces {
			children[i] = hoc.With(
				Piece(PieceProps{
					Piece:     piece,
					Player:    player,
					CanSelect: player == board.You,
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
