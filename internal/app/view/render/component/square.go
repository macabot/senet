package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/view/state"
)

type SquareProps struct {
	Position    state.Position
	Selected    bool
	Highlighted bool
	Protected   bool
	Blocking    bool
}

var iconToLabel = map[state.Icon]string{
	state.Two:   "II",
	state.Three: "III",
	state.Cross: "☓",
	state.Ankh:  "☥",
}

func Square(props SquareProps) *hypp.VNode {
	var text *hypp.VNode
	if special, ok := state.SpecialPositions[props.Position]; ok {
		text = hypp.Text(iconToLabel[special.Icon])
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
					"protected":    props.Protected,
					"blocking":     props.Blocking,
				},
				"disabled": !props.Highlighted,
			},
			text,
		),
	)
}
