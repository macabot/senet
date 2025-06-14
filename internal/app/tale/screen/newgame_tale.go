package screen

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/senet/internal/app/component/screen"
	"github.com/macabot/senet/internal/app/state"
)

func TaleNewGame() *fairytale.Tale[*state.State] {
	return fairytale.New(
		"NewGame",
		&state.State{Screen: state.NewGameScreen},
		screen.NewGame,
	)
}
