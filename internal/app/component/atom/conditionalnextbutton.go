package atom

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
)

func ConditionalNextButton(conditionIsMet bool, onClick hypp.Dispatchable) *hypp.VNode {
	if !conditionIsMet {
		onClick = nil
	}
	return html.Button(
		hypp.HProps{
			"class":   "cta",
			"onclick": onClick,
		},
		hypp.Text("Next"),
	)
}
