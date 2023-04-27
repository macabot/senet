package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/state"
)

type PieceProps struct {
	Piece         state.Piece
	Player        int
	CanSelect     bool
	DrawAttention bool
	Moving        bool
	Selected      bool
	PieceAbility  state.PieceAbility
}

func Piece(props PieceProps) *hypp.VNode {
	coordinate := props.Piece.Position.Coordinate()
	var label *hypp.VNode
	switch props.PieceAbility {
	case state.BlockingPiece:
		// TODO
	case state.ProtectedPiece:
		label = shield()
	}
	return html.Div(
		hypp.HProps{
			"class": []string{
				"piece-wrapper",
				fmt.Sprintf("row-%d", coordinate.Row),
				fmt.Sprintf("column-%d", coordinate.Column),
			},
		},
		html.Button(
			hypp.HProps{
				"class": map[string]bool{
					"piece":                                true,
					fmt.Sprintf("player-%d", props.Player): true,
					"moving":                               props.Moving,
					"selected":                             props.Selected,
					"draw-attention":                       props.DrawAttention,
				},
				"disabled": !props.CanSelect,
				"type":     "button",
			},
			label,
		),
	)
}
