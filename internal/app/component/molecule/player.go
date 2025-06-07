package molecule

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/component/atom"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
)

func Player(playerIndex int, player Player, hasTurn bool) *hypp.VNode {
	var bubble *hypp.VNode
	if player.SpeechBubble != nil {
		bubble = SpeechBubble(playerIndex, player.SpeechBubble)
	}
	return html.Div(
		hypp.HProps{
			"class": map[string]bool{
				"player-wrapper":                      true,
				fmt.Sprintf("player-%d", playerIndex): true,
			},
		},
		html.Button(
			hypp.HProps{
				"class": map[string]bool{
					"player":         true,
					"has-turn":       hasTurn,
					"draw-attention": player.DrawAttention,
				},
				"onclick": hypp.ActionAndPayload[*state.State]{
					Action:  dispatch.ToggleSpeechBubble,
					Payload: playerIndex,
				},
			},
			html.Span(nil, hypp.Text(player.Name)),
			atom.PointsIcon(player.Points),
		),
		bubble,
	)
}
