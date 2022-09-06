package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/model"
)

type SquareProps struct {
	Position    model.Position
	Selected    bool
	Highlighted bool
}

var iconToLabel = map[model.Icon]string{
	model.Two:   "II",
	model.Three: "III",
	model.Cross: "☓",
	model.Ankh:  "☥",
}

func Square(props SquareProps) *hypp.VNode {
	var text *hypp.VNode
	if icon, ok := model.SpecialPositions[props.Position]; ok {
		text = hypp.Text(iconToLabel[icon])
	}
	return html.Div(
		hypp.HProps{
			"class": map[string]bool{
				"square":                                    true,
				fmt.Sprintf("row-%d", props.Position[0]):    true,
				fmt.Sprintf("column-%d", props.Position[1]): true,
			},
		},
		html.Button(
			hypp.HProps{
				"class": map[string]bool{
					"inner-square": true,
					"selected":     props.Selected,
					"highlighted":  props.Highlighted,
				},
				"disabled": !props.Highlighted,
			},
			text,
		),
	)
}
