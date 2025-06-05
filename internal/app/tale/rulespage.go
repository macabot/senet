package tale

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/state"
)

func RulesPage() *fairytale.Tale[*state.State] {
	return fairytale.New(
		"RulesPage",
		&state.State{
			Screen: state.RulesPage,
		},
		component.Senet,
	).WithSettings(fairytale.TaleSettings{
		Target: fairytale.TaleAsHTML,
	})
}
