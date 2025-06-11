package organism

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/component/atom"
	"github.com/macabot/senet/internal/app/component/molecule"
	"github.com/macabot/senet/internal/app/state"
)

type PlayersProps struct {
	Players [2]molecule.PlayerProps
	Turn    int
}

func CreatePlayersProps(s *state.State) PlayersProps {
	return PlayersProps{
		Players: [2]molecule.PlayerProps{
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
	classes := map[string]bool{
		"players": true,
	}
	if props.Turn == 0 {
		classes["has-turn-0"] = true
	} else if props.Turn == 1 {
		classes["has-turn-1"] = true
	}

	return html.Section(
		hypp.HProps{
			"class": classes,
		},
		molecule.Player(0, props.Players[0], props.Turn == 0),
		atom.PlayerTurnIcon(),
		molecule.Player(1, props.Players[1], props.Turn == 1),
	)
}
