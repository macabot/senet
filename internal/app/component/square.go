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
	Selected           *state.Piece
	CanClick           bool
	InvalidDestination bool
	IsStart            bool
	ShowDirection      bool
}

func iconToLabel(icon state.Icon) *hypp.VNode {
	switch icon {
	case state.Protected:
		return ProtectedIcon()
	case state.ReturnToStart:
		return ReturnToStartIcon()
	default:
		panic(fmt.Errorf("invalid icon %v", icon))
	}
}

func positionToArrowIcon(pos state.Position) *hypp.VNode {
	if pos < 9 {
		return FlowLeftIcon()
	} else if pos == 9 {
		return FlowUpIcon()
	} else if pos < 19 {
		return FlowRightIcon()
	} else if pos == 19 {
		return FlowUpIcon()
	} else {
		return FlowLeftIcon()
	}
}

func Square(props SquareProps) *hypp.VNode {
	coordinate := props.Position.Coordinate()
	hProps := hypp.HProps{
		"class": map[string]bool{
			"square":                                    true,
			fmt.Sprintf("row-%d", coordinate.Row):       true,
			fmt.Sprintf("column-%d", coordinate.Column): true,
			fmt.Sprintf("pos-%d", props.Position):       true,
		},
	}
	if props.CanClick && props.Selected != nil {
		hProps["onclick"] = hypp.ActionAndPayload[*state.State]{
			Action: dispatch.MoveToSquare,
			Payload: dispatch.Move{
				From: props.Selected.Position,
				To:   props.Position,
			},
		}
	}

	validReturnToStart := false
	var label *hypp.VNode
	if props.ShowDirection {
		label = positionToArrowIcon(props.Position)
	} else if special, ok := state.SpecialPositions[props.Position]; ok {
		label = iconToLabel(special.Icon)
		if props.CanClick && special.ReturnToStart {
			validReturnToStart = true
		}
	} else if props.IsStart {
		label = StartIcon()
	}

	validDestination := props.CanClick
	if validDestination && validReturnToStart {
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
				},
				"disabled":   !props.CanClick,
				"type":       "button",
				"aria-label": fmt.Sprintf("square %d", props.Position),
			},
			label,
		),
	)
}
