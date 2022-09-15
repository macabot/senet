package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
)

func Stick(up bool) *hypp.VNode {
	return html.Div(
		hypp.HProps{
			"class": map[string]bool{
				"stick": true,
				"up":    up,
			},
		},
	)
}
