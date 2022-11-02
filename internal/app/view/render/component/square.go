package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/view/state"
)

type SquareProps struct {
	Position           state.Position
	Protected          bool
	Blocking           bool
	ValidDestination   bool
	InvalidDestination bool
	Selected           bool
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
	coordinate := props.Position.Coordinate()
	return html.Div(
		hypp.HProps{
			"class": map[string]bool{
				"square":                                    true,
				fmt.Sprintf("row-%d", coordinate.Row):       true,
				fmt.Sprintf("column-%d", coordinate.Column): true,
			},
		},
		html.Button(
			hypp.HProps{
				"class": map[string]bool{
					"inner-square":        true,
					"protected":           props.Protected,
					"blocking":            props.Blocking,
					"valid-destination":   props.ValidDestination,
					"invalid-destination": props.InvalidDestination,
					"selected":            props.Selected,
				},
				"disabled": !props.ValidDestination,
				"type":     "button",
			},
			text,
		),
	)
}
