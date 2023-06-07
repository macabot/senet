package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
)

type SticksProps struct {
	Sticks        *state.Sticks
	DrawAttention bool
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
		throwButton(!props.DrawAttention),
		steps(props.Sticks),
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

func steps(sticks *state.Sticks) *hypp.VNode {
	return html.Div(
		hypp.HProps{
			"class": map[string]bool{
				"steps-wrapper": true,
				"disabled":      !sticks.HasThrown,
			},
		},
		html.Div(
			hypp.HProps{
				"class": map[string]bool{
					"steps":        true,
					"can-go-again": sticks.CanGoAgain(),
				},
			},
			hypp.Textf("%d", sticks.Steps()),
		),
	)
}
