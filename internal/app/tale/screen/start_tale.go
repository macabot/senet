package screen

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/component/screen"
	"github.com/macabot/senet/internal/app/state"
)

func TaleStart() *fairytale.Tale[*state.State] {
	return fairytale.New(
		"Start",
		&state.State{
			Screen: state.StartScreen,
		},
		func(s *state.State) *hypp.VNode {
			return screen.Start()
		},
	).WithSettings(fairytale.TaleSettings{
		Target: fairytale.TaleAsHTML,
	})
}
