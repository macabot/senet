package tale

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/state"
	mycontrol "github.com/macabot/senet/internal/app/tale/control"
)

func GamePage() *fairytale.Tale[*state.State] {
	game := state.NewGame()
	game.TurnMode = state.IsPlayer1
	return fairytale.New(
		"GamePage",
		&state.State{
			Game: game,
			Page: state.GamePage,
		},
		component.Senet,
	).WithControls(
		mycontrol.Configuration(),
		mycontrol.Steps(),
		mycontrol.PlayerTurn(),
	).WithSettings(fairytale.TaleSettings{
		Target: fairytale.TaleAsHTML,
	})
}
