package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
)

type PieceProps struct {
	Piece         *state.Piece
	Player        int
	CanClick      bool
	DrawAttention bool
	Selected      bool
}

func Piece(props PieceProps) *hypp.VNode {
	coordinate := props.Piece.Position.Coordinate()
	hProps := hypp.HProps{
		"id": fmt.Sprintf("piece-%d", props.Piece.ID),
		"class": map[string]bool{
			"piece-wrapper":                             true,
			fmt.Sprintf("row-%d", coordinate.Row):       true,
			fmt.Sprintf("column-%d", coordinate.Column): true,
			"can-click": props.CanClick,
		},
	}
	if props.CanClick {
		hProps["onclick"] = hypp.ActionAndPayload[*state.State]{
			Action:  dispatch.SelectPiece,
			Payload: props.Piece.ID,
		}
	}
	return html.Div(
		hProps,
		html.Button(
			hypp.HProps{
				"class": map[string]bool{
					"piece":                                true,
					fmt.Sprintf("player-%d", props.Player): true,
					"selected":                             props.Selected,
					"draw-attention":                       props.DrawAttention,
					"protected":                            props.Piece.Ability == state.ProtectedPiece,
					"blocking":                             props.Piece.Ability == state.BlockingPiece,
				},
				"disabled": !props.CanClick,
				"type":     "button",
				"aria-label": fmt.Sprintf(
					"piece %d of player %d",
					((props.Piece.ID-1)%5)+1,
					props.Player+1,
				),
			},
			BlockingIcon(),
			ProtectedIcon(),
		),
	)
}
