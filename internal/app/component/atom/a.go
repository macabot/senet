package atom

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
)

func A(label string, href string, props hypp.HProps) *hypp.VNode {
	return html.A(
		hypp.MergeHProps(
			hypp.HProps{
				"href": href,
			},
			props,
		),
		hypp.Text(label),
	)
}
