package tale

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/state"
	mycontrol "github.com/macabot/senet/internal/app/tale/control"
)

func GamePage() *fairytale.Tale[*state.State] {
	game := state.NewGame()
	game.HasTurn = true
	return fairytale.New(
		"GamePage",
		&state.State{
			Game: game,
		},
		component.GamePage,
	).WithControls(
		mycontrol.Configuration(),
		mycontrol.Steps(),
		mycontrol.PlayerTurn(),
	)
}
