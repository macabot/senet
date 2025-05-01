package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
)

func Stick(flips int) *hypp.VNode {
	const clones = 20
	children := make([]*hypp.VNode, clones)
	for i := 0; i < clones; i++ {
		children[i] = innerStick(flips, i)
	}
	return html.Div(
		hypp.HProps{
			"class": "stick-container",
		},
		children...,
	)
}

func innerStick(flips int, i int) *hypp.VNode {
	return html.Div(
		hypp.HProps{
			"class": "stick",
		},
		html.Div(
			hypp.HProps{
				"class": []string{"stick-inner", fmt.Sprintf("delay-%d", i)},
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
