package page

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/component/page"
	"github.com/macabot/senet/internal/app/state"
)

func TaleRules() *fairytale.Tale[*state.State] {
	return fairytale.New(
		"Rules",
		&state.State{},
		func(s *state.State) *hypp.VNode {
			return page.Rules()
		},
	).WithSettings(fairytale.TaleSettings{
		Target: fairytale.TaleAsHTML,
	})
}
