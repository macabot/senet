package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/view/state"
)

func Sticks(props state.Sticks) *hypp.VNode {
	return html.Section(
		hypp.HProps{
			"class": "sticks",
		},
		Stick(props.Flips[0]),
		Stick(props.Flips[1]),
		Stick(props.Flips[2]),
		Stick(props.Flips[3]),
		throwButton(props.HasThrown),
	)
}

func throwButton(disabled bool) *hypp.VNode {
	return html.Button(
		hypp.HProps{
			"class":    "throw-button",
			"disabled": disabled,
			"type":     "button",
		},
	)
}
