package tale

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/state"
	"github.com/macabot/senet/internal/app/tale/control"
)

func Players() *fairytale.Tale[*state.State] {
	return fairytale.New(
		"Players",
		&state.State{
			Game: state.NewGame(),
		},
		func(s *state.State) *hypp.VNode {
			return component.Players(component.CreatePlayersProps(s))
		},
	).WithControls(
		control.PlayerTurn(),
	)
}
