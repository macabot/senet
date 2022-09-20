package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
)

func Stick(flips int) *hypp.VNode {
	return html.Div(
		hypp.HProps{
			"class": "stick",
		},
		html.Div(
			hypp.HProps{
				"class": "stick-inner",
				"style": map[string]string{
					"transform": fmt.Sprintf("rotateX(%ddeg)", 180*flips),
				},
			},
			html.Div(
				hypp.HProps{
					"class": "stick-front",
				},
			),
			html.Div(
				hypp.HProps{
					"class": "stick-back",
				},
			),
		),
	)
}
