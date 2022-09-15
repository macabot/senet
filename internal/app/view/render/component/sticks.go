package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
)

func Sticks(sticks [4]bool) *hypp.VNode {
	return html.Section(
		hypp.HProps{
			"class": "sticks",
		},
		Stick(sticks[0]),
		Stick(sticks[1]),
		Stick(sticks[2]),
		Stick(sticks[3]),
	)
}
