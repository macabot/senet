package tale

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/state"
)

func HomePage() *fairytale.Tale[*state.State] {
	return fairytale.New(
		"HomePage",
		&state.State{Page: state.HomePage},
		component.Senet,
	).WithSettings(fairytale.TaleSettings{
		Target: fairytale.TaleAsHTML,
	})
}
