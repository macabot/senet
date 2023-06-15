package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/state"
)

type SquareProps struct {
	Position           state.Position
	InvalidDestination bool
	IsStart            bool
	OnClick            hypp.Dispatchable
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
	coordinate := props.Position.Coordinate()
	hProps := hypp.HProps{
		"class": map[string]bool{
			"square":                                    true,
			fmt.Sprintf("row-%d", coordinate.Row):       true,
			fmt.Sprintf("column-%d", coordinate.Column): true,
		},
	}
	if props.OnClick != nil {
		hProps["onclick"] = props.OnClick
	}
	validDestination := props.OnClick != nil
	validReturnToStart := false
	var label *hypp.VNode
	if special, ok := state.SpecialPositions[props.Position]; ok {
		label = iconToLabel(special.Icon)
		if validDestination && special.ReturnToStart {
			validReturnToStart = true
		}
	}
	if props.IsStart {
		label = startIcon()
	}
	if validDestination && (validReturnToStart || props.IsStart) {
		validDestination = false
	}
	return html.Div(
		hProps,
		html.Button(
			hypp.HProps{
				"class": map[string]bool{
					"inner-square":          true,
					"valid-destination":     validDestination,
					"invalid-destination":   props.InvalidDestination,
					"valid-return-to-start": validReturnToStart,
					"is-start":              props.IsStart,
				},
				"disabled": props.OnClick == nil,
				"type":     "button",
			},
			label,
		),
	)
}
