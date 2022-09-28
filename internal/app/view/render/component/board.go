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
			position := state.Position{row, column}
			children[i] = hoc.With(
				Square(SquareProps{
					Position:    position,
					Selected:    props.Game.Board.Selected != nil && *&props.Game.Board.Selected.Position == position,
					Highlighted: false, // TODO
					Protected:   board.IsProtected(position),
					Blocking:    board.IsBlocking(position),
				}),
				hoc.Key(fmt.Sprintf("square-%d-%d", position[0], position[1])),
			)
			i++
		}
	}
	for player, pieces := range board.PlayerPieces {
		for _, piece := range pieces {
			children[i] = hoc.With(
				Piece(PieceProps{
					Piece:     piece,
					Player:    player,
					CanSelect: props.Game.Board.Selected == nil && player == props.Game.You && props.Game.Sticks.HasThrown,
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
