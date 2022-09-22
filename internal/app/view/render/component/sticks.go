package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
)

type SticksProps struct {
	Sticks   [4]int
	CanThrow bool
}

func Sticks(props SticksProps) *hypp.VNode {
	return html.Section(
		hypp.HProps{
			"class": "sticks",
		},
		Stick(props.Sticks[0]),
		Stick(props.Sticks[1]),
		Stick(props.Sticks[2]),
		Stick(props.Sticks[3]),
		throwButton(!props.CanThrow),
	)
}

func throwButton(disabled bool) *hypp.VNode {
	return html.Button(
		hypp.HProps{
			"class":    "throw-button",
			"disabled": disabled,
		},
	)
}
