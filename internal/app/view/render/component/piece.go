package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/view/state"
)

type PieceProps struct {
	Piece     *state.Piece
	Player    int
	CanSelect bool
	Moving    bool
}

func Piece(props PieceProps) *hypp.VNode {
	return html.Div(
		hypp.HProps{
			"class": []string{
				"piece-wrapper",
				fmt.Sprintf("row-%d", props.Piece.Position[0]),
				fmt.Sprintf("column-%d", props.Piece.Position[1]),
			},
		},
		html.Button(
			hypp.HProps{
				"class": map[string]bool{
					"piece":                                true,
					fmt.Sprintf("player-%d", props.Player): true,
					"moving":                               props.Moving,
				},
				"disabled": !props.CanSelect,
			},
		),
	)
}
