package tale

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/state"
)

func StartPage() *fairytale.Tale[*state.State] {
	return fairytale.New(
		"StartPage",
		&state.State{
			Page: state.StartPage,
		},
		component.Senet,
	).WithSettings(fairytale.TaleSettings{
		Target: fairytale.TaleAsHTML,
	})
}
