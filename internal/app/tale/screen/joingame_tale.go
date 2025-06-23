package screen

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/senet/internal/app/component/screen"
	"github.com/macabot/senet/internal/app/state"
)

func TaleJoinGame() *fairytale.Tale[*state.State] {
	return fairytale.New(
		"JoinGame",
		&state.State{Screen: state.JoinGameScreen},
		screen.JoinGame,
	)
}
