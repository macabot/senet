package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/state"
)

type Player struct {
	*state.Player
	Points int
}

type PlayersProps struct {
	Players [2]Player
	Turn    int
}

func CreatePlayersProps(s *state.State) PlayersProps {
	return PlayersProps{
		Players: [2]Player{
			{
				Player: s.Game.Players[0],
				Points: s.Game.Board.Points(0),
			},
			{
				Player: s.Game.Players[1],
				Points: s.Game.Board.Points(1),
			},
		},
		Turn: s.Game.Turn,
	}
}

func Players(props PlayersProps) *hypp.VNode {
	return html.Section(
		hypp.HProps{
			"class": map[string]bool{
				"players":                              true,
				fmt.Sprintf("has-turn-%d", props.Turn): true,
			},
		},
		player(0, props.Players[0], props.Turn == 0),
		PlayerTurnIcon(),
		player(1, props.Players[1], props.Turn == 1),
	)
}
