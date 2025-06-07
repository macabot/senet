package atom

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
)

// TODO change type of 'label' to 'string | *hypp.VNode
func Button(label string, onClick hypp.Dispatchable, props hypp.HProps) *hypp.VNode {
	return html.Button(
		hypp.MergeHProps(
			hypp.HProps{
				"type":    "button",
				"onclick": onClick,
			},
			props,
		),
		hypp.Text(label),
	)
}
