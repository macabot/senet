package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
)

type SquareProps struct {
	Position           state.Position
	ValidDestination   bool
	InvalidDestination bool
}

func iconToLabel(icon state.Icon) *hypp.VNode {
	switch icon {
	case state.Protected:
		return protectedIcon()
	case state.ReturnToStart:
		return returnToStartIcon()
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
	hProps := hypp.HProps{
		"class": map[string]bool{
			"square":                                    true,
			fmt.Sprintf("row-%d", coordinate.Row):       true,
			fmt.Sprintf("column-%d", coordinate.Column): true,
		},
	}
	if props.ValidDestination {
		hProps["onclick"] = dispatch.MoveToSquareAction(props.Position)
	}
	return html.Div(
		hProps,
		html.Button(
			hypp.HProps{
				"class": map[string]bool{
					"inner-square":        true,
					"valid-destination":   props.ValidDestination,
					"invalid-destination": props.InvalidDestination,
				},
				"disabled": !props.ValidDestination,
				"type":     "button",
			},
			label,
		),
	)
}
