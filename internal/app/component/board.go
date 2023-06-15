package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/dispatch"
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
	startPosition := props.Game.Board.StartPosition()
	for row := 0; row < 3; row++ {
		for column := 0; column < 10; column++ {
			coordinate := state.Coordinate{Row: row, Column: column}
			position := state.PositionFromCoordinate(coordinate)
			var onClick hypp.Dispatchable
			isValidDestination := hasValidDestination && validDestination == position
			isStart := hasValidDestination && validDestination == state.ReturnToStartPosition && position == startPosition
			if isStart {
				onClick = dispatch.MoveToSquareAction(state.ReturnToStartPosition)
			} else if isValidDestination {
				onClick = dispatch.MoveToSquareAction(position)
			}
			children[i] = With(
				Square(SquareProps{
					Position:           position,
					InvalidDestination: invalidDestinations.Has(position),
					IsStart:            isStart,
					OnClick:            onClick,
				}),
				Key(fmt.Sprintf("square-%d", position)),
			)
			i++
		}
	}
	for player, piecesByPos := range board.PlayerPieces {
		pieces := piecesByPos.OrderedByID()
		drawAttention := props.Game.PiecesDrawAttention(player)
		for _, piece := range pieces {
			isSelected := props.Game.PieceIsSelected(piece)
			children[i] = With(
				Piece(PieceProps{
					Piece:         piece,
					Player:        player,
					CanClick:      props.Game.CanClickOnPiece(player, piece),
					DrawAttention: drawAttention && !isSelected,
					Selected:      isSelected,
				}),
				Key(fmt.Sprintf("piece-%d", piece.ID)),
			)
			i++
		}
	}
	return html.Section(
		hypp.HProps{
			"class": map[string]bool{
				"board":                true,
				"no-animation":         !props.Game.HasMoved,
				"selected-change-even": props.Game.SelectedChangeCounter%2 == 0,
				"selected-change-odd":  props.Game.SelectedChangeCounter%2 != 0,
			},
		},
		children...,
	)
}
