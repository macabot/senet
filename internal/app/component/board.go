package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/state"
	"github.com/macabot/senet/internal/pkg/set"
)

func Board(props *state.State) *hypp.VNode {
	board := props.Game.Board
	children := make([]*hypp.VNode, 30+len(board.PlayerPieces[0])+len(board.PlayerPieces[1]))
	i := 0
	selected := props.Game.Selected
	var validDestination state.Position
	hasValidDestination := false
	var invalidDestinations set.Set[state.Position]
	if selected != nil {
		validDestination, hasValidDestination = props.Game.ValidMoves[selected.Position]
		invalidDestinations = props.Game.InvalidMoves[selected.Position]
	}
	for row := 0; row < 3; row++ {
		for column := 0; column < 10; column++ {
			coordinate := state.Coordinate{Row: row, Column: column}
			position := state.PositionFromCoordinate(coordinate)
			children[i] = With(
				Square(SquareProps{
					Position:           position,
					ValidDestination:   hasValidDestination && validDestination == position,
					InvalidDestination: invalidDestinations.Has(position),
					Selected:           selected != nil && selected.Position == position,
				}),
				Key(fmt.Sprintf("square-%d", position)),
			)
			i++
		}
	}
	for player, piecesByPos := range board.PlayerPieces {
		pieces := piecesByPos.OrderedByID()
		for _, piece := range pieces {
			children[i] = With(
				Piece(PieceProps{
					Piece:         piece,
					Player:        player,
					CanSelect:     props.Game.CanSelect(player),
					DrawAttention: props.Game.DrawsAttention(player),
					Selected:      props.Game.IsSelected(piece),
				}),
				Key(fmt.Sprintf("piece-%d", piece.ID)),
			)
			i++
		}
	}
	return html.Section(
		hypp.HProps{
			"class": "board",
		},
		children...,
	)
}
