package atom

import (
	"github.com/macabot/hypp"
)

// TODO is this component necessary?
func ConditionalButton(label string, conditionIsMet bool, onClick hypp.Dispatchable) *hypp.VNode {
	if !conditionIsMet {
		onClick = nil
	}
	return Button(
		label,
		onClick,
		hypp.HProps{"class": map[string]bool{"cta": conditionIsMet}},
	)
}
