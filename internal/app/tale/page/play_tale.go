package page

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/senet/internal/app/component/page"
	"github.com/macabot/senet/internal/app/state"
)

func TalePlay() *fairytale.Tale[*state.State] {
	return fairytale.New(
		"Play",
		&state.State{},
		page.Play,
	).WithSettings(fairytale.TaleSettings{
		Target: fairytale.TaleAsHTML,
	})
}
