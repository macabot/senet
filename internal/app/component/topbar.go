package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/state"
)

func TopBar(s *state.State) *hypp.VNode {
	return html.Section(
		hypp.HProps{
			"class": "top-bar",
		},
		Players(CreatePlayersProps(s)),
	)
}
