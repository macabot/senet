package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/state"
)

type SquareProps struct {
	Position           state.Position
	Protected          bool
	Blocking           bool
	ValidDestination   bool
	InvalidDestination bool
	Selected           bool
}

func iconToLabel(icon state.Icon) *hypp.VNode {
	switch icon {
	case state.Shield:
		return protectedIcon()
	case state.Cross:
		return hypp.Text("☓")
	default:
		panic(fmt.Errorf("invalid icon %v", icon))
	}
}

func Square(props SquareProps) *hypp.VNode {
	var label *hypp.VNode
	if special, ok := state.SpecialPositions[props.Position]; ok {
		label = iconToLabel(special.Icon)
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
			label,
		),
	)
}
