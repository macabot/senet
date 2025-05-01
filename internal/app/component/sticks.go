package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
)

type SticksProps struct {
	Sticks        *state.Sticks
	DrawAttention bool
	NoValidMoves  bool
	IsLoading     bool
}

func Sticks(props SticksProps) *hypp.VNode {
	return html.Section(
		hypp.HProps{
			"class": "sticks",
		},
		Stick(props.Sticks.Flips[0]),
		Stick(props.Sticks.Flips[1]),
		Stick(props.Sticks.Flips[2]),
		Stick(props.Sticks.Flips[3]),
		throwButton(!props.DrawAttention || props.IsLoading),
		Loader(!props.IsLoading),
		steps(props.Sticks, props.NoValidMoves),
	)
}

func throwButton(disabled bool) *hypp.VNode {
	return html.Button(
		hypp.HProps{
			"class":    "throw-button",
			"disabled": disabled,
			"type":     "button",
			"onclick":  hypp.Action[*state.State](dispatch.ThrowSticks),
		},
	)
}

func stepsToIcon(steps int) *hypp.VNode {
	switch steps {
	case 1:
		return OneStepIcon()
	case 2:
		return TwoStepsIcon()
	case 3:
		return ThreeStepsIcon()
	case 4:
		return FourStepsIcon()
	case 6:
		return SixStepsIcon()
	default:
		panic(fmt.Errorf("there exists no icon for %d steps", steps))
	}
}

func steps(sticks *state.Sticks, noValidMoves bool) *hypp.VNode {
	return html.Div(
		hypp.HProps{
			"class": map[string]bool{
				"steps-wrapper":  true,
				"disabled":       !sticks.HasThrown,
				"can-go-again":   sticks.CanGoAgain(),
				"no-valid-moves": noValidMoves,
			},
		},
		stepsToIcon(sticks.Steps()),
		noMoveButton(),
	)
}

func noMoveButton() *hypp.VNode {
	return html.Button(
		hypp.HProps{
			"class":   "no-move-button",
			"onclick": dispatch.NoMove,
		},
		NoMoveIcon(),
	)
}
