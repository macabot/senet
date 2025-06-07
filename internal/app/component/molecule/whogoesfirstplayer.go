package molecule

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
)

func WhoGoesFirstPlayer(playerIndex int, name string) *hypp.VNode {
	return html.Div(
		hypp.HProps{
			"class": map[string]bool{
				"player-wrapper":                      true,
				fmt.Sprintf("player-%d", playerIndex): true,
			},
		},
		html.Div(
			hypp.HProps{
				"class": "player",
			},
			hypp.Text(name),
		),
	)
}
