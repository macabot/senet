package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/model"
)

type PieceProps struct {
	Position  model.Position
	Player    int
	CanSelect bool
}

func Piece(props PieceProps) *hypp.VNode {
	return html.Div(
		hypp.HProps{
			"class": []string{
				"piece-wrapper",
				fmt.Sprintf("row-%d", props.Position[0]),
				fmt.Sprintf("column-%d", props.Position[1]),
			},
		},
		html.Button(
			hypp.HProps{
				"class": []string{
					"piece",
					fmt.Sprintf("player-%d", props.Player),
				},
				"disabled": !props.CanSelect,
			},
		),
	)
}
