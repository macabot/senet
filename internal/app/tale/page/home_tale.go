package page

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/component/page"
	"github.com/macabot/senet/internal/app/state"
)

func TaleHome() *fairytale.Tale[*state.State] {
	return fairytale.New(
		"Home",
		&state.State{},
		func(s *state.State) *hypp.VNode {
			return page.Home()
		},
	).WithSettings(fairytale.TaleSettings{
		Target: fairytale.TaleAsHTML,
	})
}
