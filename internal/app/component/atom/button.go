package atom

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
)

type ButtonLabel interface {
	string | *hypp.VNode
}

func Button[T ButtonLabel](label T, onClick hypp.Dispatchable, props hypp.HProps) *hypp.VNode {
	var labelNode *hypp.VNode
	switch l := any(label).(type) {
	case string:
		labelNode = hypp.Text(l)
	case *hypp.VNode:
		labelNode = l
	}

	return html.Button(
		hypp.MergeHProps(
			hypp.HProps{
				"type":    "button",
				"onclick": onClick,
			},
			props,
		),
		labelNode,
	)
}
